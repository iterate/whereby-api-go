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

type createMeetingPayload struct {
	IsLocked        bool     `json:"isLocked"`
	RoomNamePrefix  string   `json:"roomNamePrefix,omitempty"`
	RoomNamePattern string   `json:"roomNamePattern,omitempty"`
	RoomMode        string   `json:"roomMode,omitempty"`
	Start           string   `json:"startDate,omitempty"`
	End             string   `json:"endDate"`
	Fields          []string `json:"fields,omitempty"`
}

type CreateMeetingOutput = GetMeetingOutput

// CreateMeeting creates a meeting as specified. It will also create a transient
// room that is guaranteed to be available for specified start and end time.
// Some time after the meeting has ended, the transient room will be
// automatically deleted. The URL to this room is present in the response.
//
// See https://whereby.dev/http-api/#/paths/~1meetings/post for more details.
func (c *Client) CreateMeeting(ctx context.Context, input CreateMeetingInput) (CreateMeetingOutput, error) {
	var out CreateMeetingOutput
	if err := validateCreateMeetingInput(input); err != nil {
		return out, err
	}

	payload, err := json.Marshal(c.getCreateMeetingInput(input))
	if err != nil {
		return out, fmt.Errorf("failed to encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, createMeetingEndpoint, bytes.NewBuffer(payload))
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

// createGetMeetingOutput creates the user-friendly output object from the
// internal JSON representation.
func createGetMeetingOutput(dst *GetMeetingOutput, src meeting) error {
	dst.MeetingID = src.MeetingId
	dst.URL = src.RoomURL
	if hu := src.HostRoomURL; hu != nil {
		dst.HostURL = *hu
	}

	if sd := src.StartDate; sd != nil && *sd != "" {
		start, err := time.Parse(time.RFC3339, *sd)
		if err != nil {
			return err
		}
		dst.Start = start
	}

	end, err := time.Parse(time.RFC3339, src.EndDate)
	if err != nil {
		return err
	}
	dst.End = end

	return nil
}

// getCreateMeetingInput converts the CreateMeetingInput object into the inner
// representation for JSON marshalling.
func (c *Client) getCreateMeetingInput(in CreateMeetingInput) createMeetingPayload {
	var out createMeetingPayload
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
	if input.RoomNamePrefix != "" {
		if strings.ToLower(input.RoomNamePrefix) != input.RoomNamePrefix {
			return errors.New("room name should be lowercase")
		}
	}

	if input.End.IsZero() {
		return errors.New("meeting end time must be specified")
	}

	return nil
}
