package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetClanRanking(ctx context.Context, locationId string, options *PaginationOptions) (*ClanRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result ClanRankingList
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clans", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPlayerRanking(ctx context.Context, locationId string, options *PaginationOptions) (*PlayerRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result PlayerRankingList
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetClanWarsRanking(ctx context.Context, locationId string, options *PaginationOptions) (*ClanRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result ClanRankingList
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clanwars", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTopPlayerPathOfLegendSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*PlayerPathOfLegendRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result PlayerPathOfLegendRankingList
	if err := c.do(ctx, "/locations/global/pathoflegend/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTopPlayerLeagueSeason(ctx context.Context, seasonId string) (*LeagueSeason, error) {
	query := make(url.Values)
	var result LeagueSeason
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTopPlayerLeagueSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*PlayerRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result PlayerRankingList
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ListTopPlayerLeagueSeasons(ctx context.Context) (*LeagueSeasonList, error) {
	query := make(url.Values)
	var result LeagueSeasonList
	if err := c.do(ctx, "/locations/global/seasons", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetLocations(ctx context.Context, options *PaginationOptions) (*LocationList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result LocationList
	if err := c.do(ctx, "/locations", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ListTopPlayerLeagueSeasonsV2(ctx context.Context) (*LeagueSeasonList, error) {
	query := make(url.Values)
	var result LeagueSeasonList
	if err := c.do(ctx, "/locations/global/seasonsV2", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetLocation(ctx context.Context, locationId string) (*Location, error) {
	query := make(url.Values)
	var result Location
	if err := c.do(ctx, "/locations/"+escapedPath(locationId), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetGlobalTournamentRanking(ctx context.Context, tournamentTag string, options *PaginationOptions) (*LadderTournamentRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result LadderTournamentRankingList
	if err := c.do(ctx, "/locations/global/rankings/tournaments/"+escapedPath(tournamentTag), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPlayerPathOfLegendRanking(ctx context.Context, locationId string, options *PaginationOptions) (*PlayerPathOfLegendRankingList, error) {
	query := make(url.Values)
	if options != nil {
		addPagination(query, *options)
	}
	var result PlayerPathOfLegendRankingList
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/pathoflegend/players", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
