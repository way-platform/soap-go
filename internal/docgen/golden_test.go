package docgen

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
		// Some test cases might have parsing errors
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

	// Create generator with markdown filename
	markdownFilename := tc.name + ".md"
	generator := NewGenerator(markdownFilename, defs)

	// Generate documentation
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

	// Get generated markdown content
	file := generator.File()
	content, err := file.Content()
	if err != nil {
		t.Fatalf("Failed to get content for file %s: %v", file.Filename(), err)
	}

	generatedContent := string(content)

	// If update flag is set, write the golden file and we're done
	if *update {
		if err := updateGoldenFile(tc.dir, markdownFilename, generatedContent); err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Logf("Updated golden file for test case: %s", tc.name)
		return
	}

	// Compare generated content with golden file
	if err := compareWithGolden(t, tc.dir, markdownFilename, generatedContent); err != nil {
		t.Error(err)
	}
}

func updateGoldenFile(dir, filename, content string) error {
	// Clean up existing markdown files in the test directory
	if err := cleanMarkdownFiles(dir); err != nil {
		return fmt.Errorf("failed to clean markdown files: %w", err)
	}

	// Write new golden file
	filePath := filepath.Join(dir, filename)
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("failed to write golden file %s: %w", filePath, err)
	}

	return nil
}

func cleanMarkdownFiles(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			filePath := filepath.Join(dir, entry.Name())
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to remove %s: %w", filePath, err)
			}
		}
	}

	return nil
}

func compareWithGolden(t *testing.T, dir, filename, generatedContent string) error {
	// Read existing golden file
	goldenFilePath := filepath.Join(dir, filename)
	goldenContent, err := os.ReadFile(goldenFilePath)
	if err != nil {
		return fmt.Errorf("failed to read golden file %s: %w", goldenFilePath, err)
	}

	// Compare the content using cmp.Diff
	if diff := cmp.Diff(string(goldenContent), generatedContent); diff != "" {
		return fmt.Errorf("golden file mismatch (-golden +generated):\n%s", diff)
	}

	return nil
}
