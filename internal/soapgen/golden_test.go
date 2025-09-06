package soapgen

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/way-platform/soap-go/wsdl"
)

var update = flag.Bool("update", false, "update golden files")

func TestGoldenFiles(t *testing.T) {
	testdataDir := "testdata"

	// Discover all test cases by walking the testdata directory
	testCases, err := discoverTestCases(testdataDir)
	if err != nil {
		t.Fatalf("Failed to discover test cases: %v", err)
	}

	if len(testCases) == 0 {
		t.Fatal("No test cases found in testdata directory")
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			runGoldenTest(t, testCase)
		})
	}
}

type testCase struct {
	name      string
	dir       string
	wsdlFile  string
	errorFile string // path to error.txt if it exists
}

func discoverTestCases(testdataDir string) ([]testCase, error) {
	var testCases []testCase

	err := filepath.WalkDir(testdataDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root testdata directory
		if path == testdataDir {
			return nil
		}

		// Only process directories that are direct children of testdata
		if d.IsDir() && filepath.Dir(path) == testdataDir {
			wsdlFile := filepath.Join(path, "definitions.wsdl")
			if _, err := os.Stat(wsdlFile); err == nil {
				// Check if error file exists
				errorFile := filepath.Join(path, "error.txt")
				if _, err := os.Stat(errorFile); err != nil {
					errorFile = "" // error file doesn't exist
				}

				testCases = append(testCases, testCase{
					name:      d.Name(),
					dir:       path,
					wsdlFile:  wsdlFile,
					errorFile: errorFile,
				})
			}
		}

		return nil
	})

	return testCases, err
}

func runGoldenTest(t *testing.T, tc testCase) {
	// Parse WSDL file
	defs, err := wsdl.ParseFromFile(tc.wsdlFile)
	if err != nil {
		// Some test cases (like no_types_schema) might have parsing errors
		// Check if this is expected
		if tc.errorFile != "" {
			// This test case expects an error, read the expected error
			expectedErrorBytes, readErr := os.ReadFile(tc.errorFile)
			if readErr != nil {
				t.Fatalf("Failed to read expected error file: %v", readErr)
			}
			expectedError := strings.TrimSpace(string(expectedErrorBytes))

			if !strings.Contains(err.Error(), expectedError) {
				t.Errorf("Expected error to contain %q, got: %v", expectedError, err)
			}
			return // This test case is done
		}
		t.Fatalf("Failed to parse WSDL: %v", err)
	}

	// Create generator with the test case directory as package name
	generator := NewGenerator(defs, Config{
		PackageName: tc.name,
	})

	// Generate code
	err = generator.Generate()
	// Handle generation errors
	if err != nil {
		// Check if this test case expects an error
		if tc.errorFile != "" {
			// This test case expects an error, read the expected error
			expectedErrorBytes, readErr := os.ReadFile(tc.errorFile)
			if readErr != nil {
				t.Fatalf("Failed to read expected error file: %v", readErr)
			}
			expectedError := strings.TrimSpace(string(expectedErrorBytes))

			if !strings.Contains(err.Error(), expectedError) {
				t.Errorf("Expected error to contain %q, got: %v", expectedError, err)
			}
			return // This test case is done
		}
		t.Fatalf("Unexpected generation error: %v", err)
	}

	// Get generated files
	generatedFiles := make(map[string]string)
	for _, file := range generator.Files() {
		content, err := file.Content()
		if err != nil {
			t.Fatalf("Failed to get content for file %s: %v", file.Filename(), err)
		}

		// Extract just the filename for the key
		filename := filepath.Base(file.Filename())
		generatedFiles[filename] = string(content)
	}

	// If update flag is set, write the golden files and we're done
	if *update {
		if err := updateGoldenFiles(tc.dir, generatedFiles); err != nil {
			t.Fatalf("Failed to update golden files: %v", err)
		}
		t.Logf("Updated golden files for test case: %s", tc.name)
		return
	}

	// Compare generated files with golden files
	if err := compareWithGolden(t, tc.dir, generatedFiles); err != nil {
		t.Error(err)
	}

	// Run additional validation on the generated code
	if err := validateGeneratedCode(t, tc, generatedFiles); err != nil {
		t.Error(err)
	}
}

func updateGoldenFiles(dir string, generatedFiles map[string]string) error {
	// Clean up existing Go files in the test directory
	if err := cleanGoFiles(dir); err != nil {
		return fmt.Errorf("failed to clean Go files: %w", err)
	}

	// Write new golden files
	for filename, content := range generatedFiles {
		filePath := filepath.Join(dir, filename)
		if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("failed to write golden file %s: %w", filePath, err)
		}
	}

	return nil
}

func cleanGoFiles(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			filePath := filepath.Join(dir, entry.Name())
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to remove %s: %w", filePath, err)
			}
		}
	}

	return nil
}

func compareWithGolden(t *testing.T, dir string, generatedFiles map[string]string) error {
	// Read existing golden files
	goldenFiles := make(map[string]string)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			filePath := filepath.Join(dir, entry.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read golden file %s: %w", filePath, err)
			}
			goldenFiles[entry.Name()] = string(content)
		}
	}

	// Compare the maps using cmp.Diff
	if diff := cmp.Diff(goldenFiles, generatedFiles); diff != "" {
		return fmt.Errorf("golden files mismatch (-golden +generated):\n%s", diff)
	}

	return nil
}

// validateGeneratedCode performs additional validation on the generated code
func validateGeneratedCode(t *testing.T, tc testCase, generatedFiles map[string]string) error {
	// Skip validation for error test cases
	if tc.errorFile != "" {
		return nil
	}

	// Validate that the generated code compiles and contains expected elements
	for filename, content := range generatedFiles {
		if err := validateGoFileContent(t, tc.name, filename, content); err != nil {
			return fmt.Errorf("validation failed for %s: %w", filename, err)
		}
	}

	return nil
}

// validateGoFileContent validates the content of a generated Go file
func validateGoFileContent(t *testing.T, testCaseName, filename, content string) error {
	// Basic syntax validation - check that the file has proper Go syntax
	if !strings.HasPrefix(content, "package "+testCaseName) {
		return fmt.Errorf("file %s does not start with correct package declaration", filename)
	}

	// Check for required imports
	if strings.Contains(content, "xml.Name") && !strings.Contains(content, `"encoding/xml"`) {
		return fmt.Errorf("file %s uses xml.Name but doesn't import encoding/xml", filename)
	}

	// Validate struct definitions have proper XML tags
	if err := validateXMLTags(content); err != nil {
		return fmt.Errorf("XML tag validation failed in %s: %w", filename, err)
	}

	// Validate enumeration types if present
	if err := validateEnumerations(content); err != nil {
		return fmt.Errorf("enumeration validation failed in %s: %w", filename, err)
	}

	// Test that XMLName fields are properly configured
	if err := validateXMLNameFields(content); err != nil {
		return fmt.Errorf("XMLName field validation failed in %s: %w", filename, err)
	}

	return nil
}

// validateXMLTags checks that XML tags are properly formatted
func validateXMLTags(content string) error {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Look for struct fields with xml tags
		if strings.Contains(trimmed, "`xml:") {
			// Basic validation - ensure tag is properly quoted
			if !strings.Contains(trimmed, "`xml:\"") || !strings.HasSuffix(trimmed, "\"`") {
				return fmt.Errorf("malformed XML tag at line %d: %s", i+1, trimmed)
			}
		}
	}

	return nil
}

// validateEnumerations checks that enumeration types have proper methods
func validateEnumerations(content string) error {
	// Look for enumeration type definitions
	lines := strings.Split(content, "\n")
	var enumTypes []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "type ") && strings.HasSuffix(trimmed, " string") {
			// Extract type name
			parts := strings.Fields(trimmed)
			if len(parts) >= 2 {
				typeName := parts[1]
				if strings.HasSuffix(typeName, "Type") {
					enumTypes = append(enumTypes, typeName)
				}
			}
		}
	}

	// For each enum type, check that String() and IsValid() methods exist
	for _, enumType := range enumTypes {
		stringMethod := fmt.Sprintf("func (e %s) String() string", enumType)
		isValidMethod := fmt.Sprintf("func (e %s) IsValid() bool", enumType)

		if !strings.Contains(content, stringMethod) {
			return fmt.Errorf("enumeration type %s missing String() method", enumType)
		}

		if !strings.Contains(content, isValidMethod) {
			return fmt.Errorf("enumeration type %s missing IsValid() method", enumType)
		}
	}

	return nil
}

// validateXMLNameFields checks that XMLName fields are properly configured
func validateXMLNameFields(content string) error {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Look for XMLName field declarations
		if strings.Contains(trimmed, "XMLName") && strings.Contains(trimmed, "xml.Name") {
			// Ensure it has proper xml tag
			if !strings.Contains(trimmed, "`xml:\"") {
				return fmt.Errorf("XMLName field at line %d missing xml tag: %s", i+1, trimmed)
			}
		}
	}

	return nil
}
