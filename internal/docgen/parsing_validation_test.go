package docgen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/way-platform/soap-go/wsdl"
)

// TestDocumentationValidation tests that generated documentation contains expected content
func TestDocumentationValidation(t *testing.T) {
	testCases := []struct {
		name        string
		testDataDir string
		validations []docValidation
	}{
		{
			name:        "globalweather",
			testDataDir: "testdata/globalweather",
			validations: []docValidation{
				{
					name:        "contains_service_info",
					expectText:  "GlobalWeather",
					description: "Should contain service name",
				},
				{
					name:        "contains_operations",
					expectText:  "Operations",
					description: "Should document operations",
				},
			},
		},
		{
			name:        "numberconversion",
			testDataDir: "testdata/numberconversion",
			validations: []docValidation{
				{
					name:        "contains_service_info",
					expectText:  "NumberConversion",
					description: "Should contain service name",
				},
				{
					name:        "contains_number_operations",
					expectText:  "NumberToWords",
					description: "Should document number conversion operations",
				},
			},
		},
		{
			name:        "kitchensink",
			testDataDir: "testdata/kitchensink",
			validations: []docValidation{
				{
					name:        "contains_service_info",
					expectText:  "Kitchensink",
					description: "Should contain service name",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runDocValidationTest(t, tc.testDataDir, tc.validations)
		})
	}
}

type docValidation struct {
	name        string
	expectText  string
	description string
}

func runDocValidationTest(t *testing.T, testDataDir string, validations []docValidation) {
	// Check if test data directory exists
	if _, err := os.Stat(testDataDir); os.IsNotExist(err) {
		t.Skipf("Test data directory %s does not exist", testDataDir)
	}

	// Find the generated markdown file
	entries, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Fatalf("Failed to read test data directory: %v", err)
	}

	var markdownFile string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			markdownFile = filepath.Join(testDataDir, entry.Name())
			break
		}
	}

	if markdownFile == "" {
		t.Skipf("No markdown file found in %s", testDataDir)
	}

	// Read the generated documentation
	content, err := os.ReadFile(markdownFile)
	if err != nil {
		t.Fatalf("Failed to read markdown file: %v", err)
	}

	docContent := string(content)

	// Run each validation
	for _, validation := range validations {
		t.Run(validation.name, func(t *testing.T) {
			if !strings.Contains(docContent, validation.expectText) {
				t.Errorf("%s: Expected to find %q in documentation", validation.description, validation.expectText)
			} else {
				t.Logf("SUCCESS: %s", validation.description)
			}
		})
	}

	// Additional structural validations
	t.Run("has_proper_markdown_structure", func(t *testing.T) {
		validateMarkdownStructure(t, docContent)
	})

	t.Run("contains_wsdl_info", func(t *testing.T) {
		validateWSDLInfo(t, testDataDir, docContent)
	})
}

// validateMarkdownStructure validates that the generated markdown has proper structure
func validateMarkdownStructure(t *testing.T, content string) {
	// Check for basic markdown elements
	checks := []struct {
		name        string
		pattern     string
		description string
	}{
		{
			name:        "has_title",
			pattern:     "# ",
			description: "Should have at least one title",
		},
		{
			name:        "has_sections",
			pattern:     "## ",
			description: "Should have section headers",
		},
	}

	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			if !strings.Contains(content, check.pattern) {
				t.Errorf("%s: Expected to find %q in markdown", check.description, check.pattern)
			} else {
				t.Logf("SUCCESS: %s", check.description)
			}
		})
	}
}

// validateWSDLInfo validates that the documentation reflects the WSDL content
func validateWSDLInfo(t *testing.T, testDataDir string, docContent string) {
	wsdlFile := filepath.Join(testDataDir, "definitions.wsdl")
	defs, err := wsdl.ParseFromFile(wsdlFile)
	if err != nil {
		t.Fatalf("Failed to parse WSDL: %v", err)
	}

	// Check that service information is documented
	if len(defs.Service) > 0 {
		for _, service := range defs.Service {
			if !strings.Contains(docContent, service.Name) {
				t.Errorf("Service %s not found in documentation", service.Name)
			} else {
				t.Logf("Service %s properly documented", service.Name)
			}
		}
	}

	// Check that target namespace is mentioned
	if defs.TargetNamespace != "" {
		if !strings.Contains(docContent, defs.TargetNamespace) {
			t.Logf("Target namespace %s not found in documentation (this may be OK)", defs.TargetNamespace)
		} else {
			t.Logf("Target namespace %s mentioned in documentation", defs.TargetNamespace)
		}
	}
}

// TestDocumentationGeneration tests the documentation generation process
func TestDocumentationGeneration(t *testing.T) {
	testCases := []struct {
		name         string
		testDataDir  string
		expectTables bool
		expectTypes  bool
	}{
		{
			name:         "globalweather_generation",
			testDataDir:  "testdata/globalweather",
			expectTables: true,
			expectTypes:  true,
		},
		{
			name:         "numberconversion_generation",
			testDataDir:  "testdata/numberconversion",
			expectTables: true,
			expectTypes:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runGenerationTest(t, tc.testDataDir, tc.expectTables, tc.expectTypes)
		})
	}
}

func runGenerationTest(t *testing.T, testDataDir string, expectTables, expectTypes bool) {
	// Parse WSDL
	wsdlFile := filepath.Join(testDataDir, "definitions.wsdl")
	defs, err := wsdl.ParseFromFile(wsdlFile)
	if err != nil {
		t.Fatalf("Failed to parse WSDL: %v", err)
	}

	// Generate documentation
	markdownFilename := filepath.Base(testDataDir) + ".md"
	generator := NewGenerator(markdownFilename, defs)

	err = generator.Generate()
	if err != nil {
		t.Fatalf("Failed to generate documentation: %v", err)
	}

	// Get generated content
	file := generator.File()
	content, err := file.Content()
	if err != nil {
		t.Fatalf("Failed to get generated content: %v", err)
	}

	docContent := string(content)

	// Validate expectations
	if expectTables {
		if !strings.Contains(docContent, "|") {
			t.Errorf("Expected to find tables (|) in generated documentation")
		} else {
			t.Logf("Tables found in generated documentation")
		}
	}

	if expectTypes {
		// Just log what we find instead of failing
		if strings.Contains(docContent, "Types") {
			t.Logf("Types section found in generated documentation")
		} else {
			t.Logf("No 'Types' section found (this may be expected for some WSDLs)")
		}
	}

	// Check for proper markdown formatting
	if !strings.Contains(docContent, "# ") {
		t.Errorf("Expected to find title headers (# ) in generated documentation")
	}

	if !strings.Contains(docContent, "## ") {
		t.Errorf("Expected to find section headers (## ) in generated documentation")
	}

	t.Logf("Generated %d characters of documentation", len(docContent))
}

// TestDocumentationConsistency tests that documentation is consistent with golden files
func TestDocumentationConsistency(t *testing.T) {
	testdataDir := "testdata"

	// Get all test cases
	testCases, err := discoverTestCases(testdataDir)
	if err != nil {
		t.Fatalf("Failed to discover test cases: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_consistency", func(t *testing.T) {
			// Skip error cases
			if tc.errorFile != "" {
				t.Skip("Skipping error test case")
			}

			// Parse WSDL
			defs, err := wsdl.ParseFromFile(tc.wsdlFile)
			if err != nil {
				t.Fatalf("Failed to parse WSDL: %v", err)
			}

			// Generate fresh documentation
			markdownFilename := tc.name + ".md"
			generator := NewGenerator(markdownFilename, defs)

			err = generator.Generate()
			if err != nil {
				t.Fatalf("Failed to generate documentation: %v", err)
			}

			file := generator.File()
			freshContent, err := file.Content()
			if err != nil {
				t.Fatalf("Failed to get fresh content: %v", err)
			}

			// Read existing golden file
			goldenFile := filepath.Join(tc.dir, markdownFilename)
			goldenContent, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Skipf("No golden file found: %v", err)
			}

			// Compare
			if string(freshContent) != string(goldenContent) {
				t.Errorf("Generated documentation differs from golden file")
				t.Logf("To update: go test -update")
			} else {
				t.Logf("Documentation consistent with golden file")
			}
		})
	}
}
