package telegramgateway

import (
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
