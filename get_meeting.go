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
func (c *Client) GetMeeting(ctx context.Context, meetingID string, opts ...GetMeetingOpt) (GetMeetingOutput, error) {
	var out GetMeetingOutput
	var os getMeetingOpts
	for _, f := range opts {
		if err := f(&os); err != nil {
			return out, err
		}
	}

	endpoint := strings.Replace(getMeetingEndpoint, "{meetingId}", meetingID, -1)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return out, fmt.Errorf("failed create request: %w", err)
	}

	req.Header.Set("content-type", "application/json")
	res, err := c.do(req)
	if err != nil {
		return out, fmt.Errorf("failed to make request to the Whereby API: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return out, handleBadStatus(res)
	}

	var innerRes meeting
	if err := json.NewDecoder(res.Body).Decode(&innerRes); err != nil {
		return out, fmt.Errorf("failed to decode payload from Whereby: %w", err)
	}

	if err := createGetMeetingOutput(&out, innerRes); err != nil {
		return out, err
	}
	return out, nil
}

type getMeetingOpts struct {
	WithHostURL bool
}

// GetMeetingOpt is an option for GetMeeting.
type GetMeetingOpt func(*getMeetingOpts) error

func WithHostURL(include bool) GetMeetingOpt {
	return func(os *getMeetingOpts) error {
		os.WithHostURL = include
		return nil
	}
}
