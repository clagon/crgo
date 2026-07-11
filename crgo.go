// Package crgo provides a typed client for the Clash Royale API.
package crgo

import (
	"net/http"

	"github.com/clagon/crgo/client"
)

// Client is a Clash Royale API client.
type Client = client.Client

// ClientOption configures a Client.
type ClientOption = client.ClientOption

// PaginationOptions contains cursor pagination parameters.
type PaginationOptions = client.PaginationOptions

// SearchClansOptions contains clan search and pagination parameters.
type SearchClansOptions = client.SearchClansOptions

// SearchTournamentsOptions contains tournament search and pagination parameters.
type SearchTournamentsOptions = client.SearchTournamentsOptions

// APIError represents a non-2xx response from the Clash Royale API.
type APIError = client.APIError

// NewClient constructs a client authenticated with token.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	return client.NewClient(token, opts...)
}

// WithBaseURL sets the API base URL.
func WithBaseURL(rawURL string) ClientOption { return client.WithBaseURL(rawURL) }

// WithHTTPClient sets the HTTP client used for API requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return client.WithHTTPClient(httpClient)
}

// WithUserAgent sets the User-Agent request header.
func WithUserAgent(userAgent string) ClientOption { return client.WithUserAgent(userAgent) }
