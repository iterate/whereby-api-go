package whereby

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// mustTimeFunc is a helper function for dealing with time parsing in test
// functions.
func mustTimeFunc(t *testing.T) func(time.Time, error) time.Time {
	return func(v time.Time, err error) time.Time {
		if err != nil {
			t.Fatalf("time parsing failed: %v", err)
		}
		return v
	}
}

// MockClient is a mock HTTP client.
type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func testContext(t *testing.T) (context.Context, context.CancelFunc) {
	if ddl, ok := t.Deadline(); ok {
		return context.WithDeadline(context.Background(), ddl)
	}
	return context.WithCancel(context.Background())
}
