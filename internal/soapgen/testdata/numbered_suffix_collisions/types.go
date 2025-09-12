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

// RequestElement represents the requestElement element
type RequestElement struct {
	XMLName xml.Name `xml:"requestElement"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// RequestOperation represents the requestOperation element
type RequestOperation struct {
	XMLName xml.Name `xml:"requestOperation"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// Request1 represents the Request1 element
type Request1 struct {
	XMLName xml.Name `xml:"Request1"`
	Id      string   `xml:"id"`
	Type    string   `xml:"type"`
}

// Request2 represents the Request2 element
type Request2 struct {
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

// DataElement represents the dataElement element
type DataElement struct {
	XMLName xml.Name `xml:"dataElement"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// DataOperation represents the dataOperation element
type DataOperation struct {
	XMLName xml.Name `xml:"dataOperation"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data1 represents the Data1 element
type Data1 struct {
	XMLName xml.Name `xml:"Data1"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data2 represents the Data2 element
type Data2 struct {
	XMLName xml.Name `xml:"Data2"`
	Content string   `xml:"content"`
	Format  string   `xml:"format"`
}

// Data3 represents the Data3 element
type Data3 struct {
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
