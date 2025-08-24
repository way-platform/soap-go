package numberconversion

// NumberToWords represents the NumberToWords element
type NumberToWords struct {
	UbiNum uint64 `xml:"ubiNum"`
}

// NumberToWordsResponse represents the NumberToWordsResponse element
type NumberToWordsResponse struct {
	NumberToWordsResult string `xml:"NumberToWordsResult"`
}

// NumberToDollars represents the NumberToDollars element
type NumberToDollars struct {
	DNum float64 `xml:"dNum"`
}

// NumberToDollarsResponse represents the NumberToDollarsResponse element
type NumberToDollarsResponse struct {
	NumberToDollarsResult string `xml:"NumberToDollarsResult"`
}
