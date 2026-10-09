package telegramgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Ok    bool    `json:"ok"`
	Error *string `json:"error"`
}

func (c *Client) parseResponse(
	statusCode int,
	respBody []byte,
	result any,
) error {
	var r response
	if err := json.Unmarshal(respBody, &r); err != nil {
		if statusCode < 200 || statusCode >= 300 {
			apiErr := &APIError{StatusCode: statusCode, Message: http.StatusText(statusCode)}
			return apiErr
		}
		return fmt.Errorf("%w (status %d): %w", ErrResponseParsing, statusCode, err)
	}
	if !r.Ok {
		apiErr := &APIError{StatusCode: statusCode, Message: http.StatusText(statusCode)}
		if r.Error != nil {
			apiErr.Message = *r.Error
		}
		return apiErr
	}
	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("%w (status %d): %w", ErrResponseParsing, statusCode, err)
	}
	return nil
}
