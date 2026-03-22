package soap

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

// DebugTransport is an [http.RoundTripper] that dumps requests and responses
// to stderr. When Enabled is nil or points to false, requests pass through
// to Next unchanged.
//
// Standalone CLIs can wire a --debug flag to Enabled:
//
//	var debug bool
//	t := &soap.DebugTransport{Enabled: &debug, Next: http.DefaultTransport}
//	client, _ := soap.NewClient(soap.WithHTTPClient(&http.Client{Transport: t}))
//	flag.BoolVar(&debug, "debug", false, "enable debug logging")
type DebugTransport struct {
	// Enabled controls whether debug output is produced. Checked at
	// request time, so a flag parsed after construction still works.
	Enabled *bool
	// Next is the underlying transport. If nil, [http.DefaultTransport] is used.
	Next http.RoundTripper
}

func (t *DebugTransport) next() http.RoundTripper {
	if t.Next != nil {
		return t.Next
	}
	return http.DefaultTransport
}

// RoundTrip implements [http.RoundTripper].
func (t *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Enabled == nil || !*t.Enabled {
		return t.next().RoundTrip(req)
	}
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump request for debug: %w", err)
	}
	prettyPrintDump(os.Stderr, requestDump, "> ")
	resp, err := t.next().RoundTrip(req)
	if err != nil {
		return nil, err
	}
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump response for debug: %w", err)
	}
	prettyPrintDump(os.Stderr, responseDump, "< ")
	return resp, nil
}

func prettyPrintDump(w io.Writer, dump []byte, prefix string) {
	var output bytes.Buffer
	output.Grow(len(dump) * 2)
	for line := range bytes.Lines(dump) {
		output.WriteString(prefix)
		output.Write(line)
	}
	output.WriteByte('\n')
	_, _ = w.Write(output.Bytes())
}
