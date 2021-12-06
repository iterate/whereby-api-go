//go:build functional
// +build functional

package whereby_test

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/iterate/whereby-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var apiKey = flag.String("whereby-api-key", "", "Whereby API key")

func TestClient(t *testing.T) {
	if *apiKey == "" {
		t.Fatal("no API key defined")
	}
	ctx := context.Background()
	if dl, ok := t.Deadline(); ok {
		c2, ccl := context.WithDeadline(ctx, dl)
		ctx = c2
		defer ccl()
	}
	client := whereby.NewClient(*apiKey)

	t.Run("Create, get, delete", func(t *testing.T) {
		createRes, err := client.CreateMeeting(ctx, whereby.CreateMeetingInput{
			IsLocked:        true,
			RoomNamePrefix:  "some-prefix-",
			RoomNamePattern: whereby.UUIDNamePattern,
			RoomMode:        whereby.NormalMode,
			End:             time.Now().Add(time.Hour * 2),
			WithHostURL:     true,
		})
		require.NoError(t, err)

		getRes, err := client.GetMeeting(ctx, createRes.MeetingID, whereby.WithHostURL(true))
		assert.NoError(t, err)
		assert.Equal(t, createRes.MeetingID, getRes.MeetingID)

		err = client.DeleteMeeting(ctx, getRes.MeetingID)
		assert.NoError(t, err)

		_, err = client.GetMeeting(ctx, getRes.MeetingID)
		assert.ErrorIs(t, err, whereby.ErrNotFound)
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		badClient := whereby.NewClient("invalid-credentials")

		_, err := badClient.GetMeeting(ctx, "2cba3bec-b78e-4ae3-a2cc-d781de733d2a")
		assert.ErrorIs(t, err, whereby.ErrInvalidCredentials)
	})
}
