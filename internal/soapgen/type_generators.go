package soapgen

import (
	"sort"
	"strings"

	"github.com/way-platform/soap-go/internal/codegen"
	"github.com/way-platform/soap-go/xsd"
)

// generateSimpleTypeConstants generates Go constants for simple types (mainly enumerations)
func generateSimpleTypeConstants(g *codegen.File, ctx *SchemaContext) {
	if len(ctx.simpleTypes) == 0 {
		return
	}

	g.P("// Enumeration types")
	g.P()

	// Sort simple type names for deterministic output
	var names []string
	for name := range ctx.simpleTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	// Generate each simple type
	for _, name := range names {
		simpleType := ctx.simpleTypes[name]
		if simpleType.Restriction != nil && len(simpleType.Restriction.Enumerations) > 0 {
			generateEnumType(g, simpleType)
		}
	}
}

// generateEnumType generates a Go enum type from an XSD simple type with enumerations
func generateEnumType(g *codegen.File, simpleType *xsd.SimpleType) {
	typeName := toGoName(simpleType.Name)

	// Generate the enum type definition
	g.P("// ", typeName, " represents an enumeration type")
	g.P("type ", typeName, " ", g.QualifiedGoIdent(codegen.StringIdent))
	g.P()

	// Generate the constants with typed values
	g.P("// ", typeName, " enumeration values")
	g.P("const (")

	var enumValues []string
	for _, enum := range simpleType.Restriction.Enumerations {
		constName := typeName + toGoName(enum.Value)
		enumValues = append(enumValues, constName)
		g.P("\t", constName, " ", typeName, " = \"", enum.Value, "\"")
	}
	g.P(")")
	g.P()

	// Generate String method
	g.P("// String returns the string representation of ", typeName)
	g.P("func (e ", typeName, ") String() ", g.QualifiedGoIdent(codegen.StringIdent), " {")
	g.P("\treturn ", g.QualifiedGoIdent(codegen.StringIdent), "(e)")
	g.P("}")
	g.P()

	// Generate IsValid method
	g.P("// IsValid returns true if the ", typeName, " value is valid")
	g.P("func (e ", typeName, ") IsValid() ", g.QualifiedGoIdent(codegen.BoolIdent), " {")
	g.P("\tswitch e {")
	g.P("\tcase ", strings.Join(enumValues, ", "), ":")
	g.P("\t\treturn true")
	g.P("\tdefault:")
	g.P("\t\treturn false")
	g.P("\t}")
	g.P("}")
	g.P()
}

// generateInlineEnumTypes generates Go enum types for inline enumerations
func generateInlineEnumTypes(g *codegen.File, ctx *SchemaContext) {
	if len(ctx.inlineEnums) == 0 {
		return
	}

	g.P("// Inline enumeration types")
	g.P()

	// Collect unique enum types to generate (avoiding duplicates from deduplication)
	uniqueEnums := make(map[string]InlineEnumInfo)
	for _, enumInfo := range ctx.inlineEnums {
		if !enumInfo.Generated {
			uniqueEnums[enumInfo.TypeName] = enumInfo
		}
	}

	// Sort enum type names for deterministic output
	var typeNames []string
	for typeName := range uniqueEnums {
		typeNames = append(typeNames, typeName)
	}
	sort.Strings(typeNames)

	// Generate each unique inline enum type
	for _, typeName := range typeNames {
		enumInfo := uniqueEnums[typeName]
		generateInlineEnumType(g, &enumInfo)

		// Mark all enums with this type name as generated
		for key, info := range ctx.inlineEnums {
			if info.TypeName == typeName {
				info.Generated = true
				ctx.inlineEnums[key] = info
			}
		}
	}
}

// generateInlineEnumType generates a Go enum type from an inline enum info
func generateInlineEnumType(g *codegen.File, enumInfo *InlineEnumInfo) {
	typeName := enumInfo.TypeName
	simpleType := enumInfo.SimpleType

	// Generate the enum type definition
	g.P("// ", typeName, " represents an inline enumeration type")
	g.P("type ", typeName, " ", g.QualifiedGoIdent(codegen.StringIdent))
	g.P()

	// Generate the constants with typed values
	g.P("// ", typeName, " enumeration values")
	g.P("const (")

	var enumValues []string
	for _, enum := range simpleType.Restriction.Enumerations {
		constName := typeName + toGoName(enum.Value)
		enumValues = append(enumValues, constName)
		g.P("\t", constName, " ", typeName, " = \"", enum.Value, "\"")
	}
	g.P(")")
	g.P()

	// Generate String method
	g.P("// String returns the string representation of ", typeName)
	g.P("func (e ", typeName, ") String() ", g.QualifiedGoIdent(codegen.StringIdent), " {")
	g.P("\treturn ", g.QualifiedGoIdent(codegen.StringIdent), "(e)")
	g.P("}")
	g.P()

	// Generate IsValid method
	g.P("// IsValid returns true if the ", typeName, " value is valid")
	g.P("func (e ", typeName, ") IsValid() ", g.QualifiedGoIdent(codegen.BoolIdent), " {")
	g.P("\tswitch e {")
	g.P("\tcase ", strings.Join(enumValues, ", "), ":")
	g.P("\t\treturn true")
	g.P("\tdefault:")
	g.P("\t\treturn false")
	g.P("\t}")
	g.P("}")
	g.P()
}

// generateComplexTypes generates Go structs for named complex types
func generateComplexTypes(g *codegen.File, ctx *SchemaContext) {
	if len(ctx.complexTypes) == 0 {
		return
	}

	g.P("// Complex types")
	g.P()

	// Sort complex type names for deterministic output
	var names []string
	for name := range ctx.complexTypes {
		names = append(names, name)
	}
	sort.Strings(names)

	// Generate each complex type
	for _, name := range names {
		complexType := ctx.complexTypes[name]
		generateStructFromComplexType(g, complexType, ctx)
	}
}
