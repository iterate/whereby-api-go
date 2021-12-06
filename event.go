package whereby

import "time"

type RoomNamePattern string
type RoomMode string

const (
	DefaultNamePattern       RoomNamePattern = ""
	UUIDNamePattern          RoomNamePattern = "uuid"
	HumanReadableNamePattern RoomNamePattern = "human-short"
)

const (
	DefaultMode RoomMode = ""
	GroupMode   RoomMode = "group"
	NormalMode  RoomMode = "normal"
)

type CreateMeetingInput struct {
	// The initial lock state of the room. If true, only hosts will be able to
	// let in other participants and change lock state.
	IsLocked bool

	// This will be used as the prefix for the room name. The string should be
	// lowercase, and spaces will be automatically removed.
	RoomNamePrefix string

	// The format of the randomly generated room name. uuid is the default room
	// name pattern and follows the usual 8-4-4-4-12 pattern. human-short
	// generates a shorter string made up of six distinguishable characters.
	RoomNamePattern RoomNamePattern

	// The mode of the created transient room. normal is the default room mode
	// and should be used for meetings up to 4 participants. group should be
	// used for meetings that require more than 4 participants.
	RoomMode RoomMode

	// When the meeting starts. By default in UTC but a timezone can be
	// specified, e.g. 2021-05-07T17:42:49-05:00. This date must be in
	// the future.
	//
	// Deprecated: Value is ignored
	Start time.Time

	// When the meeting ends. By default in UTC but a timezone can be
	// specified, e.g. 2021-05-07T17:42:49-05:00. This has to be the same or
	// after the current date.
	End time.Time

	// Include hostRoomUrl field in the meeting response.
	WithHostURL bool
}

type GetMeetingOutput struct {
	// The ID of the meeting.
	MeetingID string

	// When the meeting starts. Always in UTC, regardless of the input timezone.
	Start time.Time

	// When the meeting ends. Always in UTC, regardless of the input timezone.
	End time.Time

	// The URL to room where the meeting will be hosted.
	URL string

	// The URL to room where the meeting will be hosted which will also make
	// the user the host of the meeting. A host will get additional privileges
	// like locking the room, and removing and muting participants, so you
	//should be careful with whom you share this URL. The user will only become
	// a host if the meeting is on-going (some additional slack is added to
	// allow a host joining the meeting ahead of time or if the meeting goes
	// over time). This field is optional and will only provided if requested
	// through fields parameter.
	HostURL string
}

type meeting struct {
	MeetingId   string  `json:"meetingId"`
	StartDate   *string `json:"startDate,omitempty"`
	EndDate     string  `json:"endDate"`
	RoomURL     string  `json:"roomUrl"`
	HostRoomURL *string `json:"hostRoomUrl,omitempty"`
}
