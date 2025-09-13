package soap

import "fmt"

// Error contains HTTP and SOAP fault information.
type Error struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int

	// ResponseBody is the raw HTTP response body.
	ResponseBody []byte

	// Envelope is the SOAP envelope that was received, nil if parsing failed.
	Envelope *Envelope

	// Fault is the SOAP fault, nil if no fault was present.
	Fault *Fault
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Fault != nil {
		return fmt.Sprintf("SOAP fault (HTTP %d): %v", e.StatusCode, e.Fault)
	}
	return fmt.Sprintf("HTTP error %d: %s", e.StatusCode, string(e.ResponseBody))
}
