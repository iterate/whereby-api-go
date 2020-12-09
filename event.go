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
	IsLocked        bool
	RoomNamePrefix  string
	RoomNamePattern RoomNamePattern
	RoomMode        RoomMode
	Start           time.Time
	End             time.Time
	WithHostURL     bool
}

type GetMeetingOutput struct {
	MeetingID string
	Start     time.Time
	End       time.Time
	URL       string
	HostURL   string
}

type meeting struct {
	MeetingId   string `json:"meetingId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	RoomURL     string `json:"roomUrl"`
	HostRoomURL string `json:"hostRoomUrl,omitempty"`
}
