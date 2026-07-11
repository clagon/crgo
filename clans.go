package crgo

import (
	"context"
	"net/url"
	"strconv"
)

func (c *Client) GetClanWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*ClanWarLog, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result ClanWarLog
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/warlog", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type SearchClansOptions struct {
	Name       string
	LocationId int
	MinMembers int
	MaxMembers int
	MinScore   int
	Pagination PaginationOptions
}

func (c *Client) SearchClans(ctx context.Context, options *SearchClansOptions) (*ClanList, error) {
	query := make(url.Values)
	if options != nil {
		if options.Name != "" {
			query.Set("name", options.Name)
		}
		if options.LocationId != 0 {
			query.Set("locationId", strconv.Itoa(options.LocationId))
		}
		if options.MinMembers != 0 {
			query.Set("minMembers", strconv.Itoa(options.MinMembers))
		}
		if options.MaxMembers != 0 {
			query.Set("maxMembers", strconv.Itoa(options.MaxMembers))
		}
		if options.MinScore != 0 {
			query.Set("minScore", strconv.Itoa(options.MinScore))
		}
		addPagination(query, options.Pagination)
	}
	var result ClanList
	if err := c.do(ctx, "/clans", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetRiverRaceWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*RiverRaceLog, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result RiverRaceLog
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/riverracelog", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetCurrentWar(ctx context.Context, clanTag string) (*CurrentClanWar, error) {
	query := make(url.Values)
	var result CurrentClanWar
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentwar", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetClan(ctx context.Context, clanTag string) (*Clan, error) {
	query := make(url.Values)
	var result Clan
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetClanMembers(ctx context.Context, clanTag string, options *PaginationOptions) (*ClanMemberList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result ClanMemberList
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/members", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetCurrentRiverRace(ctx context.Context, clanTag string) (*CurrentRiverRace, error) {
	query := make(url.Values)
	var result CurrentRiverRace
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentriverrace", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
