package call

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/way-platform/soap-go/internal/soapcall"
)

// NewCommand creates a new [cobra.Command] for the call command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "call",
		Short:   "Call a SOAP action",
		GroupID: "network",
		Long: `Call a SOAP action with the specified endpoint, action, and XML payload.

By default, the payload file should contain only the request body content,
which will be automatically wrapped in a SOAP envelope. Use --full-envelope
if your payload file already contains a complete SOAP envelope.

The response will show only the body content by default. Use --output-envelope
to include the full SOAP envelope in the response.`,
		Example: `  # Basic usage - wrap payload in SOAP envelope
  soap call -e "http://example.com/service" -a "urn:GetWeather" -p request.xml

  # Full envelope mode
  soap call -e "http://example.com/service" -a "" -p envelope.xml --full-envelope

  # Include response envelope
  soap call -e "http://example.com/service" -a "urn:GetWeather" -p request.xml --output-envelope

  # With custom headers and timeout
  soap call -e "http://example.com/service" -a "urn:GetWeather" -p request.xml \
    --timeout 60s --headers "Authorization:Bearer token123"`,
	}

	// Required flags
	endpoint := cmd.Flags().StringP("endpoint", "e", "", "SOAP service endpoint URL (required)")
	_ = cmd.MarkFlagRequired("endpoint")

	action := cmd.Flags().StringP("action", "a", "", "SOAPAction header value (required)")
	_ = cmd.MarkFlagRequired("action")

	payload := cmd.Flags().StringP("payload", "p", "", "path to XML file containing the request payload (required)")
	_ = cmd.MarkFlagRequired("payload")
	_ = cmd.MarkFlagFilename("payload", "xml")

	// Optional flags
	fullEnvelope := cmd.Flags().Bool("full-envelope", false, "treat payload file as complete SOAP envelope")
	outputEnvelope := cmd.Flags().Bool("output-envelope", false, "include SOAP envelope in response output")
	timeout := cmd.Flags().Duration("timeout", 30*time.Second, "HTTP request timeout")
	headers := cmd.Flags().StringSlice("headers", nil, "additional HTTP headers in key:value format")
	insecure := cmd.Flags().Bool("insecure", false, "skip TLS certificate verification")
	output := cmd.Flags().StringP("output", "o", "-", "output file path (default: stdout)")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(config{
			endpoint:       *endpoint,
			action:         *action,
			payloadFile:    *payload,
			fullEnvelope:   *fullEnvelope,
			outputEnvelope: *outputEnvelope,
			timeout:        *timeout,
			headers:        *headers,
			insecure:       *insecure,
			outputFile:     *output,
		})
	}

	return cmd
}

type config struct {
	endpoint       string
	action         string
	payloadFile    string
	fullEnvelope   bool
	outputEnvelope bool
	timeout        time.Duration
	headers        []string
	insecure       bool
	outputFile     string
}

func run(cfg config) error {
	ctx := context.Background()

	// Read payload file
	payloadData, err := os.ReadFile(cfg.payloadFile)
	if err != nil {
		return fmt.Errorf("failed to read payload file %s: %w", cfg.payloadFile, err)
	}

	// Parse custom headers
	customHeaders, err := soapcall.ParseHeaders(cfg.headers)
	if err != nil {
		return fmt.Errorf("failed to parse headers: %w", err)
	}

	// Prepare XML for sending
	var xmlToSend []byte
	if cfg.fullEnvelope {
		// Use payload as-is if it's already a full envelope
		if !soapcall.IsFullEnvelope(payloadData) {
			return fmt.Errorf("payload file does not contain a valid SOAP envelope (use --full-envelope=false for raw payloads)")
		}
		xmlToSend = payloadData
	} else {
		// Wrap payload in SOAP envelope
		xmlToSend, err = soapcall.WrapInEnvelope(payloadData)
		if err != nil {
			return fmt.Errorf("failed to wrap payload in SOAP envelope: %w", err)
		}
	}

	// Add XML declaration
	xmlToSend = soapcall.AddXMLDeclaration(xmlToSend)

	// Create SOAP client
	client := soapcall.NewClient(soapcall.Config{
		Endpoint:   cfg.endpoint,
		Timeout:    cfg.timeout,
		Headers:    customHeaders,
		Insecure:   cfg.insecure,
		SOAPAction: cfg.action,
	})

	// Make the SOAP call
	responseData, err := client.Call(ctx, xmlToSend)
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}

	// Process response
	var outputData []byte
	if cfg.outputEnvelope {
		// Output the full response as-is
		outputData = responseData
	} else {
		// Extract body content from SOAP envelope
		outputData, err = soapcall.ExtractFromEnvelope(responseData)
		if err != nil {
			return fmt.Errorf("failed to process SOAP response: %w", err)
		}
	}

	// Format XML for better readability
	formattedOutput, err := soapcall.FormatXML(outputData)
	if err != nil {
		// If formatting fails, use original data
		formattedOutput = outputData
	}

	// Add XML declaration to output
	formattedOutput = soapcall.AddXMLDeclaration(formattedOutput)

	// Write output
	if cfg.outputFile == "-" {
		fmt.Print(string(formattedOutput))
	} else {
		if err := os.WriteFile(cfg.outputFile, formattedOutput, 0o644); err != nil {
			return fmt.Errorf("failed to write output file %s: %w", cfg.outputFile, err)
		}
		fmt.Printf("Response written to %s\n", cfg.outputFile)
	}

	return nil
}
