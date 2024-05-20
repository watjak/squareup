package squareup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "2024-05-15"
	defaultBaseURL = "https://connect.squareupsandbox.com/v2/"
	userAgent      = "squareup/" + libraryVersion
	mediaType      = "application/json"

	headerRateLimit     = "RateLimit-Limit"
	headerRateRemaining = "RateLimit-Remaining"
	headerRateReset     = "RateLimit-Reset"
	headerRequestID     = "x-request-id"
)

// Client manages communication with the Square Connect API.
type Client struct {
	// HTTP client used to communicate with the API.
	HTTPClient *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the Square API.
	UserAgent string

	// Services used for talking to different parts of the Square API.

	// Optional extra HTTP headers to set on every request to the API.
	headers map[string]string
}

// ListOptions struct is used for passing query parameters to list endpoints.
type ListOptions struct {
	// A cursor for use in pagination. If a cursor is not present, it is assumed to be the start of the list.
	Cursor string `url:"cursor,omitempty"`

	// The maximum number of results to return in a single page. This value cannot exceed 100.
	Limit int `url:"limit,omitempty"`
}

type ErrorResponse struct {
	// Http Response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`

	// RequestID is the unique identifier for the request
	RequestID string `json:"request_id"`
}

// NewFromToken creates a new Square client from a given access token.
func NewFromToken(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
		UserAgent:  userAgent,
		headers:    map[string]string{"Authorization": "Bearer " + token},
	}
}

// NewClient creates a new Square client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
		UserAgent:  userAgent,
	}

	c.headers = make(map[string]string)

	return c
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New creates a new Square client.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}

// SetRequestHeaders sets optional HTTP headers on the client that are
// sent on each HTTP request.
func SetRequestHeaders(headers map[string]string) ClientOpt {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}
		return nil
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}

	default:
		buf := new(bytes.Buffer)
		if body != nil {
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(method, u.String(), buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", mediaType)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}
