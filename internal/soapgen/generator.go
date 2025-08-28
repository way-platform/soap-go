package soapgen

import (
	"fmt"
	"sort"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

// Config holds configuration for code generation
type Config struct {
	PackageName string
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

	return nil
}

// Files returns the generated files
func (g *Generator) Files() []*codegen.File {
	return g.files
}

// generateTypesFile generates a Go file with types from an XSD schema
func (g *Generator) generateTypesFile(schema *xsd.Schema, packageName, filename string) (*codegen.File, error) {
	file := codegen.NewFile(filename)

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

	// Collect all required imports by analyzing the types used
	requiredImports := make(map[string]bool)

	// Always add encoding/xml for XMLName fields
	requiredImports["encoding/xml"] = true

	for _, element := range append(dataTypes, messageTypes...) {
		collectRequiredImports(element, requiredImports, ctx)
	}

	// Add imports in sorted order for deterministic output
	var imports []string
	for imp := range requiredImports {
		imports = append(imports, imp)
	}
	sort.Strings(imports)
	for _, imp := range imports {
		file.Import(imp)
	}

	if len(requiredImports) > 0 {
		file.P()
	}

	// Generate simple type constants first (for enumerations)
	generateSimpleTypeConstants(file, ctx)

	// Generate complex types that are referenced but not top-level elements
	generateComplexTypes(file, ctx)

	// Generate inline complex types first (before elements that use them)
	generateInlineComplexTypes(file, ctx, dataTypes)

	// Generate data types (for proper dependency ordering)
	for _, element := range dataTypes {
		if typeRegistry.shouldGenerate(element) {
			generateStructFromElement(file, element, ctx)
		}
	}

	// Generate message wrapper types with *Wrapper suffix to avoid conflicts
	for _, element := range messageTypes {
		wrapperTypeName := toGoName(element.Name) + "Wrapper"
		if typeRegistry.shouldGenerateWithName(element, wrapperTypeName) {
			generateStructFromElementWithWrapper(file, element, ctx)
		}
	}

	return file, nil
}

// needsRawXML checks if the schema contains any inline complex types that would require RawXML
func needsRawXML(schema *xsd.Schema) bool {
	return hasInlineComplexTypes(schema.Elements)
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
