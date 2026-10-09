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
	return fmt.Sprintf("telegramgateway: %s (status: %d)", err.Message, err.StatusCode)
}

var (
	ErrNetwork         = errors.New("telegramgateway: network error")
	ErrResponseParsing = errors.New("telegramgateway: parse response")
)
