package crgo

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const defaultBaseURL = "https://api.clashroyale.com/v1"

// Client is a Clash Royale API client.
type Client struct {
	token      string
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string
}

// NewClient constructs a client authenticated with token.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	if strings.TrimSpace(token) == "" {
		return nil, errors.New("crgo: token must not be empty")
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{token: token, baseURL: baseURL, httpClient: http.DefaultClient, userAgent: "go-cr-api/1.0"}
	for _, opt := range opts {
		if opt == nil {
			return nil, errors.New("crgo: client option must not be nil")
		}
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	if c.httpClient == nil {
		return nil, errors.New("crgo: HTTP client must not be nil")
	}
	return c, nil
}
