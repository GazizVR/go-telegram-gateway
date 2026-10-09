package telegramgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Ok     bool            `json:"ok"`
	Error  *string         `json:"error"`
	Result json.RawMessage `json:"result"`
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
	if result != nil && len(r.Result) > 0 {
		if err := json.Unmarshal(r.Result, result); err != nil {
			return fmt.Errorf("%w (status %d): %w", ErrResponseParsing, statusCode, err)
		}
	}
	return nil
}
