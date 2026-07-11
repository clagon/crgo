package client

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// ClientOption configures a Client.
type ClientOption func(*Client) error

/*
WithBaseURL

	  APIのベースURLを指定する。
	  params
		rawURL: APIのベースURL
	  return
		ClientOption
*/
func WithBaseURL(rawURL string) ClientOption {
	return func(c *Client) error {
		u, err := url.Parse(rawURL)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return errors.New("crgo: base URL must be an absolute HTTP URL")
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.New("crgo: base URL must use http or https")
		}
		u.RawQuery, u.Fragment = "", ""
		u.Path = strings.TrimRight(u.Path, "/")
		c.baseURL = u
		return nil
	}
}

/*
WithHTTPClient

	  HTTPクライアントを指定する。
	  params
		client: HTTPクライアント
	  return
		ClientOption
*/
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		if client == nil {
			return errors.New("crgo: HTTP client must not be nil")
		}
		c.httpClient = client
		return nil
	}
}

/*
WithUserAgent

	  User-Agentを指定する。
	  params
		userAgent: User-Agent
	  return
		ClientOption
*/
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error { c.userAgent = strings.TrimSpace(userAgent); return nil }
}

// PaginationOptions contains cursor pagination parameters.
type PaginationOptions struct {
	Limit  int
	After  string
	Before string
}
