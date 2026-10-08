package telegramgateway

import "net/http"

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
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
