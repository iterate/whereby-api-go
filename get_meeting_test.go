package whereby

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetMeeting(t *testing.T) {
	t.Run("With valid request/response", func(t *testing.T) {
		httpClient := &mockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			body := `{"meetingId":"1","startDate":"2020-05-12T16:42:49Z","endDate":"2020-05-12T17:42:49Z","roomUrl":"http://example.com","hostRoomUrl":"http://example.com/host-room"}`
			return &http.Response{
				StatusCode:    http.StatusCreated,
				Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
				ContentLength: int64(len(body)),
			}, nil
		}}
		c := Client{
			c: httpClient,
		}

		res, err := c.GetMeeting("1")
		if err != nil {
			t.Errorf("want no err; got %v", err)
		}

		if res.MeetingID != "1" {
			t.Errorf("res.MeetingID = %s; want %s", res.MeetingID, "1")
		}
		if got, want := res.HostURL, "http://example.com/host-room"; got != want {
			t.Errorf("res.HostUrl = %s; want %s", got, want)
		}
		if got, want := res.MeetingID, "1"; got != want {
			t.Errorf("res.MeetingID = %s; want %s", got, want)
		}
	})

	t.Run("With invalid response", func(t *testing.T) {
		httpClient := &mockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			body := `{"meetingId":"1","startDate":"2020-05-12T16:42:49Z",`
			return &http.Response{
				StatusCode:    http.StatusCreated,
				Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
				ContentLength: int64(len(body)),
			}, errors.New("something went wrong")
		}}
		c := Client{c: httpClient}

		_, err := c.GetMeeting("1")
		if err == nil {
			t.Errorf("want err; got %v", err)
		}
	})
}
