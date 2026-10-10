// Package telegramgateway is an unofficial Go client for the Telegram Gateway
// API (https://core.telegram.org/gateway/api), which delivers verification
// codes to users through Telegram.
//
// Create a Client with NewClient and call its methods:
//
//	client, err := telegramgateway.NewClient("token")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	status, err := client.SendVerificationMessage(ctx, telegramgateway.SendVerificationMessageRequest{
//		PhoneNumber: "+998901234567",
//	})
//
// The package depends only on the Go standard library.
package telegramgateway

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// defaultBaseURL is the production address of the Telegram Gateway API.
const defaultBaseURL = "https://gatewayapi.telegram.org"

// Client is a Telegram Gateway API client.
//
// Use NewClient to create one; the zero value is not usable.
// A Client is safe for concurrent use as long as its underlying
// *http.Client is.
type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

// Option configures a Client. Options are passed to NewClient.
type Option func(*Client) error

// WithHTTPClient sets the HTTP client used to send requests.
// Use it to configure timeouts, proxies or transports.
// It returns an error if httpClient is nil.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.New("http client must not be nil")
		}
		c.httpClient = httpClient
		return nil
	}
}

// WithBaseURL overrides the API base URL (default
// "https://gatewayapi.telegram.org"). It is mostly useful for tests that run
// against a local server. The URL must be absolute, use the http or https
// scheme and contain a host.
func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		u, err := url.ParseRequestURI(raw)
		if err != nil {
			return fmt.Errorf("invalid base url: %w", err)
		}
		if u.Host == "" {
			return errors.New("url host must not be empty")
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.New("url scheme must be http or https")
		}
		c.baseURL = u.String()
		return nil
	}
}

// NewClient creates a Client that authenticates with the given API token,
// which can be obtained in the Telegram Gateway account settings.
//
// It returns an error if the token is empty or any option is invalid.
// By default the client uses http.DefaultClient, which has no timeout, so
// pass a context with a deadline to every call or set your own client with
// WithHttpClient.
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
