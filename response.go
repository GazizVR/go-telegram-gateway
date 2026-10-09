package telegramgateway

import (
	"encoding/json"
	"errors"
)

type Response struct {
	Ok bool `json:"ok"`
}

type ErrorResponse struct {
	Response
	Error string `json:"error"`
}

func (c *Client) parseResponse(
	statusCode int,
	respBody []byte,
	result any,
) error {
	if statusCode >= 200 && statusCode < 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			return err
		}
		return errors.New(errResp.Error)
	}
	if err := json.Unmarshal(respBody, result); err != nil {
		return err
	}
	return nil
}
