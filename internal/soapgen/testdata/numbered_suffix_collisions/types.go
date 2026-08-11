package numbered_suffix_collisions

import (
	"encoding/xml"
)

// Complex types

// DataType represents the DataType complex type
type DataType struct {
	Content string `xml:"content"`
	Format  string `xml:"format"`
}

// ExtremeCaseType represents the ExtremeCaseType complex type
type ExtremeCaseType struct {
	Value           string  `xml:"Value"`
	ValueElem       string  `xml:"value"`
	VALUE           string  `xml:"VALUE"`
	ValueElemElem   string  `xml:"valueElem"`
	ValueElemElem1  string  `xml:"ValueElem"`
	ValueElem1      string  `xml:"ValueElem1"`
	Status          string  `xml:"Status"`
	StatusElem      string  `xml:"status"`
	StatusElemElem  string  `xml:"statusElem"`
	StatusElemElem1 string  `xml:"StatusElem"`
	ValueAttr       *string `xml:"value,attr,omitempty"`
	ValueAttrAttr   *string `xml:"valueAttr,attr,omitempty"`
	ValueAttrAttr1  *string `xml:"ValueAttr,attr,omitempty"`
	ValueAttr1      *string `xml:"ValueAttr1,attr,omitempty"`
	StatusAttr      *string `xml:"status,attr,omitempty"`
	StatusAttrAttr  *string `xml:"statusAttr,attr,omitempty"`
	StatusAttrAttr1 *string `xml:"StatusAttr,attr,omitempty"`
}

// RequestType represents the RequestType complex type
type RequestType struct {
	Id   string `xml:"id"`
	Type string `xml:"type"`
}

// RequestWrapper represents the Request element
type RequestWrapper struct {
	XMLName xml.Name `xml:"http://example.com/numbered-suffix-collisions Request"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// REQUESTWrapper represents the REQUEST element
type REQUESTWrapper struct {
	XMLName xml.Name `xml:"http://example.com/numbered-suffix-collisions REQUEST"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// RequestElementWrapper represents the requestElement element
type RequestElementWrapper struct {
	XMLName xml.Name `xml:"requestElement"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// RequestWrapperWrapper represents the requestWrapper element
type RequestWrapperWrapper struct {
	XMLName xml.Name `xml:"requestWrapper"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// RequestOperationWrapper represents the requestOperation element
type RequestOperationWrapper struct {
	XMLName xml.Name `xml:"requestOperation"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// Request1Wrapper represents the Request1 element
type Request1Wrapper struct {
	XMLName xml.Name `xml:"Request1"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// Request2Wrapper represents the Request2 element
type Request2Wrapper struct {
	XMLName xml.Name `xml:"Request2"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// DataWrapper represents the Data element
type DataWrapper struct {
	XMLName xml.Name `xml:"http://example.com/numbered-suffix-collisions Data"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// DATAWrapper represents the DATA element
type DATAWrapper struct {
	XMLName xml.Name `xml:"http://example.com/numbered-suffix-collisions DATA"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// DataElementWrapper represents the dataElement element
type DataElementWrapper struct {
	XMLName xml.Name `xml:"dataElement"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// DataWrapperWrapper represents the dataWrapper element
type DataWrapperWrapper struct {
	XMLName xml.Name `xml:"dataWrapper"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// DataOperationWrapper represents the dataOperation element
type DataOperationWrapper struct {
	XMLName xml.Name `xml:"dataOperation"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data1Wrapper represents the Data1 element
type Data1Wrapper struct {
	XMLName xml.Name `xml:"Data1"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data2Wrapper represents the Data2 element
type Data2Wrapper struct {
	XMLName xml.Name `xml:"Data2"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data3Wrapper represents the Data3 element
type Data3Wrapper struct {
	XMLName xml.Name `xml:"Data3"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// ExtremeCaseElementWrapper represents the ExtremeCaseElement element
type ExtremeCaseElementWrapper struct {
	XMLName         xml.Name `xml:"http://example.com/numbered-suffix-collisions ExtremeCaseElement"`
	Value           string   `xml:"Value"`
	ValueElem       string   `xml:"value"`
	VALUE           string   `xml:"VALUE"`
	ValueElemElem   string   `xml:"valueElem"`
	ValueElemElem1  string   `xml:"ValueElem"`
	ValueElem1      string   `xml:"ValueElem1"`
	Status          string   `xml:"Status"`
	StatusElem      string   `xml:"status"`
	StatusElemElem  string   `xml:"statusElem"`
	StatusElemElem1 string   `xml:"StatusElem"`
	ValueAttr       *string  `xml:"value,attr,omitempty"`
	ValueAttrAttr   *string  `xml:"valueAttr,attr,omitempty"`
	ValueAttrAttr1  *string  `xml:"ValueAttr,attr,omitempty"`
	ValueAttr1      *string  `xml:"ValueAttr1,attr,omitempty"`
	StatusAttr      *string  `xml:"status,attr,omitempty"`
	StatusAttrAttr  *string  `xml:"statusAttr,attr,omitempty"`
	StatusAttrAttr1 *string  `xml:"StatusAttr,attr,omitempty"`
}
