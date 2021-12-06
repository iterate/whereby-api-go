package whereby_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iterate/whereby-api-go"
)

// Create a meeting and print the URL to stdout
func ExampleClient_CreateMeeting() {
	wb := whereby.NewClient(os.Getenv("WHEREBY_API_KEY"))

	meeting, err := wb.CreateMeeting(context.Background(), whereby.CreateMeetingInput{
		End: time.Now().Add(time.Hour),
	})

	if err != nil {
		log.Printf("Something went wrong: %v\n", err)
	}

	fmt.Println(meeting.URL)
}
