package telegramgateway

import (
	"errors"
	"fmt"
)

// APIError is returned when the API rejects a request or answers with a
// non-2xx HTTP status. Use errors.As to access it:
//
//	var apiErr *telegramgateway.APIError
//	if errors.As(err, &apiErr) {
//		fmt.Println(apiErr.StatusCode, apiErr.Message)
//	}
type APIError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int
	// Message is the error text returned by the API, for example
	// "PHONE_NUMBER_INVALID", or the HTTP status text if the API sent none.
	Message string
}

// Error implements the error interface.
func (err *APIError) Error() string {
	return fmt.Sprintf("telegramgateway: %s (status: %d)", err.Message, err.StatusCode)
}

var (
	// ErrNetwork is returned when the request could not be sent or the
	// response could not be read. It also wraps context errors, so check for
	// context.Canceled and context.DeadlineExceeded with errors.Is before
	// deciding to retry.
	ErrNetwork = errors.New("telegramgateway: network error")

	// ErrResponseParsing is returned when the API answered with a successful
	// status but the response could not be decoded. The request may have been
	// processed, so do not blindly retry it.
	ErrResponseParsing = errors.New("telegramgateway: parse response")
)
