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

// Request represents the Request element
type Request struct {
	XMLName xml.Name    `xml:"Request"`
	Value   RequestType `xml:",chardata"`
}

// RequestElement represents the request element
type RequestElement struct {
	XMLName xml.Name    `xml:"request"`
	Value   RequestType `xml:",chardata"`
}

// REQUEST represents the REQUEST element
type REQUEST struct {
	XMLName xml.Name    `xml:"REQUEST"`
	Value   RequestType `xml:",chardata"`
}

// RequestElement represents the requestElement element
type RequestElement struct {
	XMLName xml.Name    `xml:"requestElement"`
	Value   RequestType `xml:",chardata"`
}

// RequestWrapper represents the requestWrapper element
type RequestWrapper struct {
	XMLName xml.Name    `xml:"requestWrapper"`
	Value   RequestType `xml:",chardata"`
}

// RequestOperation represents the requestOperation element
type RequestOperation struct {
	XMLName xml.Name    `xml:"requestOperation"`
	Value   RequestType `xml:",chardata"`
}

// Request1 represents the Request1 element
type Request1 struct {
	XMLName xml.Name    `xml:"Request1"`
	Value   RequestType `xml:",chardata"`
}

// Request2 represents the Request2 element
type Request2 struct {
	XMLName xml.Name    `xml:"Request2"`
	Value   RequestType `xml:",chardata"`
}

// Data represents the Data element
type Data struct {
	XMLName xml.Name `xml:"Data"`
	Value   DataType `xml:",chardata"`
}

// DataElement represents the data element
type DataElement struct {
	XMLName xml.Name `xml:"data"`
	Value   DataType `xml:",chardata"`
}

// DATA represents the DATA element
type DATA struct {
	XMLName xml.Name `xml:"DATA"`
	Value   DataType `xml:",chardata"`
}

// DataElement represents the dataElement element
type DataElement struct {
	XMLName xml.Name `xml:"dataElement"`
	Value   DataType `xml:",chardata"`
}

// DataWrapper represents the dataWrapper element
type DataWrapper struct {
	XMLName xml.Name `xml:"dataWrapper"`
	Value   DataType `xml:",chardata"`
}

// DataOperation represents the dataOperation element
type DataOperation struct {
	XMLName xml.Name `xml:"dataOperation"`
	Value   DataType `xml:",chardata"`
}

// Data1 represents the Data1 element
type Data1 struct {
	XMLName xml.Name `xml:"Data1"`
	Value   DataType `xml:",chardata"`
}

// Data2 represents the Data2 element
type Data2 struct {
	XMLName xml.Name `xml:"Data2"`
	Value   DataType `xml:",chardata"`
}

// Data3 represents the Data3 element
type Data3 struct {
	XMLName xml.Name `xml:"Data3"`
	Value   DataType `xml:",chardata"`
}

// ExtremeCaseElement represents the ExtremeCaseElement element
type ExtremeCaseElement struct {
	XMLName xml.Name        `xml:"ExtremeCaseElement"`
	Value   ExtremeCaseType `xml:",chardata"`
}
