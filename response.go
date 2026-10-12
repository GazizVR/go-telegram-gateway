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
// If decoding the response envelope fails:
//   - returns *UnknownStatusError if statusCode is non-2xx;
//   - returns a wrapped decode error if statusCode is 2xx.
//
// If the API reported an error code in the envelope, it returns a mapped sentinel error,
// *FloodError, or *UnknownAPIError.
// If the HTTP status is not 2xx, it returns *UnknownStatusError.
// If "ok" is false, it returns *UnknownAPIError.
// On success, it decodes the "result" field into result (if non-nil and present).
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
