# whereby-api-go

A [Whereby](https://whereby.com/) API client for Go.

## Installation

```bash
go get github.com/iterate/whereby-api-go
```

## Usage

```go
package main

import (
	"github.com/iterate/whereby-api-go"
)

func main() {
    wb := whereby.NewClient("my-api-key")
    meeting, err := wb.CreateMeeting(whereby.CreateMeetingInput{
        Start: time.Now(),
        End:   time.Now().Add(time.Hour),
    })
    
    fmt.Println(meeting.URL)
}

```

## Legal
Copyright (c) 2020 Mindcare AS.

Developed by [Iterate](https://iterate.no).

Licensed under the [MIT license](LICENSE.txt).
