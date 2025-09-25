package soap

import (
	"context"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestDefaultShouldRetry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		err        error
		statusCode int
		expected   bool
	}{
		{
			name:     "DNS error should retry",
			err:      &net.DNSError{Err: "no such host"},
			expected: true,
		},
		{
			name:       "500 error should retry for idempotent requests",
			statusCode: http.StatusInternalServerError,
			expected:   true, // Will be true because we set idempotency key
		},
		{
			name:       "400 error should not retry",
			statusCode: http.StatusBadRequest,
			expected:   false,
		},
		{
			name:       "429 error should retry",
			statusCode: http.StatusTooManyRequests,
			expected:   true,
		},
		{
			name:       "503 error should retry for idempotent requests",
			statusCode: http.StatusServiceUnavailable,
			expected:   true, // Will be true because we set idempotency key
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "http://example.com", nil)
			req.Header.Set("Idempotency-Key", "test") // Make request idempotent

			var resp *http.Response
			if tt.statusCode != 0 {
				resp = &http.Response{StatusCode: tt.statusCode, Header: make(http.Header)}
			}

			result := DefaultCheckRetry(tt.err, req, resp)
			if result != tt.expected {
				t.Errorf("defaultShouldRetry() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRetryTransport_Success(t *testing.T) {
	t.Parallel()

	requestCount := 0
	const failCount = 2 // Fail 2 times, succeed on the 3rd attempt

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount <= failCount {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("Service Unavailable"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Success"))
	}))
	defer server.Close()

	transport := &retryTransport{
		maxRetries:  3,
		next:        http.DefaultTransport,
		shouldRetry: DefaultCheckRetry,
	}

	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("POST", server.URL, strings.NewReader("test"))
	req.Header.Set("Idempotency-Key", "test")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Expected success after retries, got error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	if requestCount != failCount+1 {
		t.Errorf("Expected %d requests, got %d", failCount+1, requestCount)
	}
}

func TestRetryTransport_NonRetryableError(t *testing.T) {
	t.Parallel()

	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusBadRequest) // 400 should not be retried
		_, _ = w.Write([]byte("Bad Request"))
	}))
	defer server.Close()

	transport := &retryTransport{
		maxRetries:  3,
		next:        http.DefaultTransport,
		shouldRetry: DefaultCheckRetry,
	}

	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("POST", server.URL, strings.NewReader("test"))

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", resp.StatusCode)
	}

	// Should only make 1 request since 400 is not retryable
	if requestCount != 1 {
		t.Errorf("Expected 1 request for non-retryable error, got %d", requestCount)
	}
}

func TestRetryTransport_MaxRetriesExceeded(t *testing.T) {
	t.Parallel()

	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusServiceUnavailable) // Always fail
		_, _ = w.Write([]byte("Service Unavailable"))
	}))
	defer server.Close()

	transport := &retryTransport{
		maxRetries:  2, // Only allow 2 retries
		next:        http.DefaultTransport,
		shouldRetry: DefaultCheckRetry,
	}

	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("POST", server.URL, strings.NewReader("test"))
	req.Header.Set("Idempotency-Key", "test")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status 503, got %d", resp.StatusCode)
	}

	// Should make 3 requests total (1 initial + 2 retries)
	expectedRequests := 3
	if requestCount != expectedRequests {
		t.Errorf("Expected %d requests, got %d", expectedRequests, requestCount)
	}
}

func TestRetryTransport_ContextCancellation(t *testing.T) {
	t.Parallel()

	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusServiceUnavailable) // Always fail to trigger retries
		_, _ = w.Write([]byte("Service Unavailable"))
	}))
	defer server.Close()

	transport := &retryTransport{
		maxRetries:  5,
		next:        http.DefaultTransport,
		shouldRetry: DefaultCheckRetry,
	}

	client := &http.Client{Transport: transport}

	// Create context with short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", server.URL, strings.NewReader("test"))
	req.Header.Set("Idempotency-Key", "test")

	start := time.Now()
	_, err := client.Do(req)
	elapsed := time.Since(start)

	// Should get context deadline exceeded error
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded, got: %v", err)
	}

	// Should have completed quickly due to context cancellation
	if elapsed > 100*time.Millisecond {
		t.Errorf("Request took too long (%v), context cancellation may not be working", elapsed)
	}

	// Should have made at least one request
	if requestCount < 1 {
		t.Errorf("Expected at least 1 request, got %d", requestCount)
	}
}

func TestRetryTransport_RetryAfterHeader(t *testing.T) {
	t.Parallel()

	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount == 1 {
			w.Header().Set("Retry-After", "1") // 1 second
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte("Too Many Requests"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Success"))
	}))
	defer server.Close()

	transport := &retryTransport{
		maxRetries:  1,
		next:        http.DefaultTransport,
		shouldRetry: DefaultCheckRetry,
	}

	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("POST", server.URL, strings.NewReader("test"))

	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Expected success after retry, got error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Should respect Retry-After header (at least some delay, but allow for jitter and timing variations)
	if elapsed < 500*time.Millisecond {
		t.Errorf("Expected significant delay due to Retry-After header, got %v", elapsed)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests, got %d", requestCount)
	}
}

func TestClient_WithCheckRetry(t *testing.T) {
	t.Parallel()

	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount == 1 {
			w.WriteHeader(http.StatusInternalServerError) // 500 error
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}
		// Success response
		respEnv, _ := NewEnvelope(WithBody([]byte(`<response>Success</response>`)))
		respXML, _ := xml.Marshal(respEnv)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respXML)
	}))
	defer server.Close()

	// Test with custom retry logic
	customRetry := func(err error, req *http.Request, resp *http.Response) bool {
		// Custom retry logic that always retries 5xx errors (more aggressive than DefaultCheckRetry)
		if err != nil {
			return DefaultCheckRetry(err, req, resp)
		}
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			return true // Always retry 5xx errors
		}
		return DefaultCheckRetry(err, req, resp)
	}

	client, err := NewClient(
		WithEndpoint(server.URL),
		WithMaxRetries(1),
		WithCheckRetry(customRetry), // Use custom retry logic
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	reqEnv, _ := NewEnvelope(WithBody([]byte(`<request>Test</request>`)))
	ctx := context.Background()
	respEnv, err := client.Call(ctx, "", reqEnv)
	if err != nil {
		t.Fatalf("Expected success after retry with custom logic, got error: %v", err)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests (1 failure + 1 success with custom retry), got %d", requestCount)
	}

	if string(respEnv.Body.Content) != `<response>Success</response>` {
		t.Errorf("Unexpected response body: %s", string(respEnv.Body.Content))
	}
}
