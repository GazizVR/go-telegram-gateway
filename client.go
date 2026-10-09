package telegramgateway

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client) error

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.New("http client must not be nil")
		}
		c.httpClient = httpClient
		return nil
	}
}

func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		u, err := url.Parse(raw)
		if err != nil {
			return fmt.Errorf("invalid base url: %w", err)
		}
		c.baseURL = u.String()
		return nil
	}
}

const defaultBaseURL = "https://gatewayapi.telegram.org"

func NewClient(
	token string,
	options ...Option,
) (*Client, error) {
	if token == "" {
		return nil, errors.New("token must not be empty")
	}
	client := &Client{
		token:      token,
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
	}
	for _, opt := range options {
		if err := opt(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}
