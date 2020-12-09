/*
Package whereby contains a client for communicating with the Whereby API.
*/
package whereby

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

// Client is the Whereby client. It must be created using NewClient.
type Client struct {
	key string
	c   HTTPClient
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	getMeetingEndpoint    = "https://api.whereby.dev/v1/meetings/{meetingId}"
	createMeetingEndpoint = "https://api.whereby.dev/v1/meetings"
	deleteMeetingEndpoint = "https://api.whereby.dev/v1/meetings/{meetingId}"
)

// NewClient creates a new Whereby client.
func NewClient(key string) *Client {
	return &Client{key: key}
}

// do wraps http.Client.Do and creates the inner http.Client if it is not set.
func (c *Client) do(r *http.Request) (*http.Response, error) {
	httpClient := c.c
	if httpClient == nil {
		httpClient = c.getClient()
		c.c = httpClient
	}

	return httpClient.Do(r)
}

// getClient creates a http.Client used internally for communication with
// Whereby
func (c *Client) getClient() *http.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: c.key,
		TokenType:   "bearer",
	})
	return oauth2.NewClient(context.Background(), ts)
}
