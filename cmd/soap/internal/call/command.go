package call

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/way-platform/soap-go"
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
  soap call -e "http://example.com/service" -a "urn:GetWeather" -p request.xml --output-envelope`,
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
	output := cmd.Flags().StringP("output", "o", "-", "output file path (default: stdout)")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		debug, _ := cmd.Root().PersistentFlags().GetBool("debug")
		return run(config{
			endpoint:       *endpoint,
			action:         *action,
			payloadFile:    *payload,
			fullEnvelope:   *fullEnvelope,
			outputEnvelope: *outputEnvelope,
			outputFile:     *output,
			debug:          debug,
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
	outputFile     string
	debug          bool
}

func run(cfg config) error {
	ctx := context.Background()

	// Read payload file
	payloadData, err := os.ReadFile(cfg.payloadFile)
	if err != nil {
		return fmt.Errorf("failed to read payload file %s: %w", cfg.payloadFile, err)
	}

	// No custom headers parsing needed - action will be handled by the Call method

	// Create SOAP client
	client, err := soap.NewClient(
		soap.WithEndpoint(cfg.endpoint),
		soap.WithDebug(cfg.debug),
	)
	if err != nil {
		return fmt.Errorf("failed to create SOAP client: %w", err)
	}

	// Prepare request envelope
	var requestEnvelope *soap.Envelope
	if cfg.fullEnvelope {
		// Parse the full envelope from the payload
		requestEnvelope = &soap.Envelope{}
		if err := xml.Unmarshal(payloadData, requestEnvelope); err != nil {
			return fmt.Errorf("failed to parse SOAP envelope: %w", err)
		}
	} else {
		// Wrap payload in SOAP envelope
		requestEnvelope, err = soap.NewEnvelope(soap.WithBody(payloadData))
		if err != nil {
			return fmt.Errorf("failed to create SOAP envelope: %w", err)
		}
	}

	// Make the SOAP call
	responseEnvelope, err := client.Call(ctx, cfg.action, requestEnvelope, soap.WithEndpoint(cfg.endpoint))
	if err != nil {
		return fmt.Errorf("SOAP call failed: %w", err)
	}

	// Process response
	var outputData []byte
	if cfg.outputEnvelope {
		// Output the full response envelope as XML
		var err error
		outputData, err = xml.Marshal(responseEnvelope)
		if err != nil {
			return fmt.Errorf("failed to marshal response envelope: %w", err)
		}
	} else {
		// Extract body content from SOAP envelope
		outputData = responseEnvelope.Body.Content
	}
	// Write output
	if cfg.outputFile == "-" {
		fmt.Print(string(outputData))
	} else {
		if err := os.WriteFile(cfg.outputFile, outputData, 0o644); err != nil {
			return fmt.Errorf("failed to write output file %s: %w", cfg.outputFile, err)
		}
		fmt.Printf("Response written to %s\n", cfg.outputFile)
	}
	return nil
}
