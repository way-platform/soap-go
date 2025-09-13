package simple_content_extension

import (
	"encoding/xml"
	"time"
)

// StateElement represents the StateElement element
type StateElement struct {
	XMLName   xml.Name   `xml:"StateElement"`
	Value     string     `xml:",chardata"`
	Name      string     `xml:"name,attr"`
	Timestamp *time.Time `xml:"timestamp,attr,omitempty"`
}

// ValueElement represents the ValueElement element
type ValueElement struct {
	XMLName   xml.Name `xml:"ValueElement"`
	Value     float64  `xml:",chardata"`
	Unit      *string  `xml:"unit,attr,omitempty"`
	Precision *int32   `xml:"precision,attr,omitempty"`
}

// StatesContainerWrapper represents the StatesContainer element
type StatesContainerWrapper struct {
	XMLName      xml.Name       `xml:"http://tempuri.org/ StatesContainer"`
	StateElement []StateElement `xml:"StateElement,omitempty"`
	ValueElement []ValueElement `xml:"ValueElement,omitempty"`
}
