package numberconversion

import (
	"encoding/xml"
)

// NumberToWords represents the NumberToWords element
type NumberToWords struct {
	XMLName xml.Name `xml:"NumberToWords"`
	UbiNum  uint64   `xml:"ubiNum"`
}

// NumberToWordsResponse represents the NumberToWordsResponse element
type NumberToWordsResponse struct {
	XMLName             xml.Name `xml:"NumberToWordsResponse"`
	NumberToWordsResult string   `xml:"NumberToWordsResult"`
}

// NumberToDollars represents the NumberToDollars element
type NumberToDollars struct {
	XMLName xml.Name `xml:"NumberToDollars"`
	DNum    float64  `xml:"dNum"`
}

// NumberToDollarsResponse represents the NumberToDollarsResponse element
type NumberToDollarsResponse struct {
	XMLName               xml.Name `xml:"NumberToDollarsResponse"`
	NumberToDollarsResult string   `xml:"NumberToDollarsResult"`
}
