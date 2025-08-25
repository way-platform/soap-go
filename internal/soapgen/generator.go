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

	// Separate data types from message wrapper types
	dataTypes, messageTypes := categorizeElements(schema.Elements)

	// Collect all required imports by analyzing the types used
	requiredImports := make(map[string]bool)
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

	// Skip message wrapper types for now as they cause duplicates
	// TODO: In the future, we might want to generate these with different names
	// or in a separate package for SOAP message handling

	return file, nil
}
