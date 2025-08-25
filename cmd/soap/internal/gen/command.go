package gen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/way-platform/soap-go/internal/soapgen"
	"github.com/way-platform/soap-go/wsdl"
)

// NewCommand creates a new [cobra.Command] for the gen command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen",
		Short:   "Generate code for a SOAP API",
		GroupID: "gen",
	}
	inputFile := cmd.Flags().StringP("input", "i", "", "input WSDL file (required)")
	_ = cmd.MarkFlagRequired("input")
	outputDir := cmd.Flags().StringP("dir", "d", "", "output directory (required)")
	_ = cmd.MarkFlagRequired("dir")
	packageName := cmd.Flags().StringP("package", "p", "", "Go package name (required)")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(config{
			inputFile:   *inputFile,
			outputDir:   *outputDir,
			packageName: *packageName,
		})
	}
	return cmd
}

type config struct {
	inputFile   string
	outputDir   string
	packageName string
}

func run(cfg config) error {
	if cfg.packageName == "" {
		cfg.packageName = filepath.Base(cfg.outputDir)
	}
	// Parse the WSDL file
	defs, err := wsdl.ParseFromFile(cfg.inputFile)
	if err != nil {
		return fmt.Errorf("failed to parse WSDL file: %w", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(cfg.outputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create generator with configuration
	generator := soapgen.NewGenerator(defs, soapgen.Config{
		PackageName: cfg.packageName,
	})

	// Generate the code
	if err := generator.Generate(); err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write all generated files
	for _, file := range generator.Files() {
		content, err := file.Content()
		if err != nil {
			return fmt.Errorf("failed to generate content: %w", err)
		}

		// Write file to output directory
		outputPath := filepath.Join(cfg.outputDir, file.Filename())
		if err := os.WriteFile(outputPath, content, 0o644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		fmt.Printf("Generated %s\n", outputPath)
	}

	return nil
}
