package docgen

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
	"github.com/way-platform/soap-go/xsd"
)

// Generator generates Markdown documentation from WSDL definitions.
type Generator struct {
	definitions *wsdl.Definitions
	output      *codegen.File
}

// NewGenerator creates a new [Generator] for the given WSDL definitions.
func NewGenerator(filename string, definitions *wsdl.Definitions) *Generator {
	return &Generator{
		definitions: definitions,
		output:      codegen.NewFile(filename),
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

	// Generate documentation for each service
	for _, service := range doc.Service {
		if err := g.generateServiceDoc(&service, schemaMap); err != nil {
			return err
		}
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

// normalizeDocumentation normalizes whitespace in documentation strings
// by trimming leading/trailing whitespace and replacing any sequence of
// whitespace characters (including newlines) with a single space
func normalizeDocumentation(doc string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(doc), " "))
}

// generateServiceDoc generates documentation for a single service
func (g *Generator) generateServiceDoc(service *wsdl.Service, schemaMap map[string]*xsd.Element) error {
	// Find the corresponding PortType for this service
	portType := g.findPortTypeForService(service)

	if portType == nil {
		g.output.P("*No operations found for this service.*")
		g.output.P()
		return nil
	}

	// Generate documentation for each operation
	for _, operation := range portType.Operations {
		if err := g.generateOperationDoc(&operation, service, schemaMap); err != nil {
			return err
		}
	}

	return nil
}

// generateOperationDoc generates documentation for a single operation
func (g *Generator) generateOperationDoc(operation *wsdl.Operation, service *wsdl.Service, schemaMap map[string]*xsd.Element) error {
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
		if err := g.generateMessageDoc("Request", operation.Input.Message, schemaMap); err != nil {
			return err
		}
	}

	// Generate response documentation
	if operation.Output != nil {
		if err := g.generateMessageDoc("Response", operation.Output.Message, schemaMap); err != nil {
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
func (g *Generator) generateMessageDoc(messageType, messageName string, schemaMap map[string]*xsd.Element) error {
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
		g.generateHierarchicalFieldsTable(fields)
	} else {
		g.output.P("*No fields defined.*")
	}

	g.output.P()
	return nil
}

// generateHierarchicalFieldsTable generates a table with hierarchical field display
func (g *Generator) generateHierarchicalFieldsTable(fields []fieldInfo) {
	g.output.P("| Field | Type | Required | Description |")
	g.output.P("|-------|------|----------|-------------|")

	// Render fields with proper grouping (attributes first within each parent)
	for _, field := range fields {
		desc := field.Description
		// Leave description empty if not available instead of showing "-"
		
		// Create indented name with XML-style tags
		indentedName := g.getIndentedFieldName(field, field.Level)
		g.output.P("| ", indentedName, " | ", field.Type, " | ", field.Required, " | ", desc, " |")
	}
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
