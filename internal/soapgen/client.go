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

// generateClientOptions generates the ClientOption types and implementations
func (g *Generator) generateClientOptions(file *codegen.File) {
	file.P("// ClientOption configures a Client.")
	file.P("type ClientOption func(*clientConfig)")
	file.P()

	file.P("// clientConfig holds the configuration for a Client.")
	file.P("type clientConfig struct {")
	file.P("\thttpClient *", file.QualifiedGoIdent(codegen.HTTPClientIdent))
	file.P("\tendpoint   ", file.QualifiedGoIdent(codegen.StringIdent))
	file.P("}")
	file.P()

	file.P("// WithHTTPClient sets a custom HTTP client for the SOAP client.")
	file.P("func WithHTTPClient(client *", file.QualifiedGoIdent(codegen.HTTPClientIdent), ") ClientOption {")
	file.P("\treturn func(c *clientConfig) {")
	file.P("\t\tc.httpClient = client")
	file.P("\t}")
	file.P("}")
	file.P()

	file.P("// WithEndpoint sets the SOAP endpoint URL.")
	file.P("func WithEndpoint(endpoint ", file.QualifiedGoIdent(codegen.StringIdent), ") ClientOption {")
	file.P("\treturn func(c *clientConfig) {")
	file.P("\t\tc.endpoint = endpoint")
	file.P("\t}")
	file.P("}")
	file.P()
}

// generateClientStruct generates the main Client struct
func (g *Generator) generateClientStruct(file *codegen.File) {
	file.P("// Client is a SOAP client for this service.")
	file.P("type Client struct {")
	file.P("\thttpClient *", file.QualifiedGoIdent(codegen.HTTPClientIdent))
	file.P("\tendpoint   ", file.QualifiedGoIdent(codegen.StringIdent))
	file.P("}")
	file.P()
}

// generateNewClientFunction generates the NewClient constructor
func (g *Generator) generateNewClientFunction(file *codegen.File) {
	// Extract default endpoint from service definitions
	endpoint := g.getDefaultEndpoint()

	file.P("// NewClient creates a new SOAP client.")
	file.P("func NewClient(opts ...ClientOption) (*Client, error) {")
	file.P("\tconfig := &clientConfig{")
	file.P("\t\thttpClient: ", file.QualifiedGoIdent(codegen.GoIdent{GoImportPath: "net/http", GoName: "DefaultClient"}), ",")
	if endpoint != "" {
		file.P("\t\tendpoint:   \"", endpoint, "\",")
	} else {
		file.P("\t\tendpoint:   \"\", // Set endpoint using WithEndpoint option")
	}
	file.P("\t}")
	file.P()
	file.P("\tfor _, opt := range opts {")
	file.P("\t\topt(config)")
	file.P("\t}")
	file.P()
	file.P("\t// Validate that we have an endpoint")
	file.P("\tif config.endpoint == \"\" {")
	if endpoint == "" {
		file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"SOAP endpoint must be provided using WithEndpoint() - no default endpoint found in WSDL\")")
	} else {
		file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"SOAP endpoint is required\")")
	}
	file.P("\t}")
	file.P()
	file.P("\treturn &Client{")
	file.P("\t\thttpClient: config.httpClient,")
	file.P("\t\tendpoint:   config.endpoint,")
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
		file.P("// ", methodName, " executes the ", operation.Name, " SOAP operation.")
	}

	file.P("func (c *Client) ", methodName, "(ctx ", file.QualifiedGoIdent(codegen.ContextIdent), ", req *", inputType, ") (*", outputType, ", ", file.QualifiedGoIdent(codegen.ErrorIdent), ") {")

	file.P("\t// Marshal request to XML")
	file.P("\treqXML, err := ", file.QualifiedGoIdent(codegen.XMLMarshalIdent), "(req)")
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to marshal request: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\t// Create SOAP envelope")
	file.P("\tenvelope := &soapEnvelope{")
	file.P("\t\tXMLNS: \"http://schemas.xmlsoap.org/soap/envelope/\",")
	file.P("\t\tBody: soapBody{Content: reqXML},")
	file.P("\t}")
	file.P()

	file.P("\t// Marshal envelope to XML")
	file.P("\txmlData, err := ", file.QualifiedGoIdent(codegen.XMLMarshalIdent), "(envelope)")
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to marshal SOAP envelope: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\t// Create HTTP request")
	file.P("\thttpReq, err := ", file.QualifiedGoIdent(codegen.HTTPNewRequestWithContextIdent), "(ctx, \"POST\", c.endpoint, ", file.QualifiedGoIdent(codegen.BytesNewReaderIdent), "(xmlData))")
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to create HTTP request: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\t// Set headers")
	file.P("\thttpReq.Header.Set(\"Content-Type\", \"text/xml; charset=utf-8\")")
	if soapAction != "" {
		file.P("\thttpReq.Header.Set(\"SOAPAction\", \"", soapAction, "\")")
	}
	file.P()

	file.P("\t// Execute request")
	file.P("\tresp, err := c.httpClient.Do(httpReq)")
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to execute HTTP request: %w\", err)")
	file.P("\t}")
	file.P("\tdefer resp.Body.Close()")
	file.P()

	file.P("\t// Read response")
	file.P("\trespBody, err := ", file.QualifiedGoIdent(codegen.IOReadAllIdent), "(resp.Body)")
	file.P("\tif err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to read response body: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\t// Check for HTTP errors")
	file.P("\tif resp.StatusCode != ", file.QualifiedGoIdent(codegen.HTTPStatusOKIdent), " {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"HTTP error %d: %s\", resp.StatusCode, string(respBody))")
	file.P("\t}")
	file.P()

	file.P("\t// Parse SOAP response")
	file.P("\tvar respEnvelope soapEnvelope")
	file.P("\tif err := ", file.QualifiedGoIdent(codegen.XMLUnmarshalIdent), "(respBody, &respEnvelope); err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to unmarshal SOAP response: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\t// Extract response from SOAP body")
	file.P("\tvar result ", outputType)
	file.P("\tif err := ", file.QualifiedGoIdent(codegen.XMLUnmarshalIdent), "(respEnvelope.Body.Content, &result); err != nil {")
	file.P("\t\treturn nil, ", file.QualifiedGoIdent(codegen.FmtErrorfIdent), "(\"failed to unmarshal response body: %w\", err)")
	file.P("\t}")
	file.P()

	file.P("\treturn &result, nil")
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
		outputType = "interface{}"
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
	// Generate private SOAP envelope types
	file.P("// soapEnvelope represents a SOAP envelope.")
	file.P("type soapEnvelope struct {")
	file.P("\tXMLName ", file.QualifiedGoIdent(codegen.XMLNameIdent), " `xml:\"soap:Envelope\"`")
	file.P("\tXMLNS   string   `xml:\"xmlns:soap,attr\"`")
	file.P("\tBody    soapBody `xml:\"soap:Body\"`")
	file.P("}")
	file.P()

	file.P("// soapBody represents a SOAP body.")
	file.P("type soapBody struct {")
	file.P("\tContent []byte `xml:\",innerxml\"`")
	file.P("}")
	file.P()
}
