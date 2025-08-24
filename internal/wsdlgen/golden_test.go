package wsdlgen

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
