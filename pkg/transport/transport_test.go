package transport_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	client "github.com/grafana/grafana-openapi-client-go/client"
)

// TestRetryBackoffRespectsContextCancellation asserts that when an API operation context
// expires, the retry backoff sleep/delay is cancelled and the call returns promptly and
// does not attempt to make any further requests beyond the deadline.
//
// This test uses a shortened DefaultTimeout (100ms) so the deadline fires during the first
// retry sleep (2s). The call should return promptly once the context expires, not continue
// sleeping for the full backoff duration.
func TestRetryBackoffRespectsContextCancellation(t *testing.T) {
	// Do not run in parallel: this test temporarily mutates the package-level
	// DefaultTimeout variable.
	const numRetries = 2

	// Create a mock server that returns 503 on every request.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	// Shorten the per-operation context deadline to 100ms - arbitrarily short.
	// This simulates the scenario where the code would have slept through a context deadline
	origTimeout := httptransport.DefaultTimeout
	httptransport.DefaultTimeout = 100 * time.Millisecond
	defer func() { httptransport.DefaultTimeout = origTimeout }()

	u, _ := url.Parse(server.URL)
	grafanaClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:       u.Host,
		BasePath:   client.DefaultBasePath,
		Schemes:    []string{"http"},
		NumRetries: numRetries,
		// RetryTimeout = 0 → exponential backoff (2^n seconds). Production default.
	})

	done := make(chan error, 1)
	go func() {
		_, err := grafanaClient.Health.GetHealth() // This could be any operation
		done <- err
	}()

	// Once the context deadline (100ms) fires, the backoff sleep should be cancelled
	// and the call should return with an error well within 1 second.
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("expected an error after context expiry, got nil")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("call did not return within 1s of context expiry — " +
			"the retry backoff sleep is not context-aware and is blocking indefinitely")
	}
}
