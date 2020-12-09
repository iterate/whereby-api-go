package whereby

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type createMeetingInInt struct {
	IsLocked        bool     `json:"isLocked"`
	RoomNamePrefix  string   `json:"roomNamePrefix,omitempty"`
	RoomNamePattern string   `json:"roomNamePattern,omitempty"`
	RoomMode        string   `json:"roomMode,omitempty"`
	Start           string   `json:"startDate"`
	End             string   `json:"endDate"`
	Fields          []string `json:"fields,omitempty"`
}

// CreateMeeting creates a meeting as specified. It will also create a transient
// room that is guaranteed to be available for specified start and end time.
// Some time after the meeting has ended, the transient room will be
// automatically deleted. The URL to this room is present in the response.
//
// See https://whereby.dev/http-api/#/paths/~1meetings/post for more details.
func (c *Client) CreateMeeting(input CreateMeetingInput) (*GetMeetingOutput, error) {
	return c.CreateMeetingWithContext(context.Background(), input)
}

// CreateMeetingWithContext is the same as CreateMeeting with a user-specified
// context.
func (c *Client) CreateMeetingWithContext(ctx context.Context, input CreateMeetingInput) (*GetMeetingOutput, error) {
	if err := validateCreateMeetingInput(input); err != nil {
		return nil, err
	}

	payload, err := json.Marshal(c.getCreateMeetingInput(input))
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, createMeetingEndpoint, bytes.NewBuffer(payload))
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

// createGetMeetingOutput creates the user-friendly output object from the
// internal JSON representation.
func createGetMeetingOutput(in meeting) (*GetMeetingOutput, error) {
	var out GetMeetingOutput

	out.MeetingID = in.MeetingId
	out.URL = in.RoomURL
	out.HostURL = in.HostRoomURL

	start, err := time.Parse(time.RFC3339, in.StartDate)
	if err != nil {
		return &out, fmt.Errorf("failed to parse meeting start time %s: %w", in.StartDate, err)
	}
	out.Start = start

	end, err := time.Parse(time.RFC3339, in.EndDate)
	if err != nil {
		return &out, fmt.Errorf("failed to parse meeting end time %s: %w", in.EndDate, err)
	}
	out.End = end

	return &out, nil
}

// getCreateMeetingInput converts the CreateMeetingInput object into the inner
// representation for JSON marshalling.
func (c *Client) getCreateMeetingInput(in CreateMeetingInput) createMeetingInInt {
	var out createMeetingInInt
	out.IsLocked = in.IsLocked
	out.RoomNamePrefix = in.RoomNamePrefix
	out.RoomNamePattern = string(in.RoomNamePattern)
	out.RoomMode = string(in.RoomMode)
	out.Start = in.Start.Format(time.RFC3339)
	out.End = in.End.Format(time.RFC3339)
	if in.WithHostURL {
		out.Fields = append(out.Fields, "hostRoomUrl")
	}

	return out
}

// validateCreateMeetingInput validates the provided CreateMeetingInput.
func validateCreateMeetingInput(input CreateMeetingInput) error {
	if input.RoomNamePrefix != "" && !strings.HasPrefix(input.RoomNamePrefix, "/") {
		return errors.New(`room name prefix must begin with a slash ("/")`)
	}

	if input.Start.IsZero() || input.End.IsZero() {
		return errors.New("both start and end times must be specified")
	}

	return nil
}
