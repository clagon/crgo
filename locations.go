package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetClanRanking(ctx context.Context, locationId string, options *PaginationOptions) (*ClanRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result ClanRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clans", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPlayerRanking(ctx context.Context, locationId string, options *PaginationOptions) (*PlayerRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result PlayerRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetClanWarsRanking(ctx context.Context, locationId string, options *PaginationOptions) (*ClanRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result ClanRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clanwars", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTopPlayerPathOfLegendSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*PlayerPathOfLegendRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result PlayerPathOfLegendRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/pathoflegend/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTopPlayerLeagueSeason(ctx context.Context, seasonId string) (*LeagueSeason, error) {

	// レスポンス用構造体
	var result LeagueSeason

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTopPlayerLeagueSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*PlayerRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result PlayerRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListTopPlayerLeagueSeasons(ctx context.Context) (*LeagueSeasonList, error) {

	// レスポンス用構造体
	var result LeagueSeasonList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetLocations(ctx context.Context, options *PaginationOptions) (*LocationList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result LocationList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ListTopPlayerLeagueSeasonsV2(ctx context.Context) (*LeagueSeasonList, error) {

	// レスポンス用構造体
	var result LeagueSeasonList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasonsV2", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetLocation(ctx context.Context, locationId string) (*Location, error) {

	// レスポンス用構造体
	var result Location

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetGlobalTournamentRanking(ctx context.Context, tournamentTag string, options *PaginationOptions) (*LadderTournamentRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result LadderTournamentRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/rankings/tournaments/"+escapedPath(tournamentTag), query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPlayerPathOfLegendRanking(ctx context.Context, locationId string, options *PaginationOptions) (*PlayerPathOfLegendRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result PlayerPathOfLegendRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/pathoflegend/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
