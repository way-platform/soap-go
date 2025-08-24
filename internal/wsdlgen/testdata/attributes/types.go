package attributes

// WithAttributes represents the WithAttributes element
type WithAttributes struct {
	Content string `xml:"content"`
	Id      string `xml:"id,attr"`
	Version *int32 `xml:"version,attr"`
}
