package soapgen

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"

	"github.com/way-platform/soap-go/xsd"
)

// SchemaContext provides context for resolving references within a schema
type SchemaContext struct {
	schema         *xsd.Schema
	elementRefs    map[string]*xsd.Element
	simpleTypes    map[string]*xsd.SimpleType
	complexTypes   map[string]*xsd.ComplexType
	anonymousTypes map[string]bool           // Track generated anonymous types
	inlineEnums    map[string]InlineEnumInfo // Track inline enum types
	generator      *Generator                // Reference to generator for operation element detection
}

// InlineEnumInfo holds information about an inline enum type
type InlineEnumInfo struct {
	TypeName      string
	SimpleType    *xsd.SimpleType
	ParentName    string
	FieldName     string
	EnumSignature string // Hash of enum values for deduplication
	Generated     bool
}

// AnonymousTypeRegistry tracks generated anonymous types to prevent conflicts
type AnonymousTypeRegistry struct {
	types map[string]bool
}

// FieldRegistry tracks field names within a struct to prevent duplicates
type FieldRegistry struct {
	fields map[string]FieldInfo // field name -> field info
}

// FieldInfo holds information about a generated field
type FieldInfo struct {
	xmlName     string
	isAttribute bool
	goFieldName string
}

// TypeContext represents the context in which a type is being generated
type TypeContext int

const (
	DataElementContext TypeContext = iota
	SOAPWrapperContext
	OperationWrapperContext
)

// TypeInfo holds information about a generated type
type TypeInfo struct {
	element    *xsd.Element
	xmlName    string
	goTypeName string
	context    TypeContext
}

// TypeRegistry tracks generated types to prevent duplicates and handle collisions
type TypeRegistry struct {
	types    map[string]*TypeInfo   // Go type name -> TypeInfo
	xmlNames map[string][]*TypeInfo // XML name -> list of TypeInfo (for collision detection)
}

func newAnonymousTypeRegistry() *AnonymousTypeRegistry {
	return &AnonymousTypeRegistry{
		types: make(map[string]bool),
	}
}

func newFieldRegistry() *FieldRegistry {
	return &FieldRegistry{
		fields: make(map[string]FieldInfo),
	}
}

// generateUniqueFieldName generates a unique field name avoiding collisions
func (r *FieldRegistry) generateUniqueFieldName(xmlName string, isAttribute bool) string {
	baseName := toGoName(xmlName)
	if baseName == "" {
		return ""
	}

	// Check if this exact combination already exists
	if _, exists := r.fields[baseName]; exists {
		// Always generate unique names to avoid field redeclaration errors
		// TODO: Implement proper field combination logic for multiple references to same element

		// Collision detected - generate unique name
		var candidateName string
		if isAttribute {
			candidateName = baseName + "Attr"
		} else {
			candidateName = baseName + "Elem"
		}

		// If that's still taken, use numbered suffix
		if r.hasFieldName(candidateName) {
			candidateName = r.generateNumberedFieldName(baseName, isAttribute)
		}

		// Register the new field
		r.fields[candidateName] = FieldInfo{
			xmlName:     xmlName,
			isAttribute: isAttribute,
			goFieldName: candidateName,
		}
		return candidateName
	}

	// No collision, use base name
	r.fields[baseName] = FieldInfo{
		xmlName:     xmlName,
		isAttribute: isAttribute,
		goFieldName: baseName,
	}
	return baseName
}

// hasFieldName checks if a field name is already used
func (r *FieldRegistry) hasFieldName(fieldName string) bool {
	_, exists := r.fields[fieldName]
	return exists
}

// generateNumberedFieldName generates a unique field name with numbered suffix
func (r *FieldRegistry) generateNumberedFieldName(baseName string, isAttribute bool) string {
	suffix := "Elem"
	if isAttribute {
		suffix = "Attr"
	}

	for i := 1; ; i++ {
		candidateName := fmt.Sprintf("%s%s%d", baseName, suffix, i)
		if !r.hasFieldName(candidateName) {
			return candidateName
		}
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

func newSchemaContext(schema *xsd.Schema, generator *Generator) *SchemaContext {
	ctx := &SchemaContext{
		schema:         schema,
		elementRefs:    make(map[string]*xsd.Element),
		simpleTypes:    make(map[string]*xsd.SimpleType),
		complexTypes:   make(map[string]*xsd.ComplexType),
		anonymousTypes: make(map[string]bool),
		inlineEnums:    make(map[string]InlineEnumInfo),
		generator:      generator,
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

// getInlineEnumTypeName returns the type name for an inline enum if it exists
func (ctx *SchemaContext) getInlineEnumTypeName(parentName, fieldName string) string {
	typeName := toGoName(parentName) + "_" + toGoName(fieldName)
	if enumInfo, exists := ctx.inlineEnums[typeName]; exists {
		return enumInfo.TypeName
	}
	return ""
}

// collectInlineEnums traverses all elements and collects inline enum types
func (ctx *SchemaContext) collectInlineEnums(elements []*xsd.Element) {
	for _, element := range elements {
		ctx.collectInlineEnumsFromElement(element, element.Name)
	}
}

// collectInlineEnumsFromElement recursively collects inline enums from an element
func (ctx *SchemaContext) collectInlineEnumsFromElement(element *xsd.Element, parentName string) {
	// Check if this element has an inline simple type with enumerations
	if element.SimpleType != nil && ctx.hasEnumerations(element.SimpleType) {
		ctx.registerInlineEnum(parentName, element.Name, element.SimpleType)
	}

	// Recursively check complex type elements
	if element.ComplexType != nil {
		ctx.collectInlineEnumsFromComplexType(element.ComplexType, parentName)
	}
}

// collectInlineEnumsFromComplexType recursively collects inline enums from a complex type
func (ctx *SchemaContext) collectInlineEnumsFromComplexType(complexType *xsd.ComplexType, parentName string) {
	// Check sequence elements
	if complexType.Sequence != nil {
		for _, elem := range complexType.Sequence.Elements {
			ctx.collectInlineEnumsFromElement(&elem, parentName)
		}
	}

	// Check choice elements
	if complexType.Choice != nil {
		for _, elem := range complexType.Choice.Elements {
			ctx.collectInlineEnumsFromElement(&elem, parentName)
		}
	}

	// Check all elements
	if complexType.All != nil {
		for _, elem := range complexType.All.Elements {
			ctx.collectInlineEnumsFromElement(&elem, parentName)
		}
	}

	// Check attributes with inline simple types
	for _, attr := range complexType.Attributes {
		if attr.SimpleType != nil && ctx.hasEnumerations(attr.SimpleType) {
			ctx.registerInlineEnum(parentName, attr.Name, attr.SimpleType)
		}
	}
}

// hasEnumerations checks if a simple type has enumeration restrictions
func (ctx *SchemaContext) hasEnumerations(simpleType *xsd.SimpleType) bool {
	return simpleType.Restriction != nil && len(simpleType.Restriction.Enumerations) > 0
}

// registerInlineEnum registers an inline enum type with deduplication
func (ctx *SchemaContext) registerInlineEnum(parentName, fieldName string, simpleType *xsd.SimpleType) {
	typeName := toGoName(parentName) + "_" + toGoName(fieldName)

	// Generate enum signature for deduplication
	signature := ctx.generateEnumSignature(simpleType)

	// Check if we already have an enum with the same signature
	existingTypeName := ctx.findExistingEnumBySignature(signature)
	if existingTypeName != "" {
		// Reuse existing enum type
		ctx.inlineEnums[typeName] = InlineEnumInfo{
			TypeName:      existingTypeName,
			SimpleType:    simpleType,
			ParentName:    parentName,
			FieldName:     fieldName,
			EnumSignature: signature,
			Generated:     false, // Will be marked true when the original is generated
		}
		return
	}

	// Register new enum type
	ctx.inlineEnums[typeName] = InlineEnumInfo{
		TypeName:      typeName,
		SimpleType:    simpleType,
		ParentName:    parentName,
		FieldName:     fieldName,
		EnumSignature: signature,
		Generated:     false,
	}
}

// generateEnumSignature creates a signature hash for enum values to enable deduplication
func (ctx *SchemaContext) generateEnumSignature(simpleType *xsd.SimpleType) string {
	if simpleType.Restriction == nil {
		return ""
	}

	var values []string
	for _, enum := range simpleType.Restriction.Enumerations {
		values = append(values, enum.Value)
	}

	// Sort values for consistent signature
	sort.Strings(values)

	// Create hash
	h := sha256.New()
	for _, value := range values {
		h.Write([]byte(value + "|"))
	}

	return fmt.Sprintf("%x", h.Sum(nil))[:16] // Use first 16 chars of hash
}

// findExistingEnumBySignature finds an existing enum type with the same signature
func (ctx *SchemaContext) findExistingEnumBySignature(signature string) string {
	for _, enumInfo := range ctx.inlineEnums {
		if enumInfo.EnumSignature == signature && enumInfo.TypeName != "" {
			return enumInfo.TypeName
		}
	}
	return ""
}

func newTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types:    make(map[string]*TypeInfo),
		xmlNames: make(map[string][]*TypeInfo),
	}
}

// shouldGenerateWithContext checks if a type should be generated with the given context
func (r *TypeRegistry) shouldGenerateWithContext(element *xsd.Element, context TypeContext) bool {
	baseName := toGoName(element.Name)
	if baseName == "" {
		return false // Skip elements without valid names
	}

	// Generate unique type name considering context and collisions
	typeName := r.generateUniqueTypeName(baseName, element.Name, context)

	if existing, exists := r.types[typeName]; exists {
		// Compare structures to see if they're equivalent
		return !areEquivalentElements(existing.element, element)
	}

	// Register the new type
	typeInfo := &TypeInfo{
		element:    element,
		xmlName:    element.Name,
		goTypeName: typeName,
		context:    context,
	}
	r.types[typeName] = typeInfo
	r.xmlNames[element.Name] = append(r.xmlNames[element.Name], typeInfo)
	return true
}

// generateUniqueTypeName generates a unique Go type name considering context and collisions
func (r *TypeRegistry) generateUniqueTypeName(baseName, xmlName string, context TypeContext) string {
	// Check if there are existing types with the same XML name but different case
	if existingTypes := r.xmlNames[xmlName]; len(existingTypes) > 0 {
		// If we already have a type for this exact XML name, use the existing Go type name
		for _, existing := range existingTypes {
			if existing.xmlName == xmlName && existing.context == context {
				return existing.goTypeName
			}
		}
	}

	// For SOAPWrapperContext, always use wrapper suffix for consistency
	if context == SOAPWrapperContext {
		candidateName := baseName + "Wrapper"
		if !r.hasGoTypeName(candidateName) {
			return candidateName
		}
		return r.generateNumberedName(baseName + "Wrapper")
	}

	// Check for Go type name collisions (simplified approach)
	if r.hasGoTypeName(baseName) {
		// Collision detected, generate unique name with numbered suffix
		return r.generateNumberedName(baseName)
	}

	// No collision, use the base name
	return baseName
}

// hasGoTypeName checks if a Go type name is already used
func (r *TypeRegistry) hasGoTypeName(goName string) bool {
	_, exists := r.types[goName]
	return exists
}

// generateNumberedName generates a unique name with numbered suffix
func (r *TypeRegistry) generateNumberedName(baseName string) string {
	if !r.hasGoTypeName(baseName) {
		return baseName
	}

	for i := 2; ; i++ {
		candidateName := fmt.Sprintf("%s%d", baseName, i)
		if !r.hasGoTypeName(candidateName) {
			return candidateName
		}
	}
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
		} else if elem.SimpleType != nil {
			// Elements with inline simple types should also be processed as data types
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

// Import collection is now handled automatically via QualifiedGoIdent calls in codegen
