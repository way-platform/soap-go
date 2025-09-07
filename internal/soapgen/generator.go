package soapgen

import (
	"fmt"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

// Config holds configuration for code generation
type Config struct {
	PackageName    string
	GenerateClient bool // Whether to generate SOAP client code
}

// Generator generates Go code from WSDL definitions
type Generator struct {
	definitions *wsdl.Definitions
	config      Config
	files       []*codegen.File
}

// NewGenerator creates a new Generator with the given WSDL definitions and config
func NewGenerator(definitions *wsdl.Definitions, config Config) *Generator {
	return &Generator{
		definitions: definitions,
		config:      config,
		files:       make([]*codegen.File, 0),
	}
}

// Generate generates Go code files from the WSDL definitions
func (g *Generator) Generate() error {
	// Check if we have types to generate
	if g.definitions.Types == nil || len(g.definitions.Types.Schemas) == 0 {
		return fmt.Errorf("no schema types found in WSDL definition")
	}

	// Generate Go code for each schema
	for i, schema := range g.definitions.Types.Schemas {
		filename := "types.go"
		if i > 0 {
			filename = fmt.Sprintf("types_%d.go", i+1)
		}

		file, err := g.generateTypesFile(&schema, g.config.PackageName, filename)
		if err != nil {
			return fmt.Errorf("failed to generate types file: %w", err)
		}

		g.files = append(g.files, file)
	}

	// Generate client file if requested
	if g.config.GenerateClient {
		clientFile, err := g.generateClientFile(g.config.PackageName, "client.go")
		if err != nil {
			return fmt.Errorf("failed to generate client file: %w", err)
		}
		if clientFile != nil {
			g.files = append(g.files, clientFile)
		}
	}

	return nil
}

// Files returns the generated files
func (g *Generator) Files() []*codegen.File {
	return g.files
}

// generateTypesFile generates a Go file with types from an XSD schema
func (g *Generator) generateTypesFile(schema *xsd.Schema, packageName, filename string) (*codegen.File, error) {
	file := codegen.NewFile(filename, packageName)

	// Set custom package name for soap-go to use "soap" instead of "soapgo"
	file.SetPackageName("github.com/way-platform/soap-go", "soap")

	// Create schema context for reference resolution
	ctx := newSchemaContext(schema)

	// Create type registry to prevent duplicates
	typeRegistry := newTypeRegistry()

	// Add package declaration
	file.P("package ", packageName)
	file.P()

	// Generate RawXML type definition if needed
	if needsRawXML(schema) {
		file.P("// RawXML captures raw XML content for untyped elements.")
		file.P("type RawXML []byte")
		file.P()
	}

	// Separate data types from message wrapper types
	dataTypes, messageTypes := categorizeElements(schema.Elements)

	// Import handling is now automatic via QualifiedGoIdent calls

	// Generate simple type constants first (for enumerations)
	generateSimpleTypeConstants(file, ctx)

	// Generate complex types that are referenced but not top-level elements
	generateComplexTypes(file, ctx)

	// Generate inline complex types first (before elements that use them)
	generateInlineComplexTypes(file, ctx, dataTypes)

	// Generate data types with classification-based wrapper naming
	bindingStyle := g.getBindingStyle()
	processedElements := make(map[string]bool)
	processedGoTypes := make(map[string]bool) // Track processed Go type names to prevent duplicates

	// First pass: Generate wrapper types for operation elements
	allElements := append([]*xsd.Element{}, dataTypes...)
	allElements = append(allElements, messageTypes...)

	for _, element := range allElements {
		if processedElements[element.Name] {
			continue // Skip duplicates
		}
		processedElements[element.Name] = true

		// Check if we would generate a duplicate Go type name
		var goTypeName string
		if g.shouldUseWrapperForElement(element.Name, bindingStyle) {
			goTypeName = toGoName(element.Name) + "Wrapper"
		} else {
			goTypeName = toGoName(element.Name)
		}

		if processedGoTypes[goTypeName] {
			continue // Skip elements that would generate duplicate Go type names
		}
		processedGoTypes[goTypeName] = true

		if g.shouldUseWrapperForElement(element.Name, bindingStyle) {
			if typeRegistry.shouldGenerateWithContext(element, SOAPWrapperContext) {
				generateStructFromElementWithWrapper(file, element, ctx, typeRegistry)
			}
		} else {
			if typeRegistry.shouldGenerateWithContext(element, DataElementContext) {
				generateStructFromElement(file, element, ctx, typeRegistry)
			}
		}
	}

	// All elements have been processed in the two passes above

	return file, nil
}

// needsRawXML checks if the schema contains any constructs that would require RawXML
func needsRawXML(schema *xsd.Schema) bool {
	// Check for inline complex types in elements
	if hasInlineComplexTypes(schema.Elements) {
		return true
	}

	// Check for xs:any elements in schema elements
	if hasAnyElements(schema.Elements) {
		return true
	}

	// Check for xs:any elements in named complex types
	for _, complexType := range schema.ComplexTypes {
		if hasAnyElementsInComplexType(&complexType) {
			return true
		}
	}

	// Check for untyped/unknown elements that will fallback to RawXML
	if hasUntypedElements(schema.Elements) {
		return true
	}

	// Check for untyped elements in named complex types
	for _, complexType := range schema.ComplexTypes {
		if hasUntypedElementsInComplexType(&complexType) {
			return true
		}
	}

	return false
}

// hasInlineComplexTypes recursively checks if any elements have inline complex types
func hasInlineComplexTypes(elements []xsd.Element) bool {
	for _, element := range elements {
		if hasInlineComplexType(&element) {
			return true
		}
	}
	return false
}

// hasInlineComplexType checks if an element has inline complex types
func hasInlineComplexType(element *xsd.Element) bool {
	if element.ComplexType != nil {
		// Check sequence elements
		if element.ComplexType.Sequence != nil {
			for _, field := range element.ComplexType.Sequence.Elements {
				if field.ComplexType != nil {
					return true
				}
				if field.Ref != "" {
					// Element references might also need checking, but for now we're conservative
					continue
				}
			}
		}

		// Check extension sequence elements
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, field := range ext.Sequence.Elements {
					if field.ComplexType != nil {
						return true
					}
				}
			}
		}
	}
	return false
}

// hasAnyElements checks if any elements contain xs:any elements
func hasAnyElements(elements []xsd.Element) bool {
	for _, element := range elements {
		if hasAnyElementInElement(&element) {
			return true
		}
	}
	return false
}

// hasAnyElementInElement checks if an element contains xs:any elements
func hasAnyElementInElement(element *xsd.Element) bool {
	if element.ComplexType != nil {
		return hasAnyElementsInComplexType(element.ComplexType)
	}
	return false
}

// hasAnyElementsInComplexType checks if a complex type contains xs:any elements
func hasAnyElementsInComplexType(complexType *xsd.ComplexType) bool {
	// Check sequence for xs:any elements
	if complexType.Sequence != nil && len(complexType.Sequence.Any) > 0 {
		return true
	}

	// Check extension sequences for xs:any elements
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil && len(ext.Sequence.Any) > 0 {
			return true
		}
	}

	return false
}

// hasUntypedElements checks if any elements will fallback to RawXML due to unknown types
func hasUntypedElements(elements []xsd.Element) bool {
	for _, element := range elements {
		if hasUntypedElementInElement(&element) {
			return true
		}
	}
	return false
}

// hasUntypedElementInElement checks if an element contains untyped elements
func hasUntypedElementInElement(element *xsd.Element) bool {
	if element.ComplexType != nil {
		return hasUntypedElementsInComplexType(element.ComplexType)
	}
	return false
}

// hasUntypedElementsInComplexType checks if a complex type contains elements that will use RawXML
func hasUntypedElementsInComplexType(complexType *xsd.ComplexType) bool {
	// Check sequence elements
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			// Elements with inline complex types but no proper type mapping will use RawXML
			if field.Type == "" && field.ComplexType != nil {
				return true
			}
			// Elements with empty/unknown types will fallback to RawXML
			if field.Type == "" && field.ComplexType == nil && field.Ref == "" {
				return true
			}
		}
	}

	// Check extension sequence elements
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil {
			for _, field := range ext.Sequence.Elements {
				// Elements with inline complex types but no proper type mapping will use RawXML
				if field.Type == "" && field.ComplexType != nil {
					return true
				}
				// Elements with empty/unknown types will fallback to RawXML
				if field.Type == "" && field.ComplexType == nil && field.Ref == "" {
					return true
				}
			}
		}
	}

	return false
}
