package soapgen

import (
	"fmt"
	"strings"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/wsdl"
)

// generateClientFile generates a Go file with SOAP client implementation
func (g *Generator) generateClientFile(packageName, filename string) (*codegen.File, error) {
	// Check if there are any SOAP operations to generate client for
	soapBindings := g.getSOAPBindings()
	if len(soapBindings) == 0 {
		// No SOAP bindings found, don't generate client file
		return nil, nil
	}

	// Check if any bindings have operations
	hasOperations := false
	for _, binding := range soapBindings {
		portType := g.getPortTypeForBinding(binding)
		if portType != nil && len(portType.Operations) > 0 {
			hasOperations = true
			break
		}
	}

	if !hasOperations {
		// No operations found, don't generate client file
		return nil, nil
	}

	file := codegen.NewFile(filename, packageName)

	// Set custom package name for soap-go to use "soap" instead of "soapgo"
	file.SetPackageName("github.com/way-platform/soap-go", "soap")

	// Add package declaration
	file.P("package ", packageName)
	file.P()

	// Imports are now handled automatically via QualifiedGoIdent calls

	// Generate client option types
	g.generateClientOptions(file)

	// Generate client struct
	g.generateClientStruct(file)

	// Generate NewClient function
	g.generateNewClientFunction(file)

	// Generate operation methods
	err := g.generateOperationMethods(file)
	if err != nil {
		return nil, fmt.Errorf("failed to generate operation methods: %w", err)
	}

	// Generate helper functions
	g.generateHelperFunctions(file)

	return file, nil
}

// generateClientOptions generates type aliases for SOAP client options
func (g *Generator) generateClientOptions(file *codegen.File) {
	file.P("// ClientOption configures a Client.")
	file.P("type ClientOption = ", file.QualifiedGoIdent(codegen.SOAPClientOptionIdent))
	file.P()
}

// generateClientStruct generates the main Client struct
func (g *Generator) generateClientStruct(file *codegen.File) {
	file.P("// Client is a SOAP client for this service.")
	file.P("type Client struct {")
	file.P("\t*", file.QualifiedGoIdent(codegen.SOAPClientIdent))
	file.P("}")
	file.P()
}

// generateNewClientFunction generates the NewClient constructor
func (g *Generator) generateNewClientFunction(file *codegen.File) {
	// Extract default endpoint from service definitions
	endpoint := g.getDefaultEndpoint()

	file.P("// NewClient creates a new SOAP client.")
	file.P("func NewClient(opts ...ClientOption) (*Client, error) {")
	if endpoint != "" {
		file.P("\tsoapOpts := append([]", file.QualifiedGoIdent(codegen.SOAPClientOptionIdent), "{")
		file.P("\t\t", file.QualifiedGoIdent(codegen.SOAPWithEndpointIdent), "(\"", endpoint, "\"),")
		file.P("\t}, opts...)")
		file.P("\tsoapClient, err := ", file.QualifiedGoIdent(codegen.SOAPNewClientIdent), "(soapOpts...)")
	} else {
		file.P("\tsoapClient, err := ", file.QualifiedGoIdent(codegen.SOAPNewClientIdent), "(opts...)")
	}
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to create SOAP client: %w\", err)")
	file.P("\t}")
	file.P("\treturn &Client{")
	file.P("\t\tClient: soapClient,")
	file.P("\t}, nil")
	file.P("}")
	file.P()
}

// getDefaultEndpoint extracts the default endpoint from service definitions
func (g *Generator) getDefaultEndpoint() string {
	for _, service := range g.definitions.Service {
		for _, port := range service.Ports {
			if port.SOAP11Address != nil {
				return port.SOAP11Address.Location
			}
			if port.SOAP12Address != nil {
				return port.SOAP12Address.Location
			}
		}
	}
	return ""
}

// generateOperationMethods generates methods for each SOAP operation
func (g *Generator) generateOperationMethods(file *codegen.File) error {
	// Find SOAP bindings and their operations
	soapBindings := g.getSOAPBindings()

	for _, binding := range soapBindings {
		portType := g.getPortTypeForBinding(binding)
		if portType == nil {
			continue
		}

		for _, operation := range portType.Operations {
			err := g.generateOperationMethod(file, &operation, binding)
			if err != nil {
				return fmt.Errorf("failed to generate method for operation %s: %w", operation.Name, err)
			}
		}
	}

	return nil
}

// getSOAPBindings returns SOAP bindings from the WSDL, preferring SOAP 1.1 over SOAP 1.2
func (g *Generator) getSOAPBindings() []*wsdl.Binding {
	var soap11Bindings []*wsdl.Binding
	var soap12Bindings []*wsdl.Binding

	for i := range g.definitions.Binding {
		binding := &g.definitions.Binding[i]
		if binding.SOAP11Binding != nil {
			soap11Bindings = append(soap11Bindings, binding)
		} else if binding.SOAP12Binding != nil {
			soap12Bindings = append(soap12Bindings, binding)
		}
	}

	// Prefer SOAP 1.1 bindings as per README.md specification
	if len(soap11Bindings) > 0 {
		return soap11Bindings
	}

	// Fall back to SOAP 1.2 if no SOAP 1.1 bindings are available
	return soap12Bindings
}

// getPortTypeForBinding finds the port type that matches the given binding
func (g *Generator) getPortTypeForBinding(binding *wsdl.Binding) *wsdl.PortType {
	// Extract the local name from the binding type (remove namespace prefix if present)
	bindingType := binding.Type
	if colonIdx := strings.LastIndex(bindingType, ":"); colonIdx != -1 {
		bindingType = bindingType[colonIdx+1:]
	}

	for i := range g.definitions.PortType {
		portType := &g.definitions.PortType[i]
		if portType.Name == bindingType {
			return portType
		}
	}
	return nil
}

// generateOperationMethod generates a single operation method
func (g *Generator) generateOperationMethod(file *codegen.File, operation *wsdl.Operation, binding *wsdl.Binding) error {
	methodName := toGoName(operation.Name)

	// Check if this is a one-way operation (no output message)
	isOneWay := operation.Output == nil

	// Get input and output message types
	inputType, outputType, err := g.getOperationTypes(operation)
	if err != nil {
		return fmt.Errorf("failed to get types for operation %s: %w", operation.Name, err)
	}

	// Get SOAP action
	soapAction := g.getSOAPActionForOperation(operation.Name, binding)

	// Generate method signature and documentation
	if operation.Documentation != "" {
		// Clean up documentation
		doc := strings.TrimSpace(operation.Documentation)
		doc = strings.ReplaceAll(doc, "\n", " ")
		file.P("// ", methodName, " ", doc)
	} else {
		if isOneWay {
			file.P("// ", methodName, " executes the ", operation.Name, " one-way SOAP operation.")
		} else {
			file.P("// ", methodName, " executes the ", operation.Name, " SOAP operation.")
		}
	}

	// Generate different method signatures for one-way vs request-response operations
	if isOneWay {
		// One-way operation: return only error
		file.P("func (c *Client) ", methodName, "(ctx ", file.QualifiedGoIdent(codegen.ContextIdent), ", req *", inputType, ", opts ...ClientOption) ", file.QualifiedGoIdent(codegen.ErrorIdent), " {")
		file.P("\treqEnvelope, err := ", file.QualifiedGoIdent(codegen.SOAPNewEnvelopeIdent), "(", file.QualifiedGoIdent(codegen.SOAPWithBodyIdent), "(req))")
		file.P("\tif err != nil {")
		file.P("\t\treturn ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to create SOAP envelope: %w\", err)")
		file.P("\t}")
		if soapAction != "" {
			file.P("\t_, err = c.Call(ctx, \"", soapAction, "\", reqEnvelope, opts...)")
		} else {
			file.P("\t_, err = c.Call(ctx, \"\", reqEnvelope, opts...)")
		}
		file.P("\tif err != nil {")
		file.P("\t\treturn ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"SOAP call failed: %w\", err)")
		file.P("\t}")
		file.P("\treturn nil")
	} else {
		// Request-response operation: return response and error
		file.P("func (c *Client) ", methodName, "(ctx ", file.QualifiedGoIdent(codegen.ContextIdent), ", req *", inputType, ", opts ...ClientOption) (*", outputType, ", ", file.QualifiedGoIdent(codegen.ErrorIdent), ") {")
		file.P("\treqEnvelope, err := ", file.QualifiedGoIdent(codegen.SOAPNewEnvelopeIdent), "(", file.QualifiedGoIdent(codegen.SOAPWithBodyIdent), "(req))")
		file.P("\tif err != nil {")
		file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to create SOAP envelope: %w\", err)")
		file.P("\t}")
		if soapAction != "" {
			file.P("\trespEnvelope, err := c.Call(ctx, \"", soapAction, "\", reqEnvelope, opts...)")
		} else {
			file.P("\trespEnvelope, err := c.Call(ctx, \"\", reqEnvelope, opts...)")
		}
		file.P("\tif err != nil {")
		file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"SOAP call failed: %w\", err)")
		file.P("\t}")
		file.P("\tvar result ", outputType)
		file.P("\tif err := ", file.QualifiedGoIdent(codegen.XMLUnmarshalIdent), "(respEnvelope.Body.Content, &result); err != nil {")
		file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to unmarshal response body: %w\", err)")
		file.P("\t}")
		file.P("\treturn &result, nil")
	}
	file.P("}")
	file.P()

	return nil
}

// getOperationTypes determines the input and output types for an operation
func (g *Generator) getOperationTypes(operation *wsdl.Operation) (inputType, outputType string, err error) {
	// Get input type
	if operation.Input != nil {
		inputType, err = g.getMessageElementType(operation.Input.Message)
		if err != nil {
			return "", "", fmt.Errorf("failed to get input type: %w", err)
		}
	}

	// Get output type
	if operation.Output != nil {
		outputType, err = g.getMessageElementType(operation.Output.Message)
		if err != nil {
			return "", "", fmt.Errorf("failed to get output type: %w", err)
		}
	}

	// Provide default types if not found
	if inputType == "" {
		inputType = "interface{}"
	}
	if outputType == "" {
		// For operations without output messages, use an empty struct
		// This is more appropriate than interface{} for acknowledgment responses
		outputType = "struct{}"
	}

	return inputType, outputType, nil
}

// getMessageElementType gets the Go type name for a message element
func (g *Generator) getMessageElementType(messageName string) (string, error) {
	// Remove namespace prefix if present
	if colonIdx := strings.LastIndex(messageName, ":"); colonIdx != -1 {
		messageName = messageName[colonIdx+1:]
	}

	// Get the binding style for consistent naming
	bindingStyle := g.getBindingStyle()

	// Find the message definition
	for _, message := range g.definitions.Messages {
		if message.Name == messageName {
			// Get the element from the message part
			if len(message.Parts) > 0 {
				part := message.Parts[0]
				if part.Element != "" {
					// Extract element name
					elementName := part.Element
					if colonIdx := strings.LastIndex(elementName, ":"); colonIdx != -1 {
						elementName = elementName[colonIdx+1:]
					}

					// Use consistent type naming based on binding style
					return g.getConsistentTypeName(elementName, bindingStyle), nil
				}
			}
		}
	}

	return "", fmt.Errorf("message %s not found", messageName)
}

// getSOAPActionForOperation gets the SOAP action for an operation from binding
func (g *Generator) getSOAPActionForOperation(operationName string, binding *wsdl.Binding) string {
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
	return ""
}

// generateHelperFunctions generates helper types and functions for SOAP
func (g *Generator) generateHelperFunctions(file *codegen.File) {
	// Note: SOAP envelope types are now provided by the public API
	// No need to generate private types anymore
}
