package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetPlayer(ctx context.Context, playerTag string) (*Player, error) {
	query := make(url.Values)
	var result Player
	if err := c.do(ctx, "/players/"+escapedPath(playerTag), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPlayerUpcomingChests(ctx context.Context, playerTag string) (*UpcomingChests, error) {
	query := make(url.Values)
	var result UpcomingChests
	if err := c.do(ctx, "/players/"+escapedPath(playerTag)+"/upcomingchests", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPlayerBattles(ctx context.Context, playerTag string) (*BattleList, error) {
	query := make(url.Values)
	var result BattleList
	if err := c.do(ctx, "/players/"+escapedPath(playerTag)+"/battlelog", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
