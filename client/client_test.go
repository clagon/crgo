package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func TestNewClientValidation(t *testing.T) {
	tests := []struct {
		name  string
		token string
		opts  []ClientOption
	}{
		{name: "empty_token", token: ""},
		{name: "invalid_base_url", token: "token", opts: []ClientOption{WithBaseURL("relative")}},
		{name: "nil_http_client", token: "token", opts: []ClientOption{WithHTTPClient(nil)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewClient(tt.token, tt.opts...); err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

func TestTransportErrorAndCancellation(t *testing.T) {
	transportErr := errors.New("transport failed")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	tests := []struct {
		name      string
		ctx       context.Context
		transport roundTripFunc
		wantErr   error
	}{
		{
			name: "transport_error",
			ctx:  context.Background(),
			transport: func(*http.Request) (*http.Response, error) {
				return nil, transportErr
			},
			wantErr: transportErr,
		},
		{
			name: "canceled_context",
			ctx:  ctx,
			transport: func(req *http.Request) (*http.Response, error) {
				return nil, req.Context().Err()
			},
			wantErr: context.Canceled,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient("token", WithHTTPClient(&http.Client{Transport: tt.transport}))
			if err != nil {
				t.Fatal(err)
			}
			_, err = client.GetCards(tt.ctx, nil)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEmptySuccessResponse(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
	}{
		{name: "no_content", statusCode: http.StatusNoContent, body: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient("token", WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: tt.statusCode, Body: io.NopCloser(strings.NewReader(tt.body)), Header: make(http.Header)}, nil
			})}))
			if err != nil {
				t.Fatal(err)
			}
			items, err := client.GetCards(context.Background(), nil)
			if err != nil {
				t.Fatal(err)
			}
			if items == nil {
				t.Fatal("expected an allocated response value")
			}
		})
	}
}

func TestClientOptions(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		userAgent   string
		wantBaseURL string
	}{
		{
			name:        "custom_options",
			baseURL:     "https://example.com/v1/",
			userAgent:   "custom-agent",
			wantBaseURL: "https://example.com/v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := &http.Client{}
			client, err := NewClient("token", WithBaseURL(tt.baseURL), WithHTTPClient(httpClient), WithUserAgent(tt.userAgent))
			if err != nil {
				t.Fatal(err)
			}
			if client.baseURL.String() != tt.wantBaseURL {
				t.Fatalf("base URL: got %s, want %s", client.baseURL, tt.wantBaseURL)
			}
			if client.httpClient != httpClient {
				t.Fatal("HTTP client was not applied")
			}
			if client.userAgent != tt.userAgent {
				t.Fatalf("user agent: got %q, want %q", client.userAgent, tt.userAgent)
			}
		})
	}
}
