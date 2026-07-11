package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetGlobalTournaments(ctx context.Context) (*LadderTournamentList, error) {
	query := make(url.Values)
	var result LadderTournamentList
	if err := c.do(ctx, "/globaltournaments", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
