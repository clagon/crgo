package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetCards(ctx context.Context, options *PaginationOptions) (*Items, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result Items
	if err := c.do(ctx, "/cards", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
