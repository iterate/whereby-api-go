package whereby

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetMeeting returns the specified meeting.
//
// See https://whereby.dev/http-api/#/paths/~1meetings~1{meetingId}/get for more
// details.
func (c *Client) GetMeeting(meetingID string) (*GetMeetingOutput, error) {
	return c.GetMeetingWithContext(context.Background(), meetingID)
}

// GetMeetingWithContext is the same as GetMeeting with a user-specified
// context.
func (c *Client) GetMeetingWithContext(ctx context.Context, meetingID string) (*GetMeetingOutput, error) {
	endpoint := strings.Replace(getMeetingEndpoint, "{meetingId}", meetingID, -1)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed create request: %w", err)
	}

	req.Header.Set("content-type", "application/json")
	res, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to the Whereby API: %w", err)
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status %d from Whereby", res.StatusCode)
	}

	var out meeting
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("failed to decode payload from Whereby: %w", err)
	}

	return createGetMeetingOutput(out)
}
