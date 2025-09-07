package docgen

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

// Generator generates Markdown documentation from WSDL definitions.
type Generator struct {
	definitions       *wsdl.Definitions
	output            *codegen.File
	filename          string
	inlineSimpleTypes map[string]*inlineSimpleTypeInfo // Track inline simple types by element path
}

// inlineSimpleTypeInfo holds information about an inline simple type
type inlineSimpleTypeInfo struct {
	ElementPath   string
	ElementName   string
	SimpleType    *xsd.SimpleType
	Documentation string
}

// NewGenerator creates a new [Generator] for the given WSDL definitions.
func NewGenerator(filename string, definitions *wsdl.Definitions) *Generator {
	return &Generator{
		definitions:       definitions,
		output:            codegen.NewFile(filename, ""),
		filename:          filename,
		inlineSimpleTypes: make(map[string]*inlineSimpleTypeInfo),
	}
}

// File returns the generated Markdown file.
func (g *Generator) File() *codegen.File {
	return g.output
}

// Generate runs the generator.
func (g *Generator) Generate() error {
	return g.generateMarkdown()
}

// generateMarkdown generates markdown documentation for a WSDL file
func (g *Generator) generateMarkdown() error {
	doc := g.definitions

	// Title - use service name or filename as fallback if doc.Name is empty
	title := g.getDocumentTitle()
	g.output.P("# ", title)

	// Add namespace and endpoint as table
	if doc.TargetNamespace != "" || g.getServiceEndpoint() != "" {
		g.output.P()
		g.output.P("| | |")
		g.output.P("|---|---|")
		if doc.TargetNamespace != "" {
			g.output.P("| **Namespace** | `", doc.TargetNamespace, "` |")
		}
		if endpoint := g.getServiceEndpoint(); endpoint != "" {
			g.output.P("| **Endpoint** | `", endpoint, "` |")
		}
		g.output.P()
	}

	// Add overview section with operations summary
	if err := g.generateOverview(); err != nil {
		return err
	}

	// Build schema map for element lookups
	schemaMap := g.buildSchemaMap()

	// Generate operations section
	g.output.P("## Operations")
	g.output.P()

	// Build initial custom types map for hyperlinking
	customTypesMap := g.buildCustomTypesMap()

	// Generate documentation for each service
	for _, service := range doc.Service {
		if err := g.generateServiceDoc(&service, schemaMap, customTypesMap); err != nil {
			return err
		}
	}

	// Generate custom types section
	if err := g.generateCustomTypesSection(); err != nil {
		return err
	}

	return nil
}

// getDocumentTitle returns the best available title for the document
func (g *Generator) getDocumentTitle() string {
	doc := g.definitions

	// First try the document name
	if doc.Name != "" {
		return doc.Name + " API Documentation"
	}

	// Fall back to the first service name
	if len(doc.Service) > 0 && doc.Service[0].Name != "" {
		return doc.Service[0].Name + " API Documentation"
	}

	// Extract from filename as last resort
	filename := g.output.Filename()
	if filename != "" {
		// Remove .md extension and capitalize
		name := strings.TrimSuffix(filename, ".md")
		caser := cases.Title(language.English)
		return caser.String(name) + " API Documentation"
	}

	return "API Documentation"
}

// getServiceEndpoint extracts the service endpoint URL from bindings
func (g *Generator) getServiceEndpoint() string {
	doc := g.definitions

	if len(doc.Service) > 0 {
		service := &doc.Service[0]
		for _, port := range service.Ports {
			if port.SOAP11Address != nil && port.SOAP11Address.Location != "" {
				return port.SOAP11Address.Location
			}
			if port.SOAP12Address != nil && port.SOAP12Address.Location != "" {
				return port.SOAP12Address.Location
			}
			if port.HTTPAddress != nil && port.HTTPAddress.Location != "" {
				return port.HTTPAddress.Location
			}
		}
	}

	return ""
}

// generateOverview creates an overview section with operations summary
func (g *Generator) generateOverview() error {
	doc := g.definitions

	// Collect all operations from all services
	var allOperations []operationInfo

	for _, service := range doc.Service {
		// Find the corresponding PortType for this service
		portType := g.findPortTypeForService(&service)
		if portType != nil {
			for _, operation := range portType.Operations {
				info := operationInfo{
					Name:        operation.Name,
					Description: normalizeDocumentation(operation.Documentation),
					ServiceName: service.Name,
				}
				allOperations = append(allOperations, info)
			}
		}
	}

	if len(allOperations) == 0 {
		return nil
	}

	// Add service description if available and single service
	if len(doc.Service) == 1 && doc.Service[0].Documentation != "" {
		g.output.P("## Overview")
		g.output.P()
		g.output.P(normalizeDocumentation(doc.Service[0].Documentation))
		g.output.P()
	}

	// Generate operations summary table
	g.output.P("## Available Operations")
	g.output.P()

	for _, op := range allOperations {
		desc := op.Description
		// Create GitHub-compatible anchor link to the operation section
		anchor := strings.ToLower(strings.ReplaceAll(op.Name, " ", "-"))
		if desc == "" {
			// Just show the operation name without description if none available
			g.output.P("- **[", op.Name, "](#", anchor, ")**")
		} else {
			g.output.P("- **[", op.Name, "](#", anchor, ")** - ", desc)
		}
	}
	g.output.P()

	return nil
}

type operationInfo struct {
	Name        string
	Description string
	ServiceName string
}

// findPortTypeForService finds the PortType associated with a service
func (g *Generator) findPortTypeForService(service *wsdl.Service) *wsdl.PortType {
	doc := g.definitions

	for _, binding := range doc.Binding {
		for _, port := range service.Ports {
			if strings.Contains(port.Binding, binding.Name) {
				// Find the PortType referenced by this binding
				// Remove any namespace prefix (e.g., "tns:", "custom:", etc.)
				typeName := binding.Type
				if colonIndex := strings.Index(typeName, ":"); colonIndex >= 0 {
					typeName = typeName[colonIndex+1:]
				}
				for i := range doc.PortType {
					if doc.PortType[i].Name == typeName {
						return &doc.PortType[i]
					}
				}
				break
			}
		}
	}
	return nil
}

// buildSchemaMap creates a map of element names to their definitions for easy lookup
func (g *Generator) buildSchemaMap() map[string]*xsd.Element {
	schemaMap := make(map[string]*xsd.Element)
	doc := g.definitions

	if doc.Types != nil {
		for _, schema := range doc.Types.Schemas {
			for i := range schema.Elements {
				element := &schema.Elements[i]
				schemaMap[element.Name] = element
			}
		}
	}

	return schemaMap
}

// buildCustomTypesMap creates a map of custom type names for hyperlinking
func (g *Generator) buildCustomTypesMap() map[string]bool {
	customTypes := make(map[string]bool)
	doc := g.definitions

	if doc.Types == nil {
		return customTypes
	}

	// Collect all custom types from all schemas
	for _, schema := range doc.Types.Schemas {
		// Collect simple types
		for i := range schema.SimpleTypes {
			simpleType := &schema.SimpleTypes[i]
			if simpleType.Name != "" {
				customTypes[simpleType.Name] = true
			}
		}

		// Collect complex types
		for i := range schema.ComplexTypes {
			complexType := &schema.ComplexTypes[i]
			if complexType.Name != "" {
				customTypes[complexType.Name] = true
			}
		}
	}

	return customTypes
}

// normalizeDocumentation normalizes whitespace in documentation strings
// by trimming leading/trailing whitespace and replacing any sequence of
// whitespace characters (including newlines) with a single space
func normalizeDocumentation(doc string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(doc), " "))
}

// generateServiceDoc generates documentation for a single service
func (g *Generator) generateServiceDoc(service *wsdl.Service, schemaMap map[string]*xsd.Element, customTypesMap map[string]bool) error {
	// Find the corresponding PortType for this service
	portType := g.findPortTypeForService(service)

	if portType == nil {
		g.output.P("*No operations found for this service.*")
		g.output.P()
		return nil
	}

	// Generate documentation for each operation
	for _, operation := range portType.Operations {
		if err := g.generateOperationDoc(&operation, service, schemaMap, customTypesMap); err != nil {
			return err
		}
	}

	return nil
}

// generateOperationDoc generates documentation for a single operation
func (g *Generator) generateOperationDoc(operation *wsdl.Operation, service *wsdl.Service, schemaMap map[string]*xsd.Element, customTypesMap map[string]bool) error {
	// Create GitHub-compatible anchor for the operation
	g.output.P("### ", operation.Name)
	g.output.P()

	// Add operation description if available
	if operation.Documentation != "" {
		g.output.P("> ", normalizeDocumentation(operation.Documentation))
		g.output.P()
	}

	// Add SOAP action and endpoint info
	soapAction := g.getSOAPActionForOperation(operation.Name, service)
	if soapAction != "" {
		g.output.P("**SOAP Action:** `", soapAction, "`")
		g.output.P()
	}

	// Generate request documentation
	if operation.Input != nil {
		if err := g.generateMessageDoc("Request", operation.Input.Message, schemaMap, customTypesMap); err != nil {
			return err
		}
	}

	// Generate response documentation
	if operation.Output != nil {
		if err := g.generateMessageDoc("Response", operation.Output.Message, schemaMap, customTypesMap); err != nil {
			return err
		}
	}

	// Generate fault documentation
	if len(operation.Faults) > 0 {
		g.output.P("#### Faults")
		g.output.P()
		for _, fault := range operation.Faults {
			g.output.P("- **", fault.Name, "**: ", fault.Message)
		}
		g.output.P()
	}

	g.output.P()
	return nil
}

// getSOAPActionForOperation extracts the SOAP action for a specific operation
func (g *Generator) getSOAPActionForOperation(operationName string, service *wsdl.Service) string {
	doc := g.definitions

	// Find the binding for this service
	for _, binding := range doc.Binding {
		for _, port := range service.Ports {
			if strings.Contains(port.Binding, binding.Name) {
				// Look for the operation in this binding
				for _, bindingOp := range binding.BindingOperations {
					if bindingOp.Name == operationName {
						if bindingOp.SOAP11Operation != nil {
							return bindingOp.SOAP11Operation.SOAPAction
						}
						if bindingOp.SOAP12Operation != nil {
							return bindingOp.SOAP12Operation.SOAPAction
						}
					}
				}
			}
		}
	}

	return ""
}

// generateMessageDoc generates documentation for a request or response message
func (g *Generator) generateMessageDoc(messageType, messageName string, schemaMap map[string]*xsd.Element, customTypesMap map[string]bool) error {
	doc := g.definitions

	g.output.P("#### ", messageType)
	g.output.P()

	// Find the message definition
	var message *wsdl.Message
	// Remove any namespace prefix from message name
	cleanMessageName := messageName
	if colonIndex := strings.Index(cleanMessageName, ":"); colonIndex >= 0 {
		cleanMessageName = cleanMessageName[colonIndex+1:]
	}
	for i := range doc.Messages {
		if doc.Messages[i].Name == cleanMessageName {
			message = &doc.Messages[i]
			break
		}
	}

	if message == nil {
		g.output.P("*Message definition not found.*")
		g.output.P()
		return nil
	}

	g.output.P("**Message:** `", cleanMessageName, "`")
	g.output.P()

	// Collect all fields from all parts
	var fields []fieldInfo
	for _, part := range message.Parts {
		if part.Element != "" {
			// Remove any namespace prefix from element name
			elementName := part.Element
			if colonIndex := strings.Index(elementName, ":"); colonIndex >= 0 {
				elementName = elementName[colonIndex+1:]
			}
			element := schemaMap[elementName]
			if element != nil {
				g.collectElementFields(element, "", &fields)
			} else {
				fields = append(fields, fieldInfo{
					Name:        part.Name,
					Type:        "element: " + part.Element,
					Required:    "Unknown",
					Description: "",
					Level:       0,
					IsAttribute: false,
				})
			}
		} else if part.Type != "" {
			fields = append(fields, fieldInfo{
				Name:        part.Name,
				Type:        part.Type,
				Required:    "Unknown",
				Description: "",
				Level:       0,
				IsAttribute: false,
			})
		}
	}

	// Generate table if we have fields
	if len(fields) > 0 {
		g.generateHierarchicalFieldsTable(fields, customTypesMap)
	} else {
		g.output.P("*No fields defined.*")
	}

	g.output.P()
	return nil
}

// generateHierarchicalFieldsTable generates a table with hierarchical field display
func (g *Generator) generateHierarchicalFieldsTable(fields []fieldInfo, customTypesMap map[string]bool) {
	g.output.P("| Field | Type | Required | Description |")
	g.output.P("|-------|------|----------|-------------|")

	// Render fields with proper grouping (attributes first within each parent)
	for _, field := range fields {
		desc := field.Description
		// Leave description empty if not available instead of showing "-"

		// Create indented name with XML-style tags
		indentedName := g.getIndentedFieldName(field, field.Level)

		// Convert type to hyperlink if it's a custom type
		typeDisplay := g.formatTypeWithHyperlink(field.Type, customTypesMap)

		g.output.P("| ", indentedName, " | ", typeDisplay, " | ", field.Required, " | ", desc, " |")
	}
}

// formatTypeWithHyperlink converts a type name to a hyperlink if it's a custom type
func (g *Generator) formatTypeWithHyperlink(typeName string, customTypesMap map[string]bool) string {
	// Handle empty type names or nil map (first pass)
	if typeName == "" || customTypesMap == nil {
		return typeName
	}

	// Check for attribute suffix and extract the base type
	var attributeSuffix string
	actualTypeName := typeName
	if strings.HasSuffix(typeName, " (attribute)") {
		attributeSuffix = " (attribute)"
		actualTypeName = strings.TrimSuffix(typeName, attributeSuffix)
	}

	// Check if this is a namespaced type (e.g., "tp:sessionidType")
	var baseTypeName string
	var prefix string

	if colonIndex := strings.Index(actualTypeName, ":"); colonIndex >= 0 {
		prefix = actualTypeName[:colonIndex+1] // Include the colon
		baseTypeName = actualTypeName[colonIndex+1:]
	} else {
		baseTypeName = actualTypeName
	}

	// Skip built-in XML Schema types
	if prefix == "xs:" || prefix == "xsd:" {
		return typeName
	}

	// Check if this is a custom type (named type or inline type)
	isCustomType := customTypesMap[baseTypeName]

	// Also check if this matches an inline type
	if !isCustomType {
		// Check if baseTypeName ends with " (inline)" and we have that inline type
		if strings.HasSuffix(baseTypeName, " (inline)") {
			elementName := strings.TrimSuffix(baseTypeName, " (inline)")
			for _, inlineType := range g.inlineSimpleTypes {
				if inlineType.ElementName == elementName {
					isCustomType = true
					break
				}
			}
		}
	}

	if isCustomType {
		// Generate GitHub-compatible anchor (lowercase, replace spaces with hyphens, remove parentheses)
		anchor := strings.ToLower(baseTypeName)
		anchor = strings.ReplaceAll(anchor, " ", "-")
		anchor = strings.ReplaceAll(anchor, "(", "")
		anchor = strings.ReplaceAll(anchor, ")", "")

		// Create markdown link with original formatting, preserving attribute suffix
		return fmt.Sprintf("[%s](#%s)%s", actualTypeName, anchor, attributeSuffix)
	}

	// Return original type name if not custom
	return typeName
}

// buildFieldTree builds a hierarchical tree structure from flat field list

// getIndentedFieldName creates an indented field name with XML-style tags
func (g *Generator) getIndentedFieldName(field fieldInfo, level int) string {
	// Extract the last part of the field name
	parts := strings.Split(field.Name, ".")
	lastPart := parts[len(parts)-1]

	// Add indentation based on level (doubled spacing)
	indent := ""
	if level > 0 {
		indent = strings.Repeat("&nbsp;&nbsp;&nbsp;&nbsp;", level)
	}

	// Format based on whether it's an attribute or element
	if strings.HasPrefix(lastPart, "@") {
		// Attributes are plain text without the @ prefix
		attributeName := lastPart[1:] // Remove the @ prefix
		return indent + attributeName
	} else {
		// Elements get XML-style tags wrapped in backticks to prevent HTML interpretation
		return indent + "`<" + lastPart + ">`"
	}
}

type fieldInfo struct {
	Name        string
	Type        string
	Required    string
	Description string
	Level       int  // Hierarchy level for indentation
	IsAttribute bool // Whether this is an attribute
}

// collectElementFields recursively collects field information from an element
func (g *Generator) collectElementFields(element *xsd.Element, prefix string, fields *[]fieldInfo) {
	// Build schema map for element lookups if not already done
	schemaMap := g.buildSchemaMap()
	visited := make(map[string]bool)
	g.collectElementFieldsWithSchema(element, prefix, fields, schemaMap, visited, 0)
}

// collectElementFieldsWithSchema recursively collects field information from an element with schema context
func (g *Generator) collectElementFieldsWithSchema(element *xsd.Element, prefix string, fields *[]fieldInfo, schemaMap map[string]*xsd.Element, visited map[string]bool, level int) {
	// Handle element references first
	if element.Ref != "" {
		// This is an element reference, resolve it
		refElementName := element.Ref
		if colonIndex := strings.Index(refElementName, ":"); colonIndex >= 0 {
			refElementName = refElementName[colonIndex+1:]
		}

		// Create a unique key for cycle detection
		cycleKey := prefix + "." + refElementName
		if visited[cycleKey] {
			// Cycle detected, add a placeholder and return
			*fields = append(*fields, fieldInfo{
				Name:        cycleKey,
				Type:        "ref: " + element.Ref + " (circular reference)",
				Required:    "Unknown",
				Description: "",
				Level:       level,
				IsAttribute: false,
			})
			return
		}

		if refElement := schemaMap[refElementName]; refElement != nil {
			// Mark as visited
			visited[cycleKey] = true
			// Recursively process the referenced element
			g.collectElementFieldsWithSchema(refElement, prefix, fields, schemaMap, visited, level)
			// Unmark after processing
			delete(visited, cycleKey)
		} else {
			// Fallback if reference not found
			*fields = append(*fields, fieldInfo{
				Name:        prefix + "." + refElementName,
				Type:        "ref: " + element.Ref,
				Required:    "Unknown",
				Description: "",
				Level:       level,
				IsAttribute: false,
			})
		}
		return
	}

	fieldName := element.Name
	if prefix != "" {
		fieldName = prefix + "." + fieldName
	}

	// Determine if field is required
	required := "Yes"
	if element.MinOccurs == "0" {
		required = "No"
	}
	if element.MaxOccurs != "" && element.MaxOccurs != "1" {
		if element.MaxOccurs == "unbounded" {
			required += " (0..∞)"
		} else {
			min := element.MinOccurs
			if min == "" {
				min = "1"
			}
			required += fmt.Sprintf(" (%s..%s)", min, element.MaxOccurs)
		}
	}

	// Handle inline complex types
	if element.ComplexType != nil {
		// Create a unique key for cycle detection
		cycleKey := fieldName + ":inline"
		if visited[cycleKey] {
			// Cycle detected, add a placeholder and return
			*fields = append(*fields, fieldInfo{
				Name:        fieldName,
				Type:        "object (circular reference)",
				Required:    required,
				Description: "",
				Level:       level,
				IsAttribute: false,
			})
			return
		}

		// Mark as visited
		visited[cycleKey] = true
		// Add the complex type itself
		*fields = append(*fields, fieldInfo{
			Name:        fieldName,
			Type:        "object",
			Required:    required,
			Description: "",
			Level:       level,
			IsAttribute: false,
		})

		// Recursively collect fields from the complex type
		g.collectComplexTypeFieldsWithSchema(element.ComplexType, fieldName, fields, schemaMap, visited, level+1)
		// Unmark after processing
		delete(visited, cycleKey)
		return
	}

	// Handle inline simple types
	if element.SimpleType != nil {
		fieldType := g.deriveTypeFromInlineSimpleType(element.SimpleType)

		// Store inline simple type for later documentation
		if g.shouldDocumentInlineType(element.SimpleType) {
			typeKey := g.generateInlineTypeKey(fieldName, element.Name)
			g.inlineSimpleTypes[typeKey] = &inlineSimpleTypeInfo{
				ElementPath:   fieldName,
				ElementName:   element.Name,
				SimpleType:    element.SimpleType,
				Documentation: g.extractElementDocumentation(element),
			}
			// Use the element name as the type so it can be hyperlinked
			fieldType = element.Name + " (inline)"
		}

		*fields = append(*fields, fieldInfo{
			Name:        fieldName,
			Type:        fieldType,
			Required:    required,
			Description: "",
			Level:       level,
			IsAttribute: false,
		})
		return
	}

	// Handle type references
	if element.Type != "" {
		fieldType := element.Type

		// Remove namespace prefix for type lookup
		cleanTypeName := fieldType
		if colonIndex := strings.Index(cleanTypeName, ":"); colonIndex >= 0 {
			cleanTypeName = cleanTypeName[colonIndex+1:]
		}

		// Check if this type references another element or complex type
		if referencedElement := schemaMap[cleanTypeName]; referencedElement != nil {
			// Check if the referenced element has a simple type - if so, don't recurse
			if referencedElement.Type != "" && referencedElement.ComplexType == nil {
				// This is a simple type reference, treat as simple type
				*fields = append(*fields, fieldInfo{
					Name:        fieldName,
					Type:        fieldType,
					Required:    required,
					Description: "",
					Level:       level,
					IsAttribute: false,
				})
				return
			}

			// Create a unique key for cycle detection
			cycleKey := fieldName + ":" + cleanTypeName
			if visited[cycleKey] {
				// Cycle detected, add a placeholder and return
				*fields = append(*fields, fieldInfo{
					Name:        fieldName,
					Type:        fieldType + " (circular reference)",
					Required:    required,
					Description: "",
					Level:       level,
					IsAttribute: false,
				})
				return
			}

			// Mark as visited
			visited[cycleKey] = true
			// This type references an element, recursively process it
			g.collectElementFieldsWithSchema(referencedElement, fieldName, fields, schemaMap, visited, level)
			// Unmark after processing
			delete(visited, cycleKey)
			return
		}

		// Check if this is a complex type reference
		complexType := g.findComplexTypeByName(cleanTypeName)
		if complexType != nil {
			// Create a unique key for cycle detection
			cycleKey := fieldName + ":complex:" + cleanTypeName
			if visited[cycleKey] {
				// Cycle detected, add a placeholder and return
				*fields = append(*fields, fieldInfo{
					Name:        fieldName,
					Type:        "object (circular reference)",
					Required:    required,
					Description: "",
					Level:       level,
					IsAttribute: false,
				})
				return
			}

			// Mark as visited
			visited[cycleKey] = true
			// Add the complex type itself
			*fields = append(*fields, fieldInfo{
				Name:        fieldName,
				Type:        "object",
				Required:    required,
				Description: "",
				Level:       level,
				IsAttribute: false,
			})

			// Recursively collect fields from the complex type
			g.collectComplexTypeFieldsWithSchema(complexType, fieldName, fields, schemaMap, visited, level+1)
			// Unmark after processing
			delete(visited, cycleKey)
			return
		}

		// Simple type or unknown type
		*fields = append(*fields, fieldInfo{
			Name:        fieldName,
			Type:        fieldType,
			Required:    required,
			Description: "",
			Level:       level,
			IsAttribute: false,
		})
		return
	}

	// Element with no type or ref (shouldn't happen in well-formed XSD)
	*fields = append(*fields, fieldInfo{
		Name:        fieldName,
		Type:        "",
		Required:    required,
		Description: "",
		Level:       level,
		IsAttribute: false,
	})
}

// shouldDocumentInlineType determines if an inline simple type should be documented
func (g *Generator) shouldDocumentInlineType(simpleType *xsd.SimpleType) bool {
	// Document if it has enumerations, patterns, or other interesting constraints
	if simpleType.Restriction != nil {
		return len(simpleType.Restriction.Enumerations) > 0 ||
			len(simpleType.Restriction.Patterns) > 0 ||
			simpleType.Restriction.MinLength != nil ||
			simpleType.Restriction.MaxLength != nil ||
			simpleType.Restriction.MinInclusive != nil ||
			simpleType.Restriction.MaxInclusive != nil
	}
	// Also document list and union types
	return simpleType.List != nil || simpleType.Union != nil
}

// generateInlineTypeKey generates a unique key for an inline simple type
func (g *Generator) generateInlineTypeKey(fieldName, elementName string) string {
	return fmt.Sprintf("%s::%s", fieldName, elementName)
}

// extractElementDocumentation extracts documentation from an element
func (g *Generator) extractElementDocumentation(element *xsd.Element) string {
	if element.Annotation != nil && len(element.Annotation.Documentation) > 0 {
		return normalizeDocumentation(element.Annotation.Documentation[0].Content)
	}
	return ""
}

// createInlineTypeDedupeKey creates a deduplication key for inline simple types
func (g *Generator) createInlineTypeDedupeKey(inlineType *inlineSimpleTypeInfo) string {
	// Create a key based on element name, base type, and restrictions
	baseType := g.getSimpleTypeBase(inlineType.SimpleType)
	restrictions := g.getSimpleTypeRestrictions(inlineType.SimpleType)

	// Create a signature from restrictions
	var restrictionSig strings.Builder
	for _, restriction := range restrictions {
		restrictionSig.WriteString(fmt.Sprintf("%s:%s;", restriction.Type, restriction.Value))
	}

	return fmt.Sprintf("%s::%s::%s", inlineType.ElementName, baseType, restrictionSig.String())
}

// deriveTypeFromInlineSimpleType derives a type description from an inline simple type
func (g *Generator) deriveTypeFromInlineSimpleType(simpleType *xsd.SimpleType) string {
	if simpleType.Restriction != nil {
		// Start with the base type
		baseType := simpleType.Restriction.Base
		if baseType == "" {
			baseType = "string" // Default fallback
		}

		// If there are enumerations, it's an enum type
		if len(simpleType.Restriction.Enumerations) > 0 {
			return baseType + " (enum)"
		}

		// If there are patterns, it's a constrained type
		if len(simpleType.Restriction.Patterns) > 0 {
			return baseType + " (pattern)"
		}

		// If there are other constraints, it's a restricted type
		if simpleType.Restriction.MinInclusive != nil || simpleType.Restriction.MaxInclusive != nil ||
			simpleType.Restriction.MinLength != nil || simpleType.Restriction.MaxLength != nil {
			return baseType + " (restricted)"
		}

		// Otherwise, just use the base type
		return baseType
	}

	// Handle list types
	if simpleType.List != nil {
		itemType := simpleType.List.ItemType
		if itemType == "" {
			itemType = "string"
		}
		return itemType + " (list)"
	}

	// Handle union types
	if simpleType.Union != nil {
		return "union"
	}

	// Fallback
	return "string"
}

// findComplexTypeByName finds a complex type definition by name across all schemas
func (g *Generator) findComplexTypeByName(typeName string) *xsd.ComplexType {
	doc := g.definitions
	if doc.Types != nil {
		for _, schema := range doc.Types.Schemas {
			for i := range schema.ComplexTypes {
				if schema.ComplexTypes[i].Name == typeName {
					return &schema.ComplexTypes[i]
				}
			}
		}
	}
	return nil
}

// collectComplexTypeFieldsWithSchema collects fields from a complex type with schema context
func (g *Generator) collectComplexTypeFieldsWithSchema(complexType *xsd.ComplexType, prefix string, fields *[]fieldInfo, schemaMap map[string]*xsd.Element, visited map[string]bool, level int) {
	// Collect attributes first
	for _, attr := range complexType.Attributes {
		required := "No"
		if attr.Use == "required" {
			required = "Yes"
		}

		*fields = append(*fields, fieldInfo{
			Name:        prefix + ".@" + attr.Name,
			Type:        attr.Type + " (attribute)",
			Required:    required,
			Description: "",
			Level:       level,
			IsAttribute: true,
		})
	}

	// Then collect elements
	if complexType.Sequence != nil {
		g.collectSequenceFieldsWithSchema(complexType.Sequence, prefix, fields, schemaMap, visited, level)
	}
	if complexType.Choice != nil {
		g.collectChoiceFieldsWithSchema(complexType.Choice, prefix, fields, schemaMap, visited, level)
	}
	if complexType.All != nil {
		g.collectAllFieldsWithSchema(complexType.All, prefix, fields, schemaMap, visited, level)
	}
}

// collectSequenceFieldsWithSchema collects fields from a sequence with schema context
func (g *Generator) collectSequenceFieldsWithSchema(sequence *xsd.Sequence, prefix string, fields *[]fieldInfo, schemaMap map[string]*xsd.Element, visited map[string]bool, level int) {
	for i := range sequence.Elements {
		g.collectElementFieldsWithSchema(&sequence.Elements[i], prefix, fields, schemaMap, visited, level)
	}
	for i := range sequence.Sequences {
		g.collectSequenceFieldsWithSchema(&sequence.Sequences[i], prefix, fields, schemaMap, visited, level)
	}
	for i := range sequence.Choices {
		g.collectChoiceFieldsWithSchema(&sequence.Choices[i], prefix, fields, schemaMap, visited, level)
	}
}

// collectChoiceFieldsWithSchema collects fields from a choice with schema context
func (g *Generator) collectChoiceFieldsWithSchema(choice *xsd.Choice, prefix string, fields *[]fieldInfo, schemaMap map[string]*xsd.Element, visited map[string]bool, level int) {
	// Add a note about choice
	*fields = append(*fields, fieldInfo{
		Name:        prefix + " (choice)",
		Type:        "one of the following",
		Required:    "Yes",
		Description: "Choose one of the following options",
		Level:       level,
		IsAttribute: false,
	})

	for i := range choice.Elements {
		g.collectElementFieldsWithSchema(&choice.Elements[i], prefix, fields, schemaMap, visited, level)
	}
	for i := range choice.Sequences {
		g.collectSequenceFieldsWithSchema(&choice.Sequences[i], prefix, fields, schemaMap, visited, level)
	}
	for i := range choice.Choices {
		g.collectChoiceFieldsWithSchema(&choice.Choices[i], prefix, fields, schemaMap, visited, level)
	}
}

// collectAllFieldsWithSchema collects fields from an all group with schema context
func (g *Generator) collectAllFieldsWithSchema(all *xsd.All, prefix string, fields *[]fieldInfo, schemaMap map[string]*xsd.Element, visited map[string]bool, level int) {
	for i := range all.Elements {
		g.collectElementFieldsWithSchema(&all.Elements[i], prefix, fields, schemaMap, visited, level)
	}
}

// generateCustomTypesSection generates documentation for custom types defined in the schema
func (g *Generator) generateCustomTypesSection() error {
	doc := g.definitions

	if doc.Types == nil || len(doc.Types.Schemas) == 0 {
		return nil
	}

	// Collect all custom types from all schemas
	var simpleTypes []customSimpleType
	var complexTypes []customComplexType

	for _, schema := range doc.Types.Schemas {
		// Collect simple types
		for i := range schema.SimpleTypes {
			simpleType := &schema.SimpleTypes[i]
			if simpleType.Name != "" {
				customType := customSimpleType{
					Name:          simpleType.Name,
					BaseType:      g.getSimpleTypeBase(simpleType),
					Documentation: g.getSimpleTypeDocumentation(simpleType),
					Restrictions:  g.getSimpleTypeRestrictions(simpleType),
				}
				simpleTypes = append(simpleTypes, customType)
			}
		}

		// Collect complex types (only those with documentation)
		for i := range schema.ComplexTypes {
			complexType := &schema.ComplexTypes[i]
			if complexType.Name != "" {
				documentation := g.getComplexTypeDocumentation(complexType)
				// Only include complex types that have actual documentation
				if documentation != "" {
					customType := customComplexType{
						Name:          complexType.Name,
						Documentation: documentation,
					}
					complexTypes = append(complexTypes, customType)
				}
			}
		}
	}

	// Collect inline simple types (deduplicate and sort by element name)
	inlineTypesSeen := make(map[string]bool)
	var inlineTypes []customSimpleType
	for _, inlineType := range g.inlineSimpleTypes {
		// Create a deduplication key based on element name and restrictions
		dedupeKey := g.createInlineTypeDedupeKey(inlineType)
		if !inlineTypesSeen[dedupeKey] {
			customType := customSimpleType{
				Name:          inlineType.ElementName + " (inline)",
				BaseType:      g.getSimpleTypeBase(inlineType.SimpleType),
				Documentation: inlineType.Documentation,
				Restrictions:  g.getSimpleTypeRestrictions(inlineType.SimpleType),
			}
			inlineTypes = append(inlineTypes, customType)
			inlineTypesSeen[dedupeKey] = true
		}
	}

	// Sort inline types by name for consistent output
	sort.Slice(inlineTypes, func(i, j int) bool {
		return inlineTypes[i].Name < inlineTypes[j].Name
	})

	// Add sorted inline types to simpleTypes
	simpleTypes = append(simpleTypes, inlineTypes...)

	// Only generate the section if we have custom types
	if len(simpleTypes) == 0 && len(complexTypes) == 0 {
		return nil
	}

	g.output.P("## Custom Types")
	g.output.P()
	g.output.P("This section documents the custom data types defined in the schema.")
	g.output.P()

	// Generate simple types section
	if len(simpleTypes) > 0 {
		g.output.P("### Simple Types")
		g.output.P()
		for _, simpleType := range simpleTypes {
			g.generateSimpleTypeDoc(&simpleType)
		}
	}

	// Generate complex types section
	if len(complexTypes) > 0 {
		g.output.P("### Complex Types")
		g.output.P()
		for _, complexType := range complexTypes {
			g.generateComplexTypeDoc(&complexType)
		}
	}

	return nil
}

type customSimpleType struct {
	Name          string
	BaseType      string
	Documentation string
	Restrictions  []typeRestriction
}

type customComplexType struct {
	Name          string
	Documentation string
}

type typeRestriction struct {
	Type        string // "enumeration", "pattern", "length", etc.
	Value       string
	Description string
}

// getSimpleTypeBase extracts the base type from a simple type restriction
func (g *Generator) getSimpleTypeBase(simpleType *xsd.SimpleType) string {
	if simpleType.Restriction != nil && simpleType.Restriction.Base != "" {
		base := simpleType.Restriction.Base
		// Remove namespace prefix if present
		if colonIndex := strings.Index(base, ":"); colonIndex >= 0 {
			base = base[colonIndex+1:]
		}
		return base
	}
	return ""
}

// getSimpleTypeDocumentation extracts documentation from a simple type
func (g *Generator) getSimpleTypeDocumentation(simpleType *xsd.SimpleType) string {
	if simpleType.Annotation != nil && len(simpleType.Annotation.Documentation) > 0 {
		return normalizeDocumentation(simpleType.Annotation.Documentation[0].Content)
	}
	return ""
}

// getSimpleTypeRestrictions extracts restrictions from a simple type
func (g *Generator) getSimpleTypeRestrictions(simpleType *xsd.SimpleType) []typeRestriction {
	var restrictions []typeRestriction

	if simpleType.Restriction == nil {
		return restrictions
	}

	// Handle enumerations
	for _, enum := range simpleType.Restriction.Enumerations {
		restrictions = append(restrictions, typeRestriction{
			Type:  "enumeration",
			Value: enum.Value,
		})
	}

	// Handle patterns
	for _, pattern := range simpleType.Restriction.Patterns {
		restrictions = append(restrictions, typeRestriction{
			Type:  "pattern",
			Value: pattern.Value,
		})
	}

	// Handle length constraints
	if simpleType.Restriction.MinLength != nil {
		restrictions = append(restrictions, typeRestriction{
			Type:  "minLength",
			Value: simpleType.Restriction.MinLength.Value,
		})
	}
	if simpleType.Restriction.MaxLength != nil {
		restrictions = append(restrictions, typeRestriction{
			Type:  "maxLength",
			Value: simpleType.Restriction.MaxLength.Value,
		})
	}

	return restrictions
}

// getComplexTypeDocumentation extracts documentation from a complex type
func (g *Generator) getComplexTypeDocumentation(complexType *xsd.ComplexType) string {
	if complexType.Annotation != nil && len(complexType.Annotation.Documentation) > 0 {
		return normalizeDocumentation(complexType.Annotation.Documentation[0].Content)
	}
	return ""
}

// generateSimpleTypeDoc generates documentation for a single simple type
func (g *Generator) generateSimpleTypeDoc(simpleType *customSimpleType) {
	g.output.P("#### `", simpleType.Name, "`")
	g.output.P()

	if simpleType.BaseType != "" {
		g.output.P("**Base Type:** `", simpleType.BaseType, "`")
		g.output.P()
	}

	if simpleType.Documentation != "" {
		g.output.P(simpleType.Documentation)
		g.output.P()
	}

	// Document restrictions
	if len(simpleType.Restrictions) > 0 {
		// Group restrictions by type
		enums := []typeRestriction{}
		patterns := []typeRestriction{}
		lengths := []typeRestriction{}

		for _, restriction := range simpleType.Restrictions {
			switch restriction.Type {
			case "enumeration":
				enums = append(enums, restriction)
			case "pattern":
				patterns = append(patterns, restriction)
			case "minLength", "maxLength":
				lengths = append(lengths, restriction)
			}
		}

		if len(enums) > 0 {
			g.output.P("**Allowed Values:**")
			for _, enum := range enums {
				g.output.P("- `", enum.Value, "`")
			}
			g.output.P()
		}

		if len(patterns) > 0 {
			g.output.P("**Pattern:**")
			for _, pattern := range patterns {
				g.output.P("- `", pattern.Value, "`")
			}
			g.output.P()
		}

		if len(lengths) > 0 {
			g.output.P("**Length Constraints:**")
			for _, length := range lengths {
				g.output.P("- ", length.Type, ": ", length.Value)
			}
			g.output.P()
		}
	}

	g.output.P()
}

// generateComplexTypeDoc generates documentation for a single complex type
func (g *Generator) generateComplexTypeDoc(complexType *customComplexType) {
	g.output.P("#### `", complexType.Name, "`")
	g.output.P()

	if complexType.Documentation != "" {
		g.output.P(complexType.Documentation)
		g.output.P()
	}

	g.output.P()
}
