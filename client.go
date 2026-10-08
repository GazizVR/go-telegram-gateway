package telegramgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}
	return func(c *Client) {
		c.baseURL = u.String()
	}
}

const defaultBaseURL = "https://gatewayapi.telegram.org"

func NewClient(
	token string,
	options ...Option,
) *Client {
	client := &Client{
		token:      token,
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

func (c *Client) do(
	ctx context.Context,
	method string,
	body any,
	result any,
) error {
	endpoint, err := url.JoinPath(c.baseURL, method)
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBody, result); err != nil {
		return err
	}
	return nil
}
