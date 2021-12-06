package whereby

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestClient_CreateMeeting(t *testing.T) {
	httpClient := &mockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
		body := `{"meetingId":"1","startDate":"2020-05-12T16:42:49Z","endDate":"2020-05-12T17:42:49Z","roomUrl":"http://example.com"}`
		return &http.Response{
			StatusCode:    http.StatusCreated,
			Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
			ContentLength: int64(len(body)),
		}, nil
	}}
	c := Client{
		c: httpClient,
	}

	res, err := c.CreateMeeting(context.Background(), CreateMeetingInput{
		IsLocked:    true,
		Start:       mustTimeFunc(t)(time.Parse(time.RFC3339, "2020-05-12T16:42:49Z")),
		End:         mustTimeFunc(t)(time.Parse(time.RFC3339, "2020-05-12T17:42:49Z")),
		WithHostURL: false,
	})
	if err != nil {
		t.Errorf("want no err; got %v", err)
	}

	if res.MeetingID != "1" {
		t.Errorf("res.MeetingID = %s; want %s", res.MeetingID, "1")
	}

}
