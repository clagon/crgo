package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAPIError(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{name: "bad_request", status: http.StatusBadRequest},
		{name: "forbidden", status: http.StatusForbidden},
		{name: "not_found", status: http.StatusNotFound},
		{name: "too_many_requests", status: http.StatusTooManyRequests},
		{name: "internal_server_error", status: http.StatusInternalServerError},
		{name: "service_unavailable", status: http.StatusServiceUnavailable},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(`{"reason":"accessDenied","message":"no access","type":"client","detail":{"key":"value"}}`))
			}))
			defer server.Close()
			c, _ := NewClient("secret", WithBaseURL(server.URL))
			_, err := c.GetCards(context.Background(), nil)
			var apiErr *APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("expected APIError, got %v", err)
			}
			if apiErr.StatusCode != tt.status || apiErr.Message != "no access" || len(apiErr.Body) == 0 {
				t.Fatalf("unexpected APIError: %+v", apiErr)
			}
		})
	}
}

func TestAPIErrorFixture(t *testing.T) {
	tests := []struct {
		name       string
		file       string
		wantReason string
	}{
		{name: "access_denied", file: "testdata/api_error.json", wantReason: "accessDenied"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			apiErr := &APIError{StatusCode: http.StatusForbidden, Body: body}
			if err := json.Unmarshal(body, apiErr); err != nil {
				t.Fatal(err)
			}
			if apiErr.Reason != tt.wantReason {
				t.Fatalf("reason: got %q, want %q", apiErr.Reason, tt.wantReason)
			}
		})
	}
}

func TestNonJSONErrorAndInvalidSuccessJSON(t *testing.T) {
	tests := []struct {
		name         string
		status       int
		body         string
		call         func(context.Context, *Client) error
		wantAPIError bool
		wantContains string
	}{
		{
			name:   "non_json_error",
			status: http.StatusBadGateway,
			body:   "bad gateway",
			call: func(ctx context.Context, client *Client) error {
				_, err := client.GetCards(ctx, nil)
				return err
			},
			wantAPIError: true,
		},
		{
			name: "invalid_json",
			body: `{`,
			call: func(ctx context.Context, client *Client) error {
				_, err := client.GetPlayer(ctx, "tag")
				return err
			},
			wantContains: "decode response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()
			client, err := NewClient("secret", WithBaseURL(server.URL))
			if err != nil {
				t.Fatal(err)
			}
			err = tt.call(context.Background(), client)
			if err == nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.wantContains != "" && !strings.Contains(err.Error(), tt.wantContains) {
				t.Fatalf("error: got %q, want substring %q", err, tt.wantContains)
			}
			if !tt.wantAPIError {
				return
			}
			var apiErr *APIError
			if !errors.As(err, &apiErr) || string(apiErr.Body) != tt.body {
				t.Fatalf("unexpected APIError: %v", err)
			}
		})
	}
}
