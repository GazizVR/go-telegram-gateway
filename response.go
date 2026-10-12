package telegramgateway

import (
	"encoding/json"
	"fmt"
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
			return &UnknownStatusError{StatusCode: statusCode}
		}
		return fmt.Errorf("decode response: %w", err)
	}
	if r.Error != nil {
		return parseError(*r.Error)
	}
	if !isSuccessStatus {
		return &UnknownStatusError{StatusCode: statusCode}
	}
	if !r.OK {
		return &UnknownAPIError{}
	}
	if result != nil && len(r.Result) > 0 {
		if err := json.Unmarshal(r.Result, result); err != nil {
			return fmt.Errorf("decode result: %w", err)
		}
	}
	return nil
}
