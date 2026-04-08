package soapgen

import "strings"

// BindingStyle represents the SOAP binding style configuration
type BindingStyle struct {
	Style         string // "document" or "rpc"
	Use           string // "literal" or "encoded"
	EncodingStyle string // optional encoding style URI
}

// getBindingStyle determines the SOAP binding style for the given WSDL
func (g *Generator) getBindingStyle() BindingStyle {
	// Find the first SOAP binding (prefer SOAP 1.1)
	soapBindings := g.getSOAPBindings()
	if len(soapBindings) == 0 {
		// Default to document/literal if no bindings found
		return BindingStyle{Style: "document", Use: "literal"}
	}

	binding := soapBindings[0]
	style := BindingStyle{}

	// Get style from binding
	if binding.SOAP11Binding != nil {
		style.Style = binding.SOAP11Binding.Style
	} else if binding.SOAP12Binding != nil {
		style.Style = binding.SOAP12Binding.Style
	}

	// Default to document if not specified
	if style.Style == "" {
		style.Style = "document"
	}

	// Get use and encodingStyle from first operation
	if len(binding.BindingOperations) > 0 {
		op := binding.BindingOperations[0]
		if op.Output != nil {
			if op.Output.SOAP11Body != nil {
				style.Use = op.Output.SOAP11Body.Use
				style.EncodingStyle = op.Output.SOAP11Body.EncodingStyle
			} else if op.Output.SOAP12Body != nil {
				style.Use = op.Output.SOAP12Body.Use
				style.EncodingStyle = op.Output.SOAP12Body.EncodingStyle
			}
		}
	}

	// Default to literal if not specified
	if style.Use == "" {
		style.Use = "literal"
	}

	return style
}

// getConsistentTypeName returns the Go type name with consistent wrapper naming
func (g *Generator) getConsistentTypeName(xmlElementName string, bindingStyle BindingStyle) string {
	baseName := toGoName(xmlElementName)

	// Use the same logic as type generation for consistency
	if g.shouldUseWrapperForElement(xmlElementName, bindingStyle) {
		return baseName + "Wrapper"
	}

	// For non-operation elements or non-wrapper styles, use the base name
	return baseName
}

// isOperationMessageElement checks if the given element name is used in any SOAP operation message
func (g *Generator) isOperationMessageElement(xmlElementName string) bool {
	// Check all messages referenced by operations
	for _, message := range g.definitions.Messages {
		for _, part := range message.Parts {
			if part.Element != "" {
				// Extract element name (remove namespace prefix)
				elementName := part.Element
				if colonIdx := strings.LastIndex(elementName, ":"); colonIdx != -1 {
					elementName = elementName[colonIdx+1:]
				}
				if elementName == xmlElementName {
					return true
				}
			}
		}
	}
	return false
}

// shouldUseWrapperForElement determines if a specific element should use wrapper naming.
// An element gets the Wrapper suffix when:
//   - it is used as a SOAP operation message element (existing behaviour), OR
//   - its name collides with a simpleType or complexType in the same WSDL,
//     which would otherwise produce duplicate Go type declarations.
func (g *Generator) shouldUseWrapperForElement(elementName string, bindingStyle BindingStyle) bool {
	// Classification-based approach: Use wrapper naming for operation elements in appropriate binding styles
	if bindingStyle.Style == "rpc" {
		if g.isOperationMessageElement(elementName) {
			return true
		}
	}

	if bindingStyle.Style == "document" && bindingStyle.Use == "literal" {
		if g.isOperationMessageElement(elementName) {
			return true
		}
	}

	// Use wrapper naming when the element name collides with an existing
	// simpleType or complexType definition. In XSD these are separate symbol
	// spaces, but in Go they would produce duplicate type declarations.
	if g.elementNameCollidesWithType(elementName) {
		return true
	}

	return false
}

// elementNameCollidesWithType checks if an element's Go type name would
// collide with any simpleType or complexType Go type name in the WSDL schemas.
// Comparison is done on Go names (after toGoName) because different XSD names
// can produce the same Go identifier (e.g., "addPolicy_Request" and
// "addPolicyRequest" both become "AddPolicyRequest").
func (g *Generator) elementNameCollidesWithType(elementName string) bool {
	if g.definitions.Types == nil {
		return false
	}
	goName := toGoName(elementName)
	for _, schema := range g.definitions.Types.Schemas {
		for _, st := range schema.SimpleTypes {
			if toGoName(st.Name) == goName {
				return true
			}
		}
		for _, ct := range schema.ComplexTypes {
			if toGoName(ct.Name) == goName {
				return true
			}
		}
	}
	return false
}
