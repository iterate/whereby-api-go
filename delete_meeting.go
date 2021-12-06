package whereby

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// DeleteMeeting deletes the specified meeting. The endpoint is idempotent,
// meaning it will return the same response even if the meeting has already been
// deleted.
//
// See https://whereby.dev/http-api/#/paths/~1meetings~1{meetingId}/delete for
// more details.
func (c *Client) DeleteMeeting(ctx context.Context, meetingID string) error {
	endpoint := strings.Replace(deleteMeetingEndpoint, "{meetingId}", meetingID, -1)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed create request: %w", err)
	}

	res, err := c.do(req)
	if err != nil {
		return fmt.Errorf("failed to make request to the Whereby API: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return handleBadStatus(res)
	}

	return nil
}
