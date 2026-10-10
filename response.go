package telegramgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// response is the envelope the Gateway API wraps every reply in:
// {"ok": true, "result": ...} or {"ok": false, "error": "..."}.
type response struct {
	OK     bool            `json:"ok"`
	Error  *string         `json:"error"`
	Result json.RawMessage `json:"result"`
}

// parseResponse turns an HTTP status code and body into an error or a result.
//
// It returns an *APIError if the status is not 2xx or the API reported
// "ok": false, and wraps ErrResponseParsing if a successful response cannot
// be decoded. The status is checked together with "ok" because a proxy may
// answer with a non-2xx status and an unrelated body. On success it decodes
// the "result" field into result, which must be a pointer or nil.
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
