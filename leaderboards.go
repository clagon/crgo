package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetLeaderboard(ctx context.Context, leaderboardId int, options *PaginationOptions) (*LeaderboardList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result LeaderboardList
	if err := c.do(ctx, "/leaderboard/"+escapedPath(leaderboardId), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetLeaderboards(ctx context.Context) (*LeaderboardList, error) {
	query := make(url.Values)
	var result LeaderboardList
	if err := c.do(ctx, "/leaderboards", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
