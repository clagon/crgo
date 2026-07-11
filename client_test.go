package crgo

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
	if _, err := NewClient(""); err == nil {
		t.Fatal("expected empty token error")
	}
	if _, err := NewClient("token", WithBaseURL("relative")); err == nil {
		t.Fatal("expected invalid URL error")
	}
	if _, err := NewClient("token", WithHTTPClient(nil)); err == nil {
		t.Fatal("expected nil HTTP client error")
	}
}

func TestTransportErrorAndCancellation(t *testing.T) {
	transportErr := errors.New("transport failed")
	c, err := NewClient("token", WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, transportErr
	})}))
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.GetCards(context.Background(), nil)
	if !errors.Is(err, transportErr) {
		t.Fatalf("expected transport error, got %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c, _ = NewClient("token", WithHTTPClient(&http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, req.Context().Err()
	})}))
	_, err = c.GetCards(ctx, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
}

func TestEmptySuccessResponse(t *testing.T) {
	c, _ := NewClient("token", WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusNoContent, Body: io.NopCloser(strings.NewReader("")), Header: make(http.Header)}, nil
	})}))
	items, err := c.GetCards(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if items == nil {
		t.Fatal("expected an allocated response value")
	}
}

func TestClientOptions(t *testing.T) {
	hc := &http.Client{}
	c, err := NewClient("token", WithBaseURL("https://example.com/v1/"), WithHTTPClient(hc), WithUserAgent("custom-agent"))
	if err != nil {
		t.Fatal(err)
	}
	if c.baseURL.String() != "https://example.com/v1" {
		t.Fatalf("unexpected base URL: %s", c.baseURL)
	}
	if c.httpClient != hc || c.userAgent != "custom-agent" {
		t.Fatal("options were not applied")
	}
}
