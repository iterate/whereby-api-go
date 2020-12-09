# whereby-api-go

A [Whereby](https://whereby.com/) API client for Go.

## Installation

```bash
go get github.com/iterate/whereby-api-go
```

## Usage

```go
wb := whereby.NewClient("my-api-key")
meeting, err := wb.CreateMeeting(whereby.CreateMeetingInput{
    Start: time.Now(),
    End:   time.Now().Add(time.Hour),
})
```

## Legal
Copyright (C) Mindcare AS.

Unauthorized use, copying or distribution of this library, via any medium, is strictly prohibited.