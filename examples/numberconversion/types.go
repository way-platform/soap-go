package numberconversion

import (
	"encoding/xml"
)

// NumberToWordsWrapper represents the NumberToWords element
type NumberToWordsWrapper struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`
	UbiNum  uint64   `xml:"ubiNum"`
}

// NumberToWordsResponseWrapper represents the NumberToWordsResponse element
type NumberToWordsResponseWrapper struct {
	XMLName             xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWordsResponse"`
	NumberToWordsResult string   `xml:"NumberToWordsResult"`
}

// NumberToDollarsWrapper represents the NumberToDollars element
type NumberToDollarsWrapper struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToDollars"`
	DNum    float64  `xml:"dNum"`
}

// NumberToDollarsResponseWrapper represents the NumberToDollarsResponse element
type NumberToDollarsResponseWrapper struct {
	XMLName               xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToDollarsResponse"`
	NumberToDollarsResult string   `xml:"NumberToDollarsResult"`
}
