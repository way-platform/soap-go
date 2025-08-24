package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

// NewCommand creates a new [cobra.Command] for the gen command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen",
		Short:   "Generate code for a SOAP API",
		GroupID: "gen",
	}
	inputFile := cmd.Flags().StringP("input", "i", "", "input WSDL file (required)")
	_ = cmd.MarkFlagRequired("input")
	outputDir := cmd.Flags().StringP("dir", "d", "", "output directory (required)")
	_ = cmd.MarkFlagRequired("dir")
	packageName := cmd.Flags().StringP("package", "p", "", "Go package name (required)")
	_ = cmd.MarkFlagRequired("package")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(config{
			inputFile:   *inputFile,
			outputDir:   *outputDir,
			packageName: *packageName,
		})
	}
	return cmd
}

type config struct {
	inputFile   string
	outputDir   string
	packageName string
}

func run(cfg config) error {
	// Parse the WSDL file
	defs, err := wsdl.ParseFromFile(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("failed to parse WSDL file: %w", err)
	}

	// Check if we have types to generate
	if defs.Types == nil || len(defs.Types.Schemas) == 0 {
		return fmt.Errorf("no schema types found in WSDL file")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(cfg.outputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate Go code for each schema
	for i, schema := range defs.Types.Schemas {
		filename := "types.go"
		if i > 0 {
			filename = fmt.Sprintf("types_%d.go", i+1)
		}

		if err := generateTypesFile(&schema, cfg.packageName, cfg.outputDir, filename); err != nil {
			return fmt.Errorf("failed to generate types file: %w", err)
		}
	}

	return nil
}

// generateTypesFile generates a Go file with types from an XSD schema
func generateTypesFile(schema *xsd.Schema, packageName, outputDir, filename string) error {
	outputPath := filepath.Join(outputDir, filename)
	g := codegen.NewFile(outputPath)

	// Add package declaration
	g.P("package ", packageName)
	g.P()

	// Collect all required imports by analyzing the types used
	requiredImports := make(map[string]bool)
	for _, element := range schema.Elements {
		if element.ComplexType != nil {
			collectRequiredImports(&element, requiredImports)
		}
	}

	// Add imports
	for imp := range requiredImports {
		g.Import(imp)
	}

	if len(requiredImports) > 0 {
		g.P()
	}

	// Generate types for each top-level element
	for _, element := range schema.Elements {
		if element.ComplexType != nil {
			generateStructFromElement(g, &element)
		}
	}

	// Get the generated content
	content, err := g.Content()
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, content, 0o644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

// collectRequiredImports recursively collects all import statements needed for the given element
func collectRequiredImports(element *xsd.Element, imports map[string]bool) {
	if element.ComplexType != nil && element.ComplexType.Sequence != nil {
		for _, fieldElement := range element.ComplexType.Sequence.Elements {
			// Parse the type and check if it requires any imports
			parsedType := xsd.ParseType(fieldElement.Type)
			for _, imp := range parsedType.RequiresImport() {
				imports[imp] = true
			}
		}
	}
}

// generateStructFromElement generates a Go struct from an XSD element
func generateStructFromElement(g *codegen.File, element *xsd.Element) {
	structName := toGoName(element.Name)

	// Add comment
	g.P("// ", structName, " represents the ", element.Name, " element")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	// Generate fields from the complex type
	if element.ComplexType != nil && element.ComplexType.Sequence != nil {
		for _, field := range element.ComplexType.Sequence.Elements {
			generateStructField(g, &field)
		}
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructField generates a Go struct field from an XSD element
func generateStructField(g *codegen.File, element *xsd.Element) {
	fieldName := toGoName(element.Name)
	goType := mapXSDTypeToGo(element.Type)
	xmlName := element.Name

	// Generate the field with XML tag
	g.P("	", fieldName, " ", goType, " `xml:\"", xmlName, "\"`")
}

// toGoName converts an XML name to a Go identifier (PascalCase)
func toGoName(name string) string {
	if name == "" {
		return ""
	}

	// Split on common separators and capitalize each part
	parts := strings.FieldsFunc(name, func(r rune) bool {
		return r == '_' || r == '-' || r == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	// Handle the case where name doesn't need splitting
	if len(parts) <= 1 {
		result.Reset()
		result.WriteString(strings.ToUpper(name[:1]))
		if len(name) > 1 {
			result.WriteString(name[1:])
		}
	}

	return result.String()
}

// mapXSDTypeToGo maps XSD types to Go types using the xsd10 type system.
func mapXSDTypeToGo(xsdType string) string {
	// Parse the XSD type using our comprehensive type system
	parsedType := xsd.ParseType(xsdType)
	return parsedType.ToGoType()
}
