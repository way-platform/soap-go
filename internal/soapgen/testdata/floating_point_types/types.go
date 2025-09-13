package floating_point_types

import (
	"encoding/xml"
)

// Complex types

// FloatingPointAttributes represents the FloatingPointAttributes complex type
type FloatingPointAttributes struct {
	FloatAttr           *float64 `xml:"floatAttr,attr,omitempty"`
	DoubleAttr          *float64 `xml:"doubleAttr,attr,omitempty"`
	DecimalAttr         *float64 `xml:"decimalAttr,attr,omitempty"`
	OptionalFloatAttr   *float64 `xml:"optionalFloatAttr,attr,omitempty"`
	OptionalDoubleAttr  *float64 `xml:"optionalDoubleAttr,attr,omitempty"`
	OptionalDecimalAttr *float64 `xml:"optionalDecimalAttr,attr,omitempty"`
}

// FloatingPointContainer represents the FloatingPointContainer complex type
type FloatingPointContainer struct {
	FloatValue           float64  `xml:"FloatValue"`
	DoubleValue          float64  `xml:"DoubleValue"`
	DecimalValue         float64  `xml:"DecimalValue"`
	OptionalFloatValue   *float64 `xml:"OptionalFloatValue,omitempty"`
	OptionalDoubleValue  *float64 `xml:"OptionalDoubleValue,omitempty"`
	OptionalDecimalValue *float64 `xml:"OptionalDecimalValue,omitempty"`
}

// FloatElement represents the FloatElement element
type FloatElement struct {
	XMLName xml.Name `xml:"FloatElement"`
	Value   float64  `xml:",chardata"`
}

// DoubleElement represents the DoubleElement element
type DoubleElement struct {
	XMLName xml.Name `xml:"DoubleElement"`
	Value   float64  `xml:",chardata"`
}

// DecimalElement represents the DecimalElement element
type DecimalElement struct {
	XMLName xml.Name `xml:"DecimalElement"`
	Value   float64  `xml:",chardata"`
}

// FloatingPointContainer represents the FloatingPointContainer element
type FloatingPointContainer struct {
	XMLName              xml.Name `xml:"FloatingPointContainer"`
	FloatValue           float64  `xml:"FloatValue"`
	DoubleValue          float64  `xml:"DoubleValue"`
	DecimalValue         float64  `xml:"DecimalValue"`
	OptionalFloatValue   *float64 `xml:"OptionalFloatValue,omitempty"`
	OptionalDoubleValue  *float64 `xml:"OptionalDoubleValue,omitempty"`
	OptionalDecimalValue *float64 `xml:"OptionalDecimalValue,omitempty"`
}

// FloatingPointAttributes represents the FloatingPointAttributes element
type FloatingPointAttributes struct {
	XMLName             xml.Name `xml:"FloatingPointAttributes"`
	FloatAttr           *float64 `xml:"floatAttr,attr,omitempty"`
	DoubleAttr          *float64 `xml:"doubleAttr,attr,omitempty"`
	DecimalAttr         *float64 `xml:"decimalAttr,attr,omitempty"`
	OptionalFloatAttr   *float64 `xml:"optionalFloatAttr,attr,omitempty"`
	OptionalDoubleAttr  *float64 `xml:"optionalDoubleAttr,attr,omitempty"`
	OptionalDecimalAttr *float64 `xml:"optionalDecimalAttr,attr,omitempty"`
}
