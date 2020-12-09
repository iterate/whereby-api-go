package whereby

import (
	"net/http"
	"testing"
)

func TestClient_DeleteMeeting(t *testing.T) {
	httpClient := &mockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode:    http.StatusNoContent,
			ContentLength: 0,
		}, nil
	}}
	c := Client{
		c: httpClient,
	}

	err := c.DeleteMeeting("1")
	if err != nil {
		t.Errorf("want no err; got %v", err)
	}
}
