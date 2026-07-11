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
	for _, status := range []int{400, 403, 404, 429, 500, 503} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"reason":"accessDenied","message":"no access","type":"client","detail":{"key":"value"}}`))
			}))
			defer server.Close()
			c, _ := NewClient("secret", WithBaseURL(server.URL))
			_, err := c.GetCards(context.Background(), nil)
			var apiErr *APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("expected APIError, got %v", err)
			}
			if apiErr.StatusCode != status || apiErr.Message != "no access" || len(apiErr.Body) == 0 {
				t.Fatalf("unexpected APIError: %+v", apiErr)
			}
		})
	}
}

func TestAPIErrorFixture(t *testing.T) {
	body, err := os.ReadFile("testdata/api_error.json")
	if err != nil {
		t.Fatal(err)
	}
	apiErr := &APIError{StatusCode: http.StatusForbidden, Body: body}
	if err := json.Unmarshal(body, apiErr); err != nil {
		t.Fatal(err)
	}
	if apiErr.Reason != "accessDenied" {
		t.Fatalf("unexpected APIError: %+v", apiErr)
	}
}

func TestNonJSONErrorAndInvalidSuccessJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "cards") {
			w.WriteHeader(502)
			_, _ = w.Write([]byte("bad gateway"))
			return
		}
		_, _ = w.Write([]byte(`{`))
	}))
	defer server.Close()
	c, _ := NewClient("secret", WithBaseURL(server.URL))
	_, err := c.GetCards(context.Background(), nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) || string(apiErr.Body) != "bad gateway" {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = c.GetPlayer(context.Background(), "tag")
	if err == nil || !strings.Contains(err.Error(), "decode response") {
		t.Fatalf("unexpected decode error: %v", err)
	}
}
