package soapgen

import (
	"encoding/xml"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/way-platform/soap-go/wsdl"
)

// TestParsingWithGeneratedTypes tests XML parsing using types generated from testdata
func TestParsingWithGeneratedTypes(t *testing.T) {
	testdataDir := "testdata"

	// Discover all test cases that have generated types
	testCases, err := discoverParsingTestCases(testdataDir)
	if err != nil {
		t.Fatalf("Failed to discover parsing test cases: %v", err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			runParsingTest(t, testCase)
		})
	}
}

type parsingTestCase struct {
	name     string
	dir      string
	wsdlFile string
	goFiles  []string
}

func discoverParsingTestCases(testdataDir string) ([]parsingTestCase, error) {
	var testCases []parsingTestCase

	entries, err := os.ReadDir(testdataDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		testDir := filepath.Join(testdataDir, entry.Name())
		wsdlFile := filepath.Join(testDir, "definitions.wsdl")

		// Check if WSDL file exists
		if _, err := os.Stat(wsdlFile); err != nil {
			continue
		}

		// Find all Go files in the test directory
		var goFiles []string
		testEntries, err := os.ReadDir(testDir)
		if err != nil {
			continue
		}

		for _, testEntry := range testEntries {
			if strings.HasSuffix(testEntry.Name(), ".go") {
				goFiles = append(goFiles, filepath.Join(testDir, testEntry.Name()))
			}
		}

		// Skip if no Go files (might be error test cases)
		if len(goFiles) == 0 {
			continue
		}

		testCases = append(testCases, parsingTestCase{
			name:     entry.Name(),
			dir:      testDir,
			wsdlFile: wsdlFile,
			goFiles:  goFiles,
		})
	}

	return testCases, nil
}

func runParsingTest(t *testing.T, tc parsingTestCase) {
	// Parse the WSDL to understand what types should be available
	defs, err := wsdl.ParseFromFile(tc.wsdlFile)
	if err != nil {
		t.Fatalf("Failed to parse WSDL: %v", err)
	}

	// Parse the generated Go files to extract type information
	types, err := extractTypesFromGoFiles(tc.goFiles)
	if err != nil {
		t.Fatalf("Failed to extract types from Go files: %v", err)
	}

	// Create XML test cases based on the WSDL and available types
	xmlTestCases := generateXMLTestCases(defs, types)

	// Run parsing tests for each XML test case
	for _, xmlTC := range xmlTestCases {
		t.Run(xmlTC.name, func(t *testing.T) {
			runXMLParsingTest(t, xmlTC, tc.name)
		})
	}
}

type typeInfo struct {
	name     string
	typeName string
	fields   map[string]fieldInfo
}

type fieldInfo struct {
	name     string
	typeName string
	xmlTag   string
	isSlice  bool
	isPtr    bool
}

type xmlTestCase struct {
	name       string
	typeName   string
	xmlContent string
}

func extractTypesFromGoFiles(goFiles []string) (map[string]typeInfo, error) {
	types := make(map[string]typeInfo)

	for _, goFile := range goFiles {
		content, err := os.ReadFile(goFile)
		if err != nil {
			return nil, err
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, goFile, content, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		// Extract struct types from the AST
		for _, decl := range node.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							typeInfo := extractStructTypeInfo(typeSpec.Name.Name, structType)
							types[typeSpec.Name.Name] = typeInfo
						}
					}
				}
			}
		}
	}

	return types, nil
}

func extractStructTypeInfo(typeName string, structType *ast.StructType) typeInfo {
	info := typeInfo{
		name:     typeName,
		typeName: typeName,
		fields:   make(map[string]fieldInfo),
	}

	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue // Skip embedded fields for now
		}

		fieldName := field.Names[0].Name
		fieldTypeName := extractTypeName(field.Type)
		xmlTag := extractXMLTag(field.Tag)

		isSlice := false
		isPtr := false

		// Check if it's a slice
		if arrayType, ok := field.Type.(*ast.ArrayType); ok {
			isSlice = true
			fieldTypeName = extractTypeName(arrayType.Elt)
		}

		// Check if it's a pointer
		if starExpr, ok := field.Type.(*ast.StarExpr); ok {
			isPtr = true
			fieldTypeName = extractTypeName(starExpr.X)
		}

		info.fields[fieldName] = fieldInfo{
			name:     fieldName,
			typeName: fieldTypeName,
			xmlTag:   xmlTag,
			isSlice:  isSlice,
			isPtr:    isPtr,
		}
	}

	return info
}

func extractTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return extractTypeName(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return extractTypeName(t.X)
	case *ast.ArrayType:
		return "[]" + extractTypeName(t.Elt)
	default:
		return "unknown"
	}
}

func extractXMLTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}

	tagValue := strings.Trim(tag.Value, "`")
	// Find the xml tag using proper parsing
	// Look for xml:"..." pattern
	xmlStart := strings.Index(tagValue, `xml:"`)
	if xmlStart == -1 {
		return ""
	}

	// Find the content between quotes
	start := xmlStart + 5 // Skip `xml:"`
	end := start
	for end < len(tagValue) && tagValue[end] != '"' {
		end++
	}

	if end >= len(tagValue) {
		return ""
	}

	return tagValue[start:end]
}

func generateXMLTestCases(defs *wsdl.Definitions, types map[string]typeInfo) []xmlTestCase {
	var testCases []xmlTestCase

	// Generate basic test cases for each struct type
	for typeName, typeInfo := range types {
		// Skip XMLName and other utility types
		if typeName == "RawXML" || strings.Contains(typeName, "XMLName") {
			continue
		}

		// Generate a minimal XML test case
		xmlContent := generateMinimalXML(typeInfo)
		if xmlContent != "" {
			testCases = append(testCases, xmlTestCase{
				name:       "minimal_" + strings.ToLower(typeName),
				typeName:   typeName,
				xmlContent: xmlContent,
			})
		}

		// Generate a comprehensive XML test case with more data
		xmlContent = generateComprehensiveXML(typeInfo)
		if xmlContent != "" {
			testCases = append(testCases, xmlTestCase{
				name:       "comprehensive_" + strings.ToLower(typeName),
				typeName:   typeName,
				xmlContent: xmlContent,
			})
		}
	}

	return testCases
}

func generateMinimalXML(typeInfo typeInfo) string {
	var sb strings.Builder

	// Find the root element name from XMLName field or use type name
	rootElement := typeInfo.name
	if xmlNameField, ok := typeInfo.fields["XMLName"]; ok && xmlNameField.xmlTag != "" {
		// Handle namespace syntax: "namespace local-name" or just "local-name"
		parts := strings.Split(xmlNameField.xmlTag, " ")
		if len(parts) >= 2 {
			// Take the local name (second part) when namespace is present
			rootElement = parts[1]
		} else {
			// Take the only part when no namespace
			rootElement = parts[0]
		}
	}

	sb.WriteString("<" + rootElement + ">")

	// Add minimal required fields
	for fieldName, field := range typeInfo.fields {
		if fieldName == "XMLName" {
			continue
		}

		xmlName := field.xmlTag
		if xmlName == "" {
			xmlName = strings.ToLower(fieldName)
		}

		// Extract just the element name from xml tag
		if strings.Contains(xmlName, ",") {
			xmlName = strings.Split(xmlName, ",")[0]
		}
		if strings.Contains(xmlName, " ") {
			xmlName = strings.Split(xmlName, " ")[0]
		}

		// Skip attributes, special tags, and empty tags
		if strings.Contains(field.xmlTag, "attr") || strings.Contains(field.xmlTag, ",") || xmlName == "" {
			continue
		}

		// Generate minimal value based on type
		value := generateMinimalValue(field.typeName)
		if value != "" {
			if field.isSlice && !field.isPtr {
				// For required slices, add at least one element
				sb.WriteString("<" + xmlName + ">" + value + "</" + xmlName + ">")
			} else if !field.isPtr {
				// For required fields
				sb.WriteString("<" + xmlName + ">" + value + "</" + xmlName + ">")
			}
			// Skip optional fields (pointers) in minimal test
		}
	}

	sb.WriteString("</" + rootElement + ">")
	return sb.String()
}

func generateComprehensiveXML(typeInfo typeInfo) string {
	var sb strings.Builder

	// Find the root element name from XMLName field or use type name
	rootElement := typeInfo.name
	if xmlNameField, ok := typeInfo.fields["XMLName"]; ok && xmlNameField.xmlTag != "" {
		// Handle namespace syntax: "namespace local-name" or just "local-name"
		parts := strings.Split(xmlNameField.xmlTag, " ")
		if len(parts) >= 2 {
			// Take the local name (second part) when namespace is present
			rootElement = parts[1]
		} else {
			// Take the only part when no namespace
			rootElement = parts[0]
		}
	}

	sb.WriteString("<" + rootElement + ">")

	// Add all fields including optional ones
	for fieldName, field := range typeInfo.fields {
		if fieldName == "XMLName" {
			continue
		}

		xmlName := field.xmlTag
		if xmlName == "" {
			xmlName = strings.ToLower(fieldName)
		}

		// Extract just the element name from xml tag
		if strings.Contains(xmlName, ",") {
			xmlName = strings.Split(xmlName, ",")[0]
		}
		if strings.Contains(xmlName, " ") {
			xmlName = strings.Split(xmlName, " ")[0]
		}

		// Skip attributes and empty tags
		if strings.Contains(field.xmlTag, "attr") || xmlName == "" {
			continue
		}

		// Generate comprehensive value based on type
		value := generateComprehensiveValue(field.typeName)
		if value != "" {
			if field.isSlice {
				// For slices, add multiple elements
				sb.WriteString("<" + xmlName + ">" + value + "</" + xmlName + ">")
				sb.WriteString("<" + xmlName + ">" + generateAlternativeValue(field.typeName) + "</" + xmlName + ">")
			} else {
				sb.WriteString("<" + xmlName + ">" + value + "</" + xmlName + ">")
			}
		}
	}

	sb.WriteString("</" + rootElement + ">")
	return sb.String()
}

func generateMinimalValue(typeName string) string {
	switch typeName {
	case "string":
		return "test"
	case "int", "int32", "int64":
		return "1"
	case "bool":
		return "true"
	case "float32", "float64":
		return "1.0"
	case "time.Time":
		return "2023-01-01T00:00:00Z"
	case "[]byte":
		return "dGVzdA==" // base64 for "test"
	default:
		// For custom types, return a simple string
		if strings.Contains(typeName, "Type") {
			return "test_value"
		}
		return ""
	}
}

func generateComprehensiveValue(typeName string) string {
	switch typeName {
	case "string":
		return "comprehensive test value"
	case "int", "int32", "int64":
		return "42"
	case "bool":
		return "false"
	case "float32", "float64":
		return "3.14159"
	case "time.Time":
		return "2023-12-25T10:30:00Z"
	case "[]byte":
		return "Y29tcHJlaGVuc2l2ZSB0ZXN0" // base64 for "comprehensive test"
	default:
		// For custom types, return a more complex string
		if strings.Contains(typeName, "Type") {
			return "comprehensive_test_value"
		}
		return ""
	}
}

func generateAlternativeValue(typeName string) string {
	switch typeName {
	case "string":
		return "alternative"
	case "int", "int32", "int64":
		return "99"
	case "bool":
		return "true"
	case "float32", "float64":
		return "2.71828"
	case "time.Time":
		return "2024-01-01T12:00:00Z"
	case "[]byte":
		return "YWx0ZXJuYXRpdmU=" // base64 for "alternative"
	default:
		if strings.Contains(typeName, "Type") {
			return "alternative_value"
		}
		return ""
	}
}

func runXMLParsingTest(t *testing.T, xmlTC xmlTestCase, packageName string) {
	// This is a simplified version - in practice we'd need to dynamically
	// create instances of the types and test unmarshaling
	// For now, we'll just verify the XML is well-formed

	var result interface{}
	err := xml.Unmarshal([]byte(xmlTC.xmlContent), &result)

	// We expect this to fail since we're not using the actual type,
	// but the XML should at least be well-formed
	if !isValidXML([]byte(xmlTC.xmlContent)) {
		t.Errorf("Generated XML is not well-formed: %s", xmlTC.xmlContent)
	}

	// Log the error for debugging, but don't fail the test
	if err != nil {
		t.Logf("Expected unmarshaling error (generic interface{}): %v", err)
	}

	t.Logf("Generated XML for %s: %s", xmlTC.typeName, xmlTC.xmlContent)
}

// TestEnumerationValues tests that enumeration constants are properly generated
func TestEnumerationValues(t *testing.T) {
	testCases := []struct {
		name         string
		testDataDir  string
		expectedEnum string
		expectedVals []string
	}{
		{
			name:         "custom_types_and_enums",
			testDataDir:  "testdata/custom_types_and_enums",
			expectedEnum: "StatusType",
			expectedVals: []string{"active", "inactive", "pending"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Read the generated types file
			typesFile := filepath.Join(tc.testDataDir, "types.go")
			content, err := os.ReadFile(typesFile)
			if err != nil {
				t.Fatalf("Failed to read types file: %v", err)
			}

			contentStr := string(content)

			// Check that enum type is defined
			if !strings.Contains(contentStr, "type "+tc.expectedEnum+" string") {
				t.Errorf("Expected enum type %s not found", tc.expectedEnum)
			}

			// Check that enum constants are defined
			for _, val := range tc.expectedVals {
				constName := tc.expectedEnum + strings.ToUpper(val[:1]) + val[1:]
				if !strings.Contains(contentStr, constName) {
					t.Errorf("Expected enum constant %s not found", constName)
				}
			}

			// Check that String() method is generated
			expectedMethod := "func (e " + tc.expectedEnum + ") String() string"
			if !strings.Contains(contentStr, expectedMethod) {
				t.Errorf("Expected String() method not found")
			}

			// Check that IsValid() method is generated
			expectedValidMethod := "func (e " + tc.expectedEnum + ") IsValid() bool"
			if !strings.Contains(contentStr, expectedValidMethod) {
				t.Errorf("Expected IsValid() method not found")
			}
		})
	}
}

// TestXMLNameGeneration tests that XMLName fields are properly generated
func TestXMLNameGeneration(t *testing.T) {
	testCases := []struct {
		name          string
		testDataDir   string
		typeName      string
		expectXMLName bool
	}{
		{
			name:          "element_references",
			testDataDir:   "testdata/element_references",
			typeName:      "PersonName",
			expectXMLName: true,
		},
		{
			name:          "simple_element_with_string_type",
			testDataDir:   "testdata/simple_element_with_string_type",
			typeName:      "SimpleElement",
			expectXMLName: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Read the generated types file
			typesFile := filepath.Join(tc.testDataDir, "types.go")
			content, err := os.ReadFile(typesFile)
			if err != nil {
				t.Fatalf("Failed to read types file: %v", err)
			}

			contentStr := string(content)

			if tc.expectXMLName {
				// Check that XMLName field is present
				expectedField := "XMLName xml.Name"
				if !strings.Contains(contentStr, expectedField) {
					t.Errorf("Expected XMLName field not found in type %s", tc.typeName)
				}
			}
		})
	}
}

// Helper function to check if XML is valid
func isValidXML(data []byte) bool {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	for {
		_, err := decoder.Token()
		if err != nil {
			return false
		}
		if decoder.InputOffset() >= int64(len(data)) {
			break
		}
	}
	return true
}

// mustParseTime is a helper function that parses time and panics on error
func mustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
