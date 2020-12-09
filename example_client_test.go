package whereby

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Create a meeting and print the URL to stdout
func ExampleClient_CreateMeeting() {
	wb := NewClient(os.Getenv("WHEREBY_API_KEY"))

	meeting, err := wb.CreateMeeting(CreateMeetingInput{
		Start: time.Now(),
		End:   time.Now().Add(time.Hour),
	})

	if err != nil {
		log.Printf("Something went wrong: %v\n", err)
	}

	fmt.Println(meeting.URL)
}

