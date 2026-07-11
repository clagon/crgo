package client

import "fmt"

// APIError represents a non-2xx response from the Clash Royale API.
type APIError struct {
	StatusCode int            `json:"-"`
	Reason     string         `json:"reason,omitempty"`
	Message    string         `json:"message,omitempty"`
	Type       string         `json:"type,omitempty"`
	Detail     map[string]any `json:"detail,omitempty"`
	Body       []byte         `json:"-"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("crgo: API error (status %d): %s", e.StatusCode, e.Message)
	}
	if e.Reason != "" {
		return fmt.Sprintf("crgo: API error (status %d): %s", e.StatusCode, e.Reason)
	}
	return fmt.Sprintf("crgo: API error (status %d)", e.StatusCode)
}
