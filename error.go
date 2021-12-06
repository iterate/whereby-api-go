package whereby

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrNotFound = errors.New("meeting not found")
var ErrInvalidCredentials = errors.New("invalid credentials")

type RateLimitedError struct {
	timeLeft time.Duration
}

func (e RateLimitedError) Error() string {
	return "rate limited"
}

func (e RateLimitedError) TimeLeft() time.Duration {
	return e.timeLeft
}

func handleBadStatus(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusUnauthorized:
		return ErrInvalidCredentials
	case http.StatusTooManyRequests:
		var rlr rateLimitResponse
		if err := json.NewDecoder(r.Body).Decode(&rlr); err == nil && rlr.Data.MSLeft != nil {
			defer r.Body.Close()
			return RateLimitedError{timeLeft: time.Duration(*rlr.Data.MSLeft) * time.Millisecond}
		}
		return RateLimitedError{}
	default:
		return fmt.Errorf("unexpected status %d from Whereby", r.StatusCode)
	}
}

type rateLimitResponse struct {
	Description string `json:"error"`
	Data        struct {
		MSLeft *int `json:"ms_left"`
	} `json:"data,omitempty"`
}
