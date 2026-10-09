package telegramgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
		return 0, nil, fmt.Errorf("build url: %w", err)
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return 0, nil, fmt.Errorf("encode request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return 0, nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("%w: %w", ErrNetwork, err)
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("%w: %w", ErrNetwork, err)
	}

	return resp.StatusCode, respBody, nil
}

func (c *Client) call(
	ctx context.Context,
	method string,
	body any,
	result any,
) error {
	statusCode, respBody, err := c.do(ctx, method, body)
	if err != nil {
		return err
	}
	return c.parseResponse(statusCode, respBody, result)
}
