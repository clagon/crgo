package crgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) do(ctx context.Context, path string, query url.Values, result any) error {
	u := *c.baseURL
	rawPath := strings.TrimRight(c.baseURL.EscapedPath(), "/") + "/" + strings.TrimLeft(path, "/")
	decodedPath, err := url.PathUnescape(rawPath)
	if err != nil {
		return fmt.Errorf("crgo: build request path: %w", err)
	}
	u.Path = decodedPath
	u.RawPath = rawPath
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("crgo: create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("crgo: send request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("crgo: read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &APIError{StatusCode: resp.StatusCode, Body: body}
		_ = json.Unmarshal(body, apiErr)
		return apiErr
	}
	if result == nil || len(body) == 0 {
		return nil
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("crgo: decode response: %w", err)
	}
	return nil
}

func escapedPath(value any) string { return url.PathEscape(fmt.Sprint(value)) }

func addPagination(query url.Values, options PaginationOptions) {
	if options.Limit != 0 {
		query.Set("limit", fmt.Sprint(options.Limit))
	}
	if options.After != "" {
		query.Set("after", options.After)
	}
	if options.Before != "" {
		query.Set("before", options.Before)
	}
}
