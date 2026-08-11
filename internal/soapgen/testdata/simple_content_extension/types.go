package simple_content_extension

import (
	"encoding/xml"
	"time"
)

// StateElementWrapper represents the StateElement element
type StateElementWrapper struct {
	XMLName   xml.Name   `xml:"StateElement"`
	Value     string     `xml:",chardata"`
	Name      string     `xml:"name,attr"`
	Timestamp *time.Time `xml:"timestamp,attr,omitempty"`
}

// ValueElementWrapper represents the ValueElement element
type ValueElementWrapper struct {
	XMLName   xml.Name `xml:"ValueElement"`
	Value     float64  `xml:",chardata"`
	Unit      *string  `xml:"unit,attr,omitempty"`
	Precision *int32   `xml:"precision,attr,omitempty"`
}

// StatesContainerWrapper represents the StatesContainer element
type StatesContainerWrapper struct {
	XMLName      xml.Name              `xml:"http://tempuri.org/ StatesContainer"`
	StateElement []StateElementWrapper `xml:"StateElement,omitempty"`
	ValueElement []ValueElementWrapper `xml:"ValueElement,omitempty"`
}
