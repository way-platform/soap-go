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

	if doc.TargetNamespace != "" {
		g.output.P()
		g.output.P("**Namespace:** `", doc.TargetNamespace, "`")
	}

	// Add endpoint information if available
	if endpoint := g.getServiceEndpoint(); endpoint != "" {
		g.output.P("**Endpoint:** `", endpoint, "`")
	}
	g.output.P()

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
		if desc == "" {
			desc = "No description available"
		}
		// Create anchor link to the operation section
		anchor := strings.ToLower(strings.ReplaceAll(op.Name, " ", "-"))
		g.output.P("- **[", op.Name, "](#", anchor, ")** - ", desc)
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
				typeName := strings.TrimPrefix(binding.Type, "tns:")
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
	// Create anchor-friendly ID for the operation
	anchor := strings.ToLower(strings.ReplaceAll(operation.Name, " ", "-"))
	g.output.P("### ", operation.Name, " {#", anchor, "}")
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
	cleanMessageName := strings.TrimPrefix(messageName, "tns:")
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
			elementName := strings.TrimPrefix(part.Element, "tns:")
			element := schemaMap[elementName]
			if element != nil {
				g.collectElementFields(element, "", &fields)
			} else {
				fields = append(fields, fieldInfo{
					Name:        part.Name,
					Type:        "element: " + part.Element,
					Required:    "Unknown",
					Description: "",
				})
			}
		} else if part.Type != "" {
			fields = append(fields, fieldInfo{
				Name:        part.Name,
				Type:        part.Type,
				Required:    "Unknown",
				Description: "",
			})
		}
	}

	// Generate table if we have fields
	if len(fields) > 0 {
		g.output.P("| Field | Type | Required | Description |")
		g.output.P("|-------|------|----------|-------------|")

		for _, field := range fields {
			desc := field.Description
			if desc == "" {
				desc = "-"
			}
			g.output.P("| ", field.Name, " | ", field.Type, " | ", field.Required, " | ", desc, " |")
		}
	} else {
		g.output.P("*No fields defined.*")
	}

	g.output.P()
	return nil
}

type fieldInfo struct {
	Name        string
	Type        string
	Required    string
	Description string
}

// collectElementFields recursively collects field information from an element
func (g *Generator) collectElementFields(element *xsd.Element, prefix string, fields *[]fieldInfo) {
	fieldName := element.Name
	if prefix != "" {
		fieldName = prefix + "." + fieldName
	}

	fieldType := element.Type
	if fieldType == "" && element.ComplexType != nil {
		fieldType = "complex"
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

	// For complex types, we'll add the parent and then recurse
	if element.ComplexType != nil {
		// Add the complex type itself
		*fields = append(*fields, fieldInfo{
			Name:        fieldName,
			Type:        "object",
			Required:    required,
			Description: "",
		})

		// Recursively collect fields from the complex type
		g.collectComplexTypeFields(element.ComplexType, fieldName, fields)
	} else {
		// Simple type
		*fields = append(*fields, fieldInfo{
			Name:        fieldName,
			Type:        fieldType,
			Required:    required,
			Description: "",
		})
	}
}

// collectComplexTypeFields collects fields from a complex type
func (g *Generator) collectComplexTypeFields(complexType *xsd.ComplexType, prefix string, fields *[]fieldInfo) {
	if complexType.Sequence != nil {
		g.collectSequenceFields(complexType.Sequence, prefix, fields)
	}
	if complexType.Choice != nil {
		g.collectChoiceFields(complexType.Choice, prefix, fields)
	}
	if complexType.All != nil {
		g.collectAllFields(complexType.All, prefix, fields)
	}

	// Collect attributes
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
		})
	}
}

// collectSequenceFields collects fields from a sequence
func (g *Generator) collectSequenceFields(sequence *xsd.Sequence, prefix string, fields *[]fieldInfo) {
	for i := range sequence.Elements {
		g.collectElementFields(&sequence.Elements[i], prefix, fields)
	}
	for i := range sequence.Sequences {
		g.collectSequenceFields(&sequence.Sequences[i], prefix, fields)
	}
	for i := range sequence.Choices {
		g.collectChoiceFields(&sequence.Choices[i], prefix, fields)
	}
}

// collectChoiceFields collects fields from a choice
func (g *Generator) collectChoiceFields(choice *xsd.Choice, prefix string, fields *[]fieldInfo) {
	// Add a note about choice
	*fields = append(*fields, fieldInfo{
		Name:        prefix + " (choice)",
		Type:        "one of the following",
		Required:    "Yes",
		Description: "Choose one of the following options",
	})

	for i := range choice.Elements {
		g.collectElementFields(&choice.Elements[i], prefix, fields)
	}
	for i := range choice.Sequences {
		g.collectSequenceFields(&choice.Sequences[i], prefix, fields)
	}
	for i := range choice.Choices {
		g.collectChoiceFields(&choice.Choices[i], prefix, fields)
	}
}

// collectAllFields collects fields from an all group
func (g *Generator) collectAllFields(all *xsd.All, prefix string, fields *[]fieldInfo) {
	for i := range all.Elements {
		g.collectElementFields(&all.Elements[i], prefix, fields)
	}
}
