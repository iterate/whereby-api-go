package wherebyiface

import (
	"context"
	"github.com/iterate/whereby-api-go"
)

// Client is an interface for the Whereby client that can be used for mocking
// and dependency injection.
type Client interface {
	CreateMeeting(input whereby.CreateMeetingInput) (*whereby.GetMeetingOutput, error)
	CreateMeetingWithContext(ctx context.Context, input whereby.CreateMeetingInput) (*whereby.GetMeetingOutput, error)

	DeleteMeeting(meetingId string) error
	DeleteMeetingWithContext(ctx context.Context, meetingId string) error

	GetMeeting(meetingId string) (*whereby.GetMeetingOutput, error)
	GetMeetingWithContext(ctx context.Context, meetingId string) (*whereby.GetMeetingOutput, error)
}
