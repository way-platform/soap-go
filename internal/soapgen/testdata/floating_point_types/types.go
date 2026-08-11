package floating_point_types

import (
	"encoding/xml"
)

// Complex types

// FloatingPointAttributes represents the FloatingPointAttributes complex type
type FloatingPointAttributes struct {
	FloatAttr           *float64 `xml:"floatAttr,attr,omitempty"`
	DoubleAttr          *float64 `xml:"doubleAttr,attr,omitempty"`
	DecimalAttr         *string  `xml:"decimalAttr,attr,omitempty"`
	OptionalFloatAttr   *float64 `xml:"optionalFloatAttr,attr,omitempty"`
	OptionalDoubleAttr  *float64 `xml:"optionalDoubleAttr,attr,omitempty"`
	OptionalDecimalAttr *string  `xml:"optionalDecimalAttr,attr,omitempty"`
}

// FloatingPointContainer represents the FloatingPointContainer complex type
type FloatingPointContainer struct {
	FloatValue           float64  `xml:"FloatValue"`
	DoubleValue          float64  `xml:"DoubleValue"`
	DecimalValue         string   `xml:"DecimalValue"`
	OptionalFloatValue   *float64 `xml:"OptionalFloatValue,omitempty"`
	OptionalDoubleValue  *float64 `xml:"OptionalDoubleValue,omitempty"`
	OptionalDecimalValue *string  `xml:"OptionalDecimalValue,omitempty"`
}

// FloatElementWrapper represents the FloatElement element
type FloatElementWrapper struct {
	XMLName xml.Name `xml:"FloatElement"`
	Value   float64  `xml:",chardata"`
}

// DoubleElementWrapper represents the DoubleElement element
type DoubleElementWrapper struct {
	XMLName xml.Name `xml:"DoubleElement"`
	Value   float64  `xml:",chardata"`
}

// DecimalElementWrapper represents the DecimalElement element
type DecimalElementWrapper struct {
	XMLName xml.Name `xml:"DecimalElement"`
	Value   string   `xml:",chardata"`
}

// FloatingPointContainerWrapper represents the FloatingPointContainer element
type FloatingPointContainerWrapper struct {
	XMLName              xml.Name `xml:"FloatingPointContainer"`
	FloatValue           float64  `xml:"FloatValue"`
	DoubleValue          float64  `xml:"DoubleValue"`
	DecimalValue         string   `xml:"DecimalValue"`
	OptionalFloatValue   *float64 `xml:"OptionalFloatValue,omitempty"`
	OptionalDoubleValue  *float64 `xml:"OptionalDoubleValue,omitempty"`
	OptionalDecimalValue *string  `xml:"OptionalDecimalValue,omitempty"`
}

// FloatingPointAttributesWrapper represents the FloatingPointAttributes element
type FloatingPointAttributesWrapper struct {
	XMLName             xml.Name `xml:"FloatingPointAttributes"`
	FloatAttr           *float64 `xml:"floatAttr,attr,omitempty"`
	DoubleAttr          *float64 `xml:"doubleAttr,attr,omitempty"`
	DecimalAttr         *string  `xml:"decimalAttr,attr,omitempty"`
	OptionalFloatAttr   *float64 `xml:"optionalFloatAttr,attr,omitempty"`
	OptionalDoubleAttr  *float64 `xml:"optionalDoubleAttr,attr,omitempty"`
	OptionalDecimalAttr *string  `xml:"optionalDecimalAttr,attr,omitempty"`
}
