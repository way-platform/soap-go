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

// shouldUseWrapperNaming determines if wrapper naming should be used based on binding style
func (g *Generator) shouldUseWrapperNaming(bindingStyle BindingStyle) bool {
	switch {
	case bindingStyle.Style == "rpc":
		// RPC style always uses wrappers for operation parameters
		return true
	case bindingStyle.Style == "document" && bindingStyle.Use == "literal":
		// Document/literal: Use wrappers if the schema is designed with operation wrappers
		// We detect this by checking if there are message wrapper elements
		return g.hasMessageWrapperElements()
	case bindingStyle.Style == "document" && bindingStyle.Use == "encoded":
		// Document/encoded: Typically no additional wrappers
		return false
	default:
		// Default to wrapper naming for unknown styles
		return true
	}
}

// hasMessageWrapperElements checks if the WSDL schema contains message wrapper elements
func (g *Generator) hasMessageWrapperElements() bool {
	for _, schema := range g.definitions.Types.Schemas {
		for _, element := range schema.Elements {
			if g.isMessageWrapperElement(&element) {
				return true
			}
		}
	}
	return false
}

// getConsistentTypeName returns the Go type name with consistent wrapper naming
func (g *Generator) getConsistentTypeName(xmlElementName string, bindingStyle BindingStyle) string {
	baseName := toGoName(xmlElementName)

	// For RPC style, always use Wrapper suffix for operation elements
	if bindingStyle.Style == "rpc" {
		// In RPC style, operation elements are always wrappers
		return baseName + "Wrapper"
	}

	// For Document/Literal style with wrapper elements, use consistent naming
	if bindingStyle.Style == "document" && bindingStyle.Use == "literal" {
		// If this WSDL uses message wrapper elements, be consistent and use Wrapper suffix for all operation elements
		if g.shouldUseWrapperNaming(bindingStyle) {
			// Check if this is an operation message element (wrapper or not)
			if g.isOperationMessageElement(xmlElementName) {
				return baseName + "Wrapper"
			}
		}
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
// shouldUseWrapperForElement determines if a specific element should use wrapper naming
// based on binding style and whether it's used in SOAP operations
func (g *Generator) shouldUseWrapperForElement(elementName string, bindingStyle BindingStyle) bool {
	// Classification-based approach: Use wrapper naming for operation elements in appropriate binding styles
	if bindingStyle.Style == "rpc" {
		// RPC style: ALL operation elements use wrappers
		return g.isOperationMessageElement(elementName)
	}
	
	if bindingStyle.Style == "document" && bindingStyle.Use == "literal" {
		// Document/Literal: Use wrapper naming for ALL operation elements for consistency
		return g.isOperationMessageElement(elementName)
	}
	
	// Other binding styles: no wrapper naming
	return false
}
