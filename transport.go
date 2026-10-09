package telegramgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) do(
	ctx context.Context,
	method string,
	body any,
) (statusCode int, respBody []byte, err error) {
	endpoint, err := url.JoinPath(c.baseURL, method)
	if err != nil {
		return 0, nil, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, respBody, nil
}
