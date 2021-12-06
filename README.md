# whereby-api-go

A [Whereby](https://whereby.com/) API client for Go.

See the [pkg.go.dev](https://pkg.go.dev/github.com/iterate/whereby-api-go) for more documentation for this module. 

## Installation

```bash
go get github.com/iterate/whereby-api-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iterate/whereby-api-go"
)

func main() {
	wb := whereby.NewClient(os.Getenv("WHEREBY_API_KEY"))

	meeting, err := wb.CreateMeeting(context.Background(), whereby.CreateMeetingInput{
		End: time.Now().Add(time.Hour),
	})

	if err != nil {
		log.Printf("Something went wrong: %v\n", err)
	}

	fmt.Println(meeting.URL)
}
```

## Contributing
You may run a simple functional test by passing the `functional` tag to go test and setting the `-whereby-api-key` flag:

```shell
go test ./... -tags functional -whereby-api-key key
```

## Legal
Copyright (c) 2020 Mindcare AS.

Developed by [Iterate](https://iterate.no).

Licensed under the [MIT license](LICENSE.txt).

