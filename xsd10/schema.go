package xsd10

import (
	"encoding/xml"
	"io"
)

// Parse reads an XSD schema from an io.Reader and unmarshals it.
func Parse(r io.Reader) (*Schema, error) {
	var schema Schema
	if err := xml.NewDecoder(r).Decode(&schema); err != nil {
		return nil, err
	}
	return &schema, nil
}

// Schema represents an <xsd:schema> element.
type Schema struct {
	XMLName              xml.Name `xml:"schema"`
	TargetNamespace      string   `xml:"targetNamespace,attr"`
	ElementFormDefault   string   `xml:"elementFormDefault,attr"`
	AttributeFormDefault string   `xml:"attributeFormDefault,attr"`

	Imports         []Import         `xml:"import"`
	Includes        []Include        `xml:"include"`
	Elements        []Element        `xml:"element"`
	ComplexTypes    []ComplexType    `xml:"complexType"`
	SimpleTypes     []SimpleType     `xml:"simpleType"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	Groups          []Group          `xml:"group"`
	Annotations     []Annotation     `xml:"annotation"`
}

// Import corresponds to <xsd:import>.
type Import struct {
	Namespace      string `xml:"namespace,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
}

// Include corresponds to <xsd:include>.
type Include struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
}

// Element corresponds to <xsd:element>.
type Element struct {
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	Ref         string       `xml:"ref,attr"`
	MinOccurs   string       `xml:"minOccurs,attr"`
	MaxOccurs   string       `xml:"maxOccurs,attr"`
	Nillable    bool         `xml:"nillable,attr"`
	Default     string       `xml:"default,attr"`
	Fixed       string       `xml:"fixed,attr"`
	ComplexType *ComplexType `xml:"complexType"`
	SimpleType  *SimpleType  `xml:"simpleType"`
	Annotation  *Annotation  `xml:"annotation"`
}

// ComplexType corresponds to <xsd:complexType>.
type ComplexType struct {
	Name            string           `xml:"name,attr"`
	Mixed           bool             `xml:"mixed,attr"`
	Final           string           `xml:"final,attr"`
	Abstract        bool             `xml:"abstract,attr"`
	Sequence        *Sequence        `xml:"sequence"`
	Choice          *Choice          `xml:"choice"`
	All             *All             `xml:"all"`
	Group           *Group           `xml:"group"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
	SimpleContent   *SimpleContent   `xml:"simpleContent"`
	ComplexContent  *ComplexContent  `xml:"complexContent"`
	Annotation      *Annotation      `xml:"annotation"`
}

// Sequence corresponds to <xsd:sequence>.
type Sequence struct {
	MinOccurs  string      `xml:"minOccurs,attr"`
	MaxOccurs  string      `xml:"maxOccurs,attr"`
	Elements   []Element   `xml:"element"`
	Groups     []Group     `xml:"group"`
	Choices    []Choice    `xml:"choice"`
	Sequences  []Sequence  `xml:"sequence"`
	Any        []Any       `xml:"any"`
	Annotation *Annotation `xml:"annotation"`
}

// Choice corresponds to <xsd:choice>.
type Choice struct {
	MinOccurs  string      `xml:"minOccurs,attr"`
	MaxOccurs  string      `xml:"maxOccurs,attr"`
	Elements   []Element   `xml:"element"`
	Groups     []Group     `xml:"group"`
	Choices    []Choice    `xml:"choice"`
	Sequences  []Sequence  `xml:"sequence"`
	Any        []Any       `xml:"any"`
	Annotation *Annotation `xml:"annotation"`
}

// All corresponds to <xsd:all>.
type All struct {
	MinOccurs  string      `xml:"minOccurs,attr"`
	MaxOccurs  string      `xml:"maxOccurs,attr"`
	Elements   []Element   `xml:"element"`
	Annotation *Annotation `xml:"annotation"`
}

// Group corresponds to <xsd:group>.
type Group struct {
	Name       string      `xml:"name,attr"`
	Ref        string      `xml:"ref,attr"`
	MinOccurs  string      `xml:"minOccurs,attr"`
	MaxOccurs  string      `xml:"maxOccurs,attr"`
	Sequence   *Sequence   `xml:"sequence"`
	Choice     *Choice     `xml:"choice"`
	All        *All        `xml:"all"`
	Annotation *Annotation `xml:"annotation"`
}

// Attribute corresponds to <xsd:attribute>.
type Attribute struct {
	Name       string      `xml:"name,attr"`
	Type       string      `xml:"type,attr"`
	Use        string      `xml:"use,attr"` // optional, prohibited, required
	Ref        string      `xml:"ref,attr"`
	Default    string      `xml:"default,attr"`
	Fixed      string      `xml:"fixed,attr"`
	Form       string      `xml:"form,attr"`
	SimpleType *SimpleType `xml:"simpleType"`
	Annotation *Annotation `xml:"annotation"`
}

// AttributeGroup corresponds to <xsd:attributeGroup>.
type AttributeGroup struct {
	Name            string           `xml:"name,attr"`
	Ref             string           `xml:"ref,attr"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
	Annotation      *Annotation      `xml:"annotation"`
}

// SimpleType corresponds to <xsd:simpleType>.
type SimpleType struct {
	Name        string       `xml:"name,attr"`
	Final       string       `xml:"final,attr"`
	Restriction *Restriction `xml:"restriction"`
	List        *List        `xml:"list"`
	Union       *Union       `xml:"union"`
	Annotation  *Annotation  `xml:"annotation"`
}

// Restriction corresponds to <xsd:restriction>.
type Restriction struct {
	Base            string           `xml:"base,attr"`
	Enumerations    []Enumeration    `xml:"enumeration"`
	Patterns        []Pattern        `xml:"pattern"`
	MinInclusive    *MinInclusive    `xml:"minInclusive"`
	MaxInclusive    *MaxInclusive    `xml:"maxInclusive"`
	MinExclusive    *MinExclusive    `xml:"minExclusive"`
	MaxExclusive    *MaxExclusive    `xml:"maxExclusive"`
	Length          *Length          `xml:"length"`
	MinLength       *MinLength       `xml:"minLength"`
	MaxLength       *MaxLength       `xml:"maxLength"`
	WhiteSpace      *WhiteSpace      `xml:"whiteSpace"`
	TotalDigits     *TotalDigits     `xml:"totalDigits"`
	FractionDigits  *FractionDigits  `xml:"fractionDigits"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
	Annotation      *Annotation      `xml:"annotation"`
}

// List corresponds to <xsd:list>.
type List struct {
	ItemType   string      `xml:"itemType,attr"`
	SimpleType *SimpleType `xml:"simpleType"`
	Annotation *Annotation `xml:"annotation"`
}

// Union corresponds to <xsd:union>.
type Union struct {
	MemberTypes string       `xml:"memberTypes,attr"`
	SimpleTypes []SimpleType `xml:"simpleType"`
	Annotation  *Annotation  `xml:"annotation"`
}

// SimpleContent corresponds to <xsd:simpleContent>.
type SimpleContent struct {
	Extension   *Extension   `xml:"extension"`
	Restriction *Restriction `xml:"restriction"`
	Annotation  *Annotation  `xml:"annotation"`
}

// ComplexContent corresponds to <xsd:complexContent>.
type ComplexContent struct {
	Mixed       bool         `xml:"mixed,attr"`
	Extension   *Extension   `xml:"extension"`
	Restriction *Restriction `xml:"restriction"`
	Annotation  *Annotation  `xml:"annotation"`
}

// Extension corresponds to <xsd:extension>.
type Extension struct {
	Base            string           `xml:"base,attr"`
	Sequence        *Sequence        `xml:"sequence"`
	Choice          *Choice          `xml:"choice"`
	All             *All             `xml:"all"`
	Group           *Group           `xml:"group"`
	Attributes      []Attribute      `xml:"attribute"`
	AttributeGroups []AttributeGroup `xml:"attributeGroup"`
	AnyAttribute    *AnyAttribute    `xml:"anyAttribute"`
	Annotation      *Annotation      `xml:"annotation"`
}

// Any corresponds to <xsd:any>.
type Any struct {
	Namespace       string      `xml:"namespace,attr"`
	ProcessContents string      `xml:"processContents,attr"`
	MinOccurs       string      `xml:"minOccurs,attr"`
	MaxOccurs       string      `xml:"maxOccurs,attr"`
	Annotation      *Annotation `xml:"annotation"`
}

// AnyAttribute corresponds to <xsd:anyAttribute>.
type AnyAttribute struct {
	Namespace       string      `xml:"namespace,attr"`
	ProcessContents string      `xml:"processContents,attr"`
	Annotation      *Annotation `xml:"annotation"`
}

// Annotation corresponds to <xsd:annotation>.
type Annotation struct {
	Documentation []Documentation `xml:"documentation"`
	AppInfo       []AppInfo       `xml:"appinfo"`
}

// Documentation corresponds to <xsd:documentation>.
type Documentation struct {
	Source  string `xml:"source,attr"`
	Lang    string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Content string `xml:",chardata"`
}

// AppInfo corresponds to <xsd:appinfo>.
type AppInfo struct {
	Source  string `xml:"source,attr"`
	Content string `xml:",chardata"`
}

// Facet types for restrictions

// Enumeration corresponds to <xsd:enumeration>.
type Enumeration struct {
	Value      string      `xml:"value,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// Pattern corresponds to <xsd:pattern>.
type Pattern struct {
	Value      string      `xml:"value,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MinInclusive corresponds to <xsd:minInclusive>.
type MinInclusive struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MaxInclusive corresponds to <xsd:maxInclusive>.
type MaxInclusive struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MinExclusive corresponds to <xsd:minExclusive>.
type MinExclusive struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MaxExclusive corresponds to <xsd:maxExclusive>.
type MaxExclusive struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// Length corresponds to <xsd:length>.
type Length struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MinLength corresponds to <xsd:minLength>.
type MinLength struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// MaxLength corresponds to <xsd:maxLength>.
type MaxLength struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// WhiteSpace corresponds to <xsd:whiteSpace>.
type WhiteSpace struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// TotalDigits corresponds to <xsd:totalDigits>.
type TotalDigits struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}

// FractionDigits corresponds to <xsd:fractionDigits>.
type FractionDigits struct {
	Value      string      `xml:"value,attr"`
	Fixed      bool        `xml:"fixed,attr"`
	Annotation *Annotation `xml:"annotation"`
}
