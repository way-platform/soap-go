package gen

import (
	"fmt"
	"go/parser"
	"go/token"
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

	// Create schema context for reference resolution
	ctx := newSchemaContext(schema)

	// Create type registry to prevent duplicates
	typeRegistry := newTypeRegistry()

	// Add package declaration
	g.P("package ", packageName)
	g.P()

	// Separate data types from message wrapper types
	dataTypes, messageTypes := categorizeElements(schema.Elements)

	// Collect all required imports by analyzing the types used
	requiredImports := make(map[string]bool)
	for _, element := range append(dataTypes, messageTypes...) {
		collectRequiredImports(element, requiredImports, ctx)
	}

	// Add imports
	for imp := range requiredImports {
		g.Import(imp)
	}

	if len(requiredImports) > 0 {
		g.P()
	}

	// Generate simple type constants first (for enumerations)
	generateSimpleTypeConstants(g, ctx)

	// Generate complex types that are referenced but not top-level elements
	generateComplexTypes(g, ctx)

	// Generate inline complex types first (before elements that use them)
	generateInlineComplexTypes(g, ctx, dataTypes)

	// Generate data types (for proper dependency ordering)
	for _, element := range dataTypes {
		if typeRegistry.shouldGenerate(element) {
			generateStructFromElement(g, element, ctx)
		}
	}

	// Skip message wrapper types for now as they cause duplicates
	// TODO: In the future, we might want to generate these with different names
	// or in a separate package for SOAP message handling

	// Get the generated content
	content, err := g.Content()
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	// Validate the generated Go code
	if err := validateGeneratedCode(content); err != nil {
		return fmt.Errorf("generated invalid Go code: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, content, 0o644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}

// SchemaContext provides context for resolving references within a schema
type SchemaContext struct {
	schema         *xsd.Schema
	elementRefs    map[string]*xsd.Element
	simpleTypes    map[string]*xsd.SimpleType
	complexTypes   map[string]*xsd.ComplexType
	anonymousTypes map[string]bool // Track generated anonymous types
}

// AnonymousTypeRegistry tracks generated anonymous types to prevent conflicts
type AnonymousTypeRegistry struct {
	types map[string]bool
}

func newAnonymousTypeRegistry() *AnonymousTypeRegistry {
	return &AnonymousTypeRegistry{
		types: make(map[string]bool),
	}
}

func (r *AnonymousTypeRegistry) generateTypeName(parentName, fieldName string) string {
	// Use Outer_Inner naming convention
	baseName := toGoName(parentName) + "_" + toGoName(fieldName)
	if !r.types[baseName] {
		r.types[baseName] = true
		return baseName
	}

	// Handle conflicts with numbered suffixes
	for i := 2; ; i++ {
		name := fmt.Sprintf("%s%d", baseName, i)
		if !r.types[name] {
			r.types[name] = true
			return name
		}
	}
}

func newSchemaContext(schema *xsd.Schema) *SchemaContext {
	ctx := &SchemaContext{
		schema:         schema,
		elementRefs:    make(map[string]*xsd.Element),
		simpleTypes:    make(map[string]*xsd.SimpleType),
		complexTypes:   make(map[string]*xsd.ComplexType),
		anonymousTypes: make(map[string]bool),
	}

	// Build reference maps
	for i := range schema.Elements {
		elem := &schema.Elements[i]
		ctx.elementRefs[elem.Name] = elem
	}

	for i := range schema.SimpleTypes {
		simpleType := &schema.SimpleTypes[i]
		ctx.simpleTypes[simpleType.Name] = simpleType
	}

	for i := range schema.ComplexTypes {
		complexType := &schema.ComplexTypes[i]
		ctx.complexTypes[complexType.Name] = complexType
	}

	return ctx
}

func (ctx *SchemaContext) resolveElementRef(ref string) *xsd.Element {
	// Handle namespace prefixes (tp:LoginRequest -> LoginRequest)
	if colonIdx := strings.LastIndex(ref, ":"); colonIdx != -1 {
		ref = ref[colonIdx+1:]
	}
	return ctx.elementRefs[ref]
}

func (ctx *SchemaContext) resolveSimpleType(typeName string) *xsd.SimpleType {
	// Handle namespace prefixes (tp:sessionidType -> sessionidType)
	if colonIdx := strings.LastIndex(typeName, ":"); colonIdx != -1 {
		typeName = typeName[colonIdx+1:]
	}
	return ctx.simpleTypes[typeName]
}

func (ctx *SchemaContext) resolveComplexType(typeName string) *xsd.ComplexType {
	// Handle namespace prefixes (tp:PropertyType -> PropertyType)
	if colonIdx := strings.LastIndex(typeName, ":"); colonIdx != -1 {
		typeName = typeName[colonIdx+1:]
	}
	return ctx.complexTypes[typeName]
}

// TypeRegistry tracks generated types to prevent duplicates
type TypeRegistry struct {
	types map[string]*xsd.Element
}

func newTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types: make(map[string]*xsd.Element),
	}
}

func (r *TypeRegistry) shouldGenerate(element *xsd.Element) bool {
	name := toGoName(element.Name)
	if name == "" {
		return false // Skip elements without valid names
	}

	if existing, exists := r.types[name]; exists {
		// Compare structures to see if they're equivalent
		return !areEquivalentElements(existing, element)
	}
	r.types[name] = element
	return true
}

// areEquivalentElements checks if two elements have the same structure
func areEquivalentElements(a, b *xsd.Element) bool {
	// Simple comparison - could be enhanced for more sophisticated checking
	if a.Name != b.Name {
		return false
	}

	// Check if both have complex types with same structure
	if (a.ComplexType == nil) != (b.ComplexType == nil) {
		return false
	}

	if a.ComplexType != nil && b.ComplexType != nil {
		// Compare sequences
		if (a.ComplexType.Sequence == nil) != (b.ComplexType.Sequence == nil) {
			return false
		}

		if a.ComplexType.Sequence != nil && b.ComplexType.Sequence != nil {
			if len(a.ComplexType.Sequence.Elements) != len(b.ComplexType.Sequence.Elements) {
				return false
			}
		}

		// Compare attributes
		if len(a.ComplexType.Attributes) != len(b.ComplexType.Attributes) {
			return false
		}
	}

	return true
}

// categorizeElements separates data types from message wrapper types
func categorizeElements(elements []xsd.Element) ([]*xsd.Element, []*xsd.Element) {
	dataTypes := make([]*xsd.Element, 0)
	messageTypes := make([]*xsd.Element, 0)

	for i := range elements {
		elem := &elements[i]
		// Only process elements that have complex types or could be generated as structs
		if elem.ComplexType != nil {
			if isMessageWrapper(elem) {
				messageTypes = append(messageTypes, elem)
			} else {
				dataTypes = append(dataTypes, elem)
			}
		}
	}

	return dataTypes, messageTypes
}

// isMessageWrapper identifies SOAP message wrapper elements using generic patterns
func isMessageWrapper(element *xsd.Element) bool {
	// SOAP message wrappers typically have the following characteristics:
	// 1. Contain exactly one element that is a reference (not an inline definition)
	// 2. The referenced element name differs from the wrapper element name
	// 3. Often use camelCase while data types use PascalCase
	// 4. May have "Response" suffix for response wrappers

	if element.ComplexType == nil || element.ComplexType.Sequence == nil {
		return false
	}

	elems := element.ComplexType.Sequence.Elements

	// Must have exactly one element that is a reference
	if len(elems) != 1 || elems[0].Ref == "" {
		return false
	}

	// Extract the referenced element name (remove namespace prefix)
	refName := elems[0].Ref
	if colonIdx := strings.LastIndex(refName, ":"); colonIdx != -1 {
		refName = refName[colonIdx+1:]
	}

	// If the wrapper element name is different from the referenced element name,
	// and the wrapper uses camelCase (starts with lowercase), it's likely a message wrapper
	wrapperName := element.Name
	if len(wrapperName) > 0 && wrapperName[0] >= 'a' && wrapperName[0] <= 'z' {
		// camelCase wrapper name suggests it's a SOAP operation wrapper
		return true
	}

	// Check for common SOAP response wrapper patterns (generic)
	lowerWrapperName := strings.ToLower(wrapperName)
	lowerRefName := strings.ToLower(refName)

	// If wrapper ends with "response" and references a different element, it's likely a response wrapper
	if strings.HasSuffix(lowerWrapperName, "response") && lowerWrapperName != lowerRefName {
		return true
	}

	// If the referenced element name is significantly different from wrapper name,
	// and the wrapper doesn't start with uppercase (PascalCase), treat as wrapper
	if wrapperName != refName && len(wrapperName) > 0 && wrapperName[0] >= 'a' && wrapperName[0] <= 'z' {
		return true
	}

	return false
}

// collectRequiredImports recursively collects all import statements needed for the given element
func collectRequiredImports(element *xsd.Element, imports map[string]bool, ctx *SchemaContext) {
	if element.ComplexType != nil {
		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			for _, fieldElement := range element.ComplexType.Sequence.Elements {
				// Parse the type and check if it requires any imports
				if fieldElement.Type != "" {
					parsedType := xsd.ParseType(fieldElement.Type)
					for _, imp := range parsedType.RequiresImport() {
						imports[imp] = true
					}
				} else if fieldElement.ComplexType != nil {
					// Check if inline complex types need imports
					collectRequiredImportsFromComplexType(fieldElement.ComplexType, imports)
				}
			}
		}

		// Handle attributes
		for _, attr := range element.ComplexType.Attributes {
			if attr.Type != "" {
				parsedType := xsd.ParseType(attr.Type)
				for _, imp := range parsedType.RequiresImport() {
					imports[imp] = true
				}
			}
		}

		// Handle complex content extensions
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, fieldElement := range ext.Sequence.Elements {
					if fieldElement.Type != "" {
						parsedType := xsd.ParseType(fieldElement.Type)
						for _, imp := range parsedType.RequiresImport() {
							imports[imp] = true
						}
					}
				}
			}
		}
	}
}

// collectRequiredImportsFromComplexType checks if a complex type needs any imports
func collectRequiredImportsFromComplexType(complexType *xsd.ComplexType, imports map[string]bool) {
	if complexType.Sequence != nil {
		for _, elem := range complexType.Sequence.Elements {
			if elem.Type != "" {
				parsedType := xsd.ParseType(elem.Type)
				for _, imp := range parsedType.RequiresImport() {
					imports[imp] = true
				}
			}
		}
	}
}

// validateGeneratedCode checks if the generated Go code is valid
func validateGeneratedCode(content []byte) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	return err
}

// generateInlineComplexTypes generates structs for all inline complex types found in elements
func generateInlineComplexTypes(g *codegen.File, ctx *SchemaContext, elements []*xsd.Element) {
	registry := newAnonymousTypeRegistry()
	hasTypes := false

	for _, element := range elements {
		if generated := generateInlineTypesFromElement(g, element, "", ctx, registry); generated && !hasTypes {
			g.P("// Inline complex types")
			g.P()
			hasTypes = true
		}
	}

	if hasTypes {
		g.P()
	}
}

// generateInlineTypesFromElement recursively generates inline complex types from an element
func generateInlineTypesFromElement(g *codegen.File, element *xsd.Element, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) bool {
	generated := false

	if element.ComplexType != nil && element.ComplexType.Sequence != nil {
		for _, field := range element.ComplexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate inline complex type using Outer_Inner naming
				typeName := registry.generateTypeName(element.Name, field.Name)
				generateInlineComplexTypeStruct(g, typeName, field.ComplexType, ctx)
				generated = true

				// Register this type so we can reference it later
				ctx.anonymousTypes[typeName] = true

				// Recursively check for nested inline types
				if generateInlineTypesFromComplexType(g, field.ComplexType, typeName, ctx, registry) {
					generated = true
				}
			}
		}
	}

	return generated
}

// generateInlineTypesFromComplexType recursively generates inline complex types from a complex type
func generateInlineTypesFromComplexType(g *codegen.File, complexType *xsd.ComplexType, parentName string, ctx *SchemaContext, registry *AnonymousTypeRegistry) bool {
	generated := false

	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if field.ComplexType != nil {
				// Generate inline complex type using Outer_Inner naming
				typeName := registry.generateTypeName(parentName, field.Name)
				generateInlineComplexTypeStruct(g, typeName, field.ComplexType, ctx)
				generated = true

				// Register this type so we can reference it later
				ctx.anonymousTypes[typeName] = true

				// Recursively check for nested inline types
				if generateInlineTypesFromComplexType(g, field.ComplexType, typeName, ctx, registry) {
					generated = true
				}
			}
		}
	}

	return generated
}

// generateInlineComplexTypeStruct generates a struct for an inline complex type
func generateInlineComplexTypeStruct(g *codegen.File, typeName string, complexType *xsd.ComplexType, ctx *SchemaContext) {
	// Add comment
	g.P("// ", typeName, " represents an inline complex type")

	// Start struct declaration
	g.P("type ", typeName, " struct {")

	hasFields := false

	// Generate fields from the sequence
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypes(g, &field, ctx) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeField(g, &attr, ctx) {
			hasFields = true
		}
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined")
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructFromElement generates a Go struct from an XSD element
func generateStructFromElement(g *codegen.File, element *xsd.Element, ctx *SchemaContext) {
	structName := toGoName(element.Name)
	if structName == "" {
		return // Skip elements without valid names
	}

	// Add comment
	g.P("// ", structName, " represents the ", element.Name, " element")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	// Track if we've added any fields
	hasFields := false

	// Generate fields from the complex type
	if element.ComplexType != nil {
		// Handle sequence elements
		if element.ComplexType.Sequence != nil {
			for _, field := range element.ComplexType.Sequence.Elements {
				if generateStructFieldWithInlineTypes(g, &field, ctx) {
					hasFields = true
				}
			}
		}

		// Handle attributes
		for _, attr := range element.ComplexType.Attributes {
			if generateAttributeField(g, &attr, ctx) {
				hasFields = true
			}
		}

		// Handle complex content extensions
		if element.ComplexType.ComplexContent != nil && element.ComplexType.ComplexContent.Extension != nil {
			ext := element.ComplexType.ComplexContent.Extension
			if ext.Sequence != nil {
				for _, field := range ext.Sequence.Elements {
					if generateStructFieldWithInlineTypes(g, &field, ctx) {
						hasFields = true
					}
				}
			}
			// Handle extension attributes
			for _, attr := range ext.Attributes {
				if generateAttributeField(g, &attr, ctx) {
					hasFields = true
				}
			}
		}
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined for this element")
	}

	// Close struct
	g.P("}")
	g.P()
}

// generateStructFieldWithInlineTypes generates a Go struct field with support for inline complex types
func generateStructFieldWithInlineTypes(g *codegen.File, element *xsd.Element, ctx *SchemaContext) bool {
	// Handle element references
	if element.Ref != "" {
		// Resolve the reference
		referencedElement := ctx.resolveElementRef(element.Ref)
		if referencedElement != nil {
			// Use the referenced element's name for the field
			fieldName := toGoName(referencedElement.Name)
			goType := toGoName(referencedElement.Name) // Reference the generated type
			xmlName := referencedElement.Name

			if fieldName != "" {
				// Handle optional elements
				if element.MinOccurs == "0" {
					goType = "*" + goType
				}
				// Handle multiple elements
				if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
					// For []byte (raw XML capture), don't create [][]byte - keep as []byte
					// For most other types, create slice of the type
					if goType != "[]byte" && goType != "*[]byte" {
						goType = "[]" + strings.TrimPrefix(goType, "*")
					}
				}

				g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, "\"`")
				return true
			}
		}
		return false
	}

	// Skip elements without names
	if element.Name == "" {
		return false
	}

	fieldName := toGoName(element.Name)
	if fieldName == "" {
		return false
	}

	// Determine the Go type
	var goType string
	if element.Type != "" {
		goType = mapXSDTypeToGoWithContext(element.Type, ctx)
	} else if element.ComplexType != nil {
		// For inline complex types without explicit type generation, use []byte to capture raw XML
		// This allows consumers to parse the XML content later when stronger typing is implemented
		goType = "[]byte"
	} else {
		goType = "string" // fallback
	}

	xmlName := element.Name

	// Handle optional elements
	if element.MinOccurs == "0" {
		if !strings.HasPrefix(goType, "*") && !strings.HasPrefix(goType, "[]") {
			goType = "*" + goType
		}
	}

	// Handle multiple elements
	if element.MaxOccurs == "unbounded" || (element.MaxOccurs != "" && element.MaxOccurs != "1") {
		// For []byte (raw XML capture), don't create [][]byte - keep as []byte to capture all XML content
		if goType != "[]byte" && goType != "*[]byte" {
			goType = "[]" + strings.TrimPrefix(goType, "*")
		}
	}

	// Generate the field with XML tag
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, "\"`")
	return true
}

// generateAttributeField generates a Go struct field from an XSD attribute
func generateAttributeField(g *codegen.File, attr *xsd.Attribute, ctx *SchemaContext) bool {
	if attr.Name == "" {
		return false
	}

	fieldName := toGoName(attr.Name)
	if fieldName == "" {
		return false
	}

	goType := mapXSDTypeToGoWithContext(attr.Type, ctx)
	xmlName := attr.Name

	// Handle optional attributes
	if attr.Use != "required" {
		if !strings.HasPrefix(goType, "*") {
			goType = "*" + goType
		}
	}

	// Generate the field with XML attribute tag
	g.P("\t", fieldName, " ", goType, " `xml:\"", xmlName, ",attr\"`")
	return true
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

// mapXSDTypeToGoWithContext maps XSD types to Go types using schema context for better resolution
func mapXSDTypeToGoWithContext(xsdType string, ctx *SchemaContext) string {
	if xsdType == "" {
		return "[]byte" // fallback for empty types - capture raw XML
	}

	// First try to resolve as a simple type in the schema
	if simpleType := ctx.resolveSimpleType(xsdType); simpleType != nil {
		return resolveSimpleTypeToGo(simpleType)
	}

	// Then try to resolve as a complex type in the schema
	if complexType := ctx.resolveComplexType(xsdType); complexType != nil {
		// For named complex types, generate a Go type name
		return toGoName(extractLocalName(xsdType))
	}

	// Check if this is a custom type (contains namespace prefix or ends with "Type")
	localName := extractLocalName(xsdType)
	if isCustomTypeName(localName) {
		// For custom types not defined in this schema, use a reasonable Go type
		return inferGoTypeFromCustomTypeName(localName)
	}

	// Try standard XSD type parsing
	parsedType := xsd.ParseType(xsdType)
	if !parsedType.IsCustomType() {
		return parsedType.ToGoType()
	}

	// For truly unknown/custom types, use []byte to capture raw XML
	return "[]byte"
}

// isCustomTypeName checks if a type name looks like a custom type
func isCustomTypeName(typeName string) bool {
	// Check if it ends with "Type" suffix (common pattern)
	if len(typeName) > 4 && strings.HasSuffix(typeName, "Type") {
		return true
	}

	// Check if it's not a standard XSD type
	parsedType := xsd.ParseType(typeName)
	return parsedType.IsCustomType()
}

// inferGoTypeFromCustomTypeName attempts to infer appropriate Go type from custom type name
func inferGoTypeFromCustomTypeName(typeName string) string {
	// Handle common patterns based on naming conventions
	name := strings.ToLower(typeName)

	// ID types are typically numeric
	if strings.Contains(name, "id") && strings.HasSuffix(name, "type") {
		return "int64"
	}

	// Timestamp types are typically strings (custom format)
	if strings.Contains(name, "timestamp") {
		return "string"
	}

	// Version types are typically numeric
	if strings.Contains(name, "version") {
		return "int64"
	}

	// Limit, offset, size types are typically numeric
	if strings.Contains(name, "limit") || strings.Contains(name, "offset") || strings.Contains(name, "size") {
		return "int64"
	}

	// Session types are typically strings
	if strings.Contains(name, "session") {
		return "string"
	}

	// For other custom types ending in "Type", assume string (safest default)
	if strings.HasSuffix(name, "type") {
		return "string"
	}

	// Default to generating a proper Go type name for complex types
	return xsd.ToGoTypeName(typeName)
}

// resolveSimpleTypeToGo converts a simple type definition to a Go type
func resolveSimpleTypeToGo(simpleType *xsd.SimpleType) string {
	if simpleType.Restriction != nil {
		// Get the base type and handle restrictions
		baseType := xsd.ParseType(simpleType.Restriction.Base)

		// Check if it's an enumeration
		if len(simpleType.Restriction.Enumerations) > 0 {
			// For enumerations, we'll use string for now
			// In the future, we could generate const declarations
			return baseType.ToGoType()
		}

		return baseType.ToGoType()
	}

	// Default to string if we can't determine the type
	return "string"
}

// extractLocalName removes namespace prefix from a type name
func extractLocalName(typeName string) string {
	if colonIdx := strings.LastIndex(typeName, ":"); colonIdx != -1 {
		return typeName[colonIdx+1:]
	}
	return typeName
}

// generateSimpleTypeConstants generates Go constants for enumeration simple types
func generateSimpleTypeConstants(g *codegen.File, ctx *SchemaContext) {
	hasEnums := false

	for _, simpleType := range ctx.simpleTypes {
		if simpleType.Restriction != nil && len(simpleType.Restriction.Enumerations) > 0 {
			if !hasEnums {
				g.P("// Enumeration constants")
				g.P()
				hasEnums = true
			}

			generateEnumConstants(g, simpleType)
		}
	}

	if hasEnums {
		g.P()
	}
}

// generateEnumConstants generates Go constants for a single enumeration type
func generateEnumConstants(g *codegen.File, simpleType *xsd.SimpleType) {
	typeName := toGoName(simpleType.Name)

	g.P("// ", typeName, " enumeration values")
	g.P("const (")

	for _, enum := range simpleType.Restriction.Enumerations {
		constName := typeName + toGoName(enum.Value)
		g.P("\t", constName, " = \"", enum.Value, "\"")
	}

	g.P(")")
	g.P()
}

// generateComplexTypes generates Go structs for named complex types
func generateComplexTypes(g *codegen.File, ctx *SchemaContext) {
	hasTypes := false

	for _, complexType := range ctx.complexTypes {
		if !hasTypes {
			g.P("// Complex types")
			g.P()
			hasTypes = true
		}

		generateStructFromComplexType(g, complexType, ctx)
	}

	if hasTypes {
		g.P()
	}
}

// generateStructFromComplexType generates a Go struct from an XSD complexType definition
func generateStructFromComplexType(g *codegen.File, complexType *xsd.ComplexType, ctx *SchemaContext) {
	structName := toGoName(complexType.Name)
	if structName == "" {
		return
	}

	// Add comment
	g.P("// ", structName, " represents the ", complexType.Name, " complex type")

	// Start struct declaration
	g.P("type ", structName, " struct {")

	hasFields := false

	// Generate fields from the sequence
	if complexType.Sequence != nil {
		for _, field := range complexType.Sequence.Elements {
			if generateStructFieldWithInlineTypes(g, &field, ctx) {
				hasFields = true
			}
		}
	}

	// Handle attributes
	for _, attr := range complexType.Attributes {
		if generateAttributeField(g, &attr, ctx) {
			hasFields = true
		}
	}

	// Handle complex content extensions
	if complexType.ComplexContent != nil && complexType.ComplexContent.Extension != nil {
		ext := complexType.ComplexContent.Extension
		if ext.Sequence != nil {
			for _, field := range ext.Sequence.Elements {
				if generateStructFieldWithInlineTypes(g, &field, ctx) {
					hasFields = true
				}
			}
		}
		// Handle extension attributes
		for _, attr := range ext.Attributes {
			if generateAttributeField(g, &attr, ctx) {
				hasFields = true
			}
		}
	}

	// If no fields were generated, add a placeholder comment
	if !hasFields {
		g.P("\t// No fields defined")
	}

	// Close struct
	g.P("}")
	g.P()
}
