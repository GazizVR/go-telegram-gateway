package telegramgateway

import (
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (err *APIError) Error() string {
	return fmt.Sprintf(
		"telegramgateway (status: %d): %s",
		err.StatusCode,
		err.Message,
	)
}

var (
	ErrResponseParsing = errors.New("telegramgateway: parse response")
)
