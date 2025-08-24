package wsdlgen

import (
	"fmt"
	"strings"

	"github.com/way-platform/soap-go/xsd"
)

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

// TypeRegistry tracks generated types to prevent duplicates
type TypeRegistry struct {
	types map[string]*xsd.Element
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
		// Process all elements that could be generated as structs
		if elem.ComplexType != nil {
			if isMessageWrapper(elem) {
				messageTypes = append(messageTypes, elem)
			} else {
				dataTypes = append(dataTypes, elem)
			}
		} else if elem.Type != "" {
			// Elements with simple types should also be processed as data types
			dataTypes = append(dataTypes, elem)
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
