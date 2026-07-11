package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetTrailEvents(ctx context.Context) (*TrailEventList, error) {
	query := make(url.Values)
	var result TrailEventList
	if err := c.do(ctx, "/events", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
