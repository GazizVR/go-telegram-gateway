package telegramgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	OK     bool            `json:"ok"`
	Error  *string         `json:"error"`
	Result json.RawMessage `json:"result"`
}

func parseResponse(
	statusCode int,
	respBody []byte,
	result any,
) error {
	var r response
	isSuccessStatus := statusCode >= 200 && statusCode < 300
	if err := json.Unmarshal(respBody, &r); err != nil {
		if !isSuccessStatus {
			apiErr := &APIError{
				StatusCode: statusCode,
				Message:    http.StatusText(statusCode),
			}
			return apiErr
		}
		return fmt.Errorf("%w (status %d): %w", ErrResponseParsing, statusCode, err)
	}
	if !r.OK || !isSuccessStatus {
		apiErr := &APIError{
			StatusCode: statusCode,
			Message:    http.StatusText(statusCode),
		}
		if r.Error != nil {
			apiErr.Message = *r.Error
		}
		if apiErr.Message == "" {
			apiErr.Message = "unexpected error"
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
