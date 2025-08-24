package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl11"
	"github.com/way-platform/soap-go/xsd10"
)

func main() {
	inputFile := flag.String("i", "", "input WSDL file")
	outputDir := flag.String("d", "", "output directory")
	packageName := flag.String("p", "", "Go package name")
	flag.Parse()
	if *inputFile == "" {
		log.Fatal("input file is required")
	}
	if *outputDir == "" {
		log.Fatal("output directory is required")
	}
	if *packageName == "" {
		log.Fatal("package name is required")
	}
	if err := run(config{
		inputFile:   *inputFile,
		outputDir:   *outputDir,
		packageName: *packageName,
	}); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	inputFile   string
	outputDir   string
	packageName string
}

func run(cfg config) error {
	// Parse the WSDL file
	defs, err := wsdl11.ParseFromFile(cfg.inputFile)
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
func generateTypesFile(schema *xsd10.Schema, packageName, outputDir, filename string) error {
	outputPath := filepath.Join(outputDir, filename)
	g := codegen.NewFile(outputPath)

	// Add package declaration
	g.P("package ", packageName)
	g.P()

	// Note: We always need encoding/xml for struct tags, but Go will complain
	// if we import it without using it in the code. We'll add it as a comment
	// for now since the XML tags don't count as "usage" for the compiler.
	g.P("// This package uses XML tags for marshaling/unmarshaling")
	g.P("// import \"encoding/xml\" // Uncomment if you need to marshal/unmarshal XML")
	g.P()

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

// generateStructFromElement generates a Go struct from an XSD element
func generateStructFromElement(g *codegen.File, element *xsd10.Element) {
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
func generateStructField(g *codegen.File, element *xsd10.Element) {
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

// mapXSDTypeToGo maps XSD types to Go types
func mapXSDTypeToGo(xsdType string) string {
	// Remove namespace prefix if present
	parts := strings.Split(xsdType, ":")
	if len(parts) == 2 {
		xsdType = parts[1]
	}
	switch xsdType {
	case "string":
		return "string"
	case "int", "integer":
		return "int"
	case "long":
		return "int64"
	case "unsignedLong":
		return "uint64"
	case "short":
		return "int16"
	case "unsignedShort":
		return "uint16"
	case "byte":
		return "int8"
	case "unsignedByte":
		return "uint8"
	case "decimal", "double":
		return "float64"
	case "float":
		return "float32"
	case "boolean":
		return "bool"
	case "dateTime":
		return "string" // TODO: Consider using time.Time
	case "date":
		return "string" // TODO: Consider using time.Time
	case "time":
		return "string" // TODO: Consider using time.Time
	case "duration":
		return "string" // TODO: Consider using time.Duration
	case "base64Binary":
		return "[]byte"
	case "hexBinary":
		return "[]byte"
	case "anyURI":
		return "string"
	case "QName":
		return "string"
	default:
		return "string" // Default to string for unknown types
	}
}
