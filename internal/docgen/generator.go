package docgen

import (
	"fmt"
	"strings"

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

	// Title
	g.output.P("# ", doc.Name)
	if doc.TargetNamespace != "" {
		g.output.P()
		g.output.P("**Namespace:** `", doc.TargetNamespace, "`")
	}
	g.output.P()

	// Build schema map for element lookups
	schemaMap := g.buildSchemaMap()

	// Generate documentation for each service
	for _, service := range doc.Service {
		if err := g.generateServiceDoc(&service, schemaMap); err != nil {
			return err
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
	doc := g.definitions

	g.output.P("## ", service.Name)
	g.output.P()

	// Add service description if available
	if service.Documentation != "" {
		g.output.P(normalizeDocumentation(service.Documentation))
		g.output.P()
	}

	// Find the corresponding PortType for this service
	var portType *wsdl.PortType
	for _, binding := range doc.Binding {
		for _, port := range service.Ports {
			if strings.Contains(port.Binding, binding.Name) {
				// Find the PortType referenced by this binding
				typeName := strings.TrimPrefix(binding.Type, "tns:")
				for i := range doc.PortType {
					if doc.PortType[i].Name == typeName {
						portType = &doc.PortType[i]
						break
					}
				}
				break
			}
		}
		if portType != nil {
			break
		}
	}

	if portType == nil {
		g.output.P("*No operations found for this service.*")
		g.output.P()
		return nil
	}

	// Generate documentation for each operation
	for _, operation := range portType.Operations {
		if err := g.generateOperationDoc(&operation, schemaMap); err != nil {
			return err
		}
	}

	return nil
}

// generateOperationDoc generates documentation for a single operation
func (g *Generator) generateOperationDoc(operation *wsdl.Operation, schemaMap map[string]*xsd.Element) error {
	g.output.P("### ", operation.Name)
	g.output.P()

	// Add operation description if available
	if operation.Documentation != "" {
		g.output.P(normalizeDocumentation(operation.Documentation))
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

	g.output.P()
	return nil
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

	// Generate field documentation for each part
	for _, part := range message.Parts {
		if part.Element != "" {
			elementName := strings.TrimPrefix(part.Element, "tns:")
			element := schemaMap[elementName]
			if element != nil {
				g.generateElementFields(element, 0)
			} else {
				g.output.P("- ", part.Name, " (element: ", part.Element, ")")
			}
		} else if part.Type != "" {
			g.output.P("- ", part.Name, " (type: ", part.Type, ")")
		}
	}

	g.output.P()
	return nil
}

// generateElementFields generates hierarchical bullet list for element fields
func (g *Generator) generateElementFields(element *xsd.Element, depth int) {
	indent := strings.Repeat("  ", depth)

	// Generate the field name and type
	fieldName := element.Name
	fieldType := element.Type
	if fieldType == "" && element.ComplexType != nil {
		fieldType = "complex"
	}

	// Add occurrence information
	occurrenceInfo := ""
	if element.MinOccurs != "" || element.MaxOccurs != "" {
		min := element.MinOccurs
		max := element.MaxOccurs
		if min == "" {
			min = "1"
		}
		if max == "" {
			max = "1"
		}
		if min != "1" || max != "1" {
			occurrenceInfo = fmt.Sprintf(" (%s..%s)", min, max)
		}
	}

	if fieldType != "" {
		g.output.P(indent, "- **", fieldName, "** (", fieldType, ")", occurrenceInfo)
	} else {
		g.output.P(indent, "- **", fieldName, "**", occurrenceInfo)
	}

	// If this element has a complex type, recursively generate its fields
	if element.ComplexType != nil {
		g.generateComplexTypeFields(element.ComplexType, depth+1)
	}
}

// generateComplexTypeFields generates fields for a complex type
func (g *Generator) generateComplexTypeFields(complexType *xsd.ComplexType, depth int) {
	if complexType.Sequence != nil {
		g.generateSequenceFields(complexType.Sequence, depth)
	}
	if complexType.Choice != nil {
		g.generateChoiceFields(complexType.Choice, depth)
	}
	if complexType.All != nil {
		g.generateAllFields(complexType.All, depth)
	}

	// Generate attributes
	for _, attr := range complexType.Attributes {
		indent := strings.Repeat("  ", depth)
		required := ""
		if attr.Use == "required" {
			required = " (required)"
		}
		g.output.P(indent, "- **@", attr.Name, "** (", attr.Type, ")", required, " *[attribute]*")
	}
}

// generateSequenceFields generates fields for a sequence
func (g *Generator) generateSequenceFields(sequence *xsd.Sequence, depth int) {
	for i := range sequence.Elements {
		g.generateElementFields(&sequence.Elements[i], depth)
	}
	for i := range sequence.Sequences {
		g.generateSequenceFields(&sequence.Sequences[i], depth)
	}
	for i := range sequence.Choices {
		g.generateChoiceFields(&sequence.Choices[i], depth)
	}
}

// generateChoiceFields generates fields for a choice
func (g *Generator) generateChoiceFields(choice *xsd.Choice, depth int) {
	indent := strings.Repeat("  ", depth)
	g.output.P(indent, "- **Choice of:**")

	for i := range choice.Elements {
		g.generateElementFields(&choice.Elements[i], depth+1)
	}
	for i := range choice.Sequences {
		g.generateSequenceFields(&choice.Sequences[i], depth+1)
	}
	for i := range choice.Choices {
		g.generateChoiceFields(&choice.Choices[i], depth+1)
	}
}

// generateAllFields generates fields for an all group
func (g *Generator) generateAllFields(all *xsd.All, depth int) {
	for i := range all.Elements {
		g.generateElementFields(&all.Elements[i], depth)
	}
}
