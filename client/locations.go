package client

import (
	"context"
	"net/url"

	model "github.com/clagon/crgo/model"
)

/*
GetClanRanking

	  指定したロケーションのクランランキングを取得する。
	  params
		ctx: context.Context
		locationId: ロケーションID
		options: ページネーションオプション
	  return
		クランランキング
		error
*/
func (c *Client) GetClanRanking(ctx context.Context, locationId string, options *PaginationOptions) (*model.ClanRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.ClanRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clans", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetPlayerRanking

	  指定したロケーションのプレイヤーランキングを取得する。
	  params
		ctx: context.Context
		locationId: ロケーションID
		options: ページネーションオプション
	  return
		プレイヤーランキング
		error
*/
func (c *Client) GetPlayerRanking(ctx context.Context, locationId string, options *PaginationOptions) (*model.PlayerRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.PlayerRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetClanWarsRanking

	  指定したロケーションのクラン対戦ランキングを取得する。
	  params
		ctx: context.Context
		locationId: ロケーションID
		options: ページネーションオプション
	  return
		クラン対戦ランキング
		error
*/
func (c *Client) GetClanWarsRanking(ctx context.Context, locationId string, options *PaginationOptions) (*model.ClanRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.ClanRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/rankings/clanwars", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetTopPlayerPathOfLegendSeasonRankings

	  指定したシーズンのPoL上位プレイヤーを取得する。
	  params
		ctx: context.Context
		seasonId: シーズンID
		options: ページネーションオプション
	  return
		PoLランキング
		error
*/
func (c *Client) GetTopPlayerPathOfLegendSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*model.PlayerPathOfLegendRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.PlayerPathOfLegendRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/pathoflegend/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetTopPlayerLeagueSeason

	  指定したトッププレイヤーリーグのシーズン情報を取得する。
	  params
		ctx: context.Context
		seasonId: シーズンID
	  return
		リーグシーズン情報
		error
*/
func (c *Client) GetTopPlayerLeagueSeason(ctx context.Context, seasonId string) (*model.LeagueSeason, error) {

	// レスポンス用構造体
	var result model.LeagueSeason

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetTopPlayerLeagueSeasonRankings

	  指定したシーズンのトッププレイヤーランキングを取得する。
	  params
		ctx: context.Context
		seasonId: シーズンID
		options: ページネーションオプション
	  return
		プレイヤーランキング
		error
*/
func (c *Client) GetTopPlayerLeagueSeasonRankings(ctx context.Context, seasonId string, options *PaginationOptions) (*model.PlayerRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.PlayerRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons/"+escapedPath(seasonId)+"/rankings/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
ListTopPlayerLeagueSeasons

	  トッププレイヤーリーグのシーズン一覧を取得する。
	  params
		ctx: context.Context
	  return
		リーグシーズン一覧
		error
*/
func (c *Client) ListTopPlayerLeagueSeasons(ctx context.Context) (*model.LeagueSeasonList, error) {

	// レスポンス用構造体
	var result model.LeagueSeasonList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasons", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetLocations

	  ロケーション一覧を取得する。
	  params
		ctx: context.Context
		options: ページネーションオプション
	  return
		ロケーション一覧
		error
*/
func (c *Client) GetLocations(ctx context.Context, options *PaginationOptions) (*model.LocationList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.LocationList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
ListTopPlayerLeagueSeasonsV2

	  一意のシーズンIDと終了時刻を含むリーグシーズン一覧を取得する。
	  params
		ctx: context.Context
	  return
		リーグシーズン一覧
		error
*/
func (c *Client) ListTopPlayerLeagueSeasonsV2(ctx context.Context) (*model.LeagueSeasonList, error) {

	// レスポンス用構造体
	var result model.LeagueSeasonList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/seasonsV2", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetLocation

	  指定したロケーションの情報を取得する。
	  params
		ctx: context.Context
		locationId: ロケーションID
	  return
		ロケーション情報
		error
*/
func (c *Client) GetLocation(ctx context.Context, locationId string) (*model.Location, error) {

	// レスポンス用構造体
	var result model.Location

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetGlobalTournamentRanking

	  指定したグローバル大会のランキングを取得する。
	  params
		ctx: context.Context
		tournamentTag: トーナメントタグ
		options: ページネーションオプション
	  return
		トーナメントランキング
		error
*/
func (c *Client) GetGlobalTournamentRanking(ctx context.Context, tournamentTag string, options *PaginationOptions) (*model.LadderTournamentRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.LadderTournamentRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/global/rankings/tournaments/"+escapedPath(tournamentTag), query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetPlayerPathOfLegendRanking

	  指定したロケーションのPoLランキングを取得する。
	  params
		ctx: context.Context
		locationId: ロケーションID
		options: ページネーションオプション
	  return
		PoLのランキング
		error
*/
func (c *Client) GetPlayerPathOfLegendRanking(ctx context.Context, locationId string, options *PaginationOptions) (*model.PlayerPathOfLegendRankingList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.PlayerPathOfLegendRankingList

	// APIリクエストを送信
	if err := c.do(ctx, "/locations/"+escapedPath(locationId)+"/pathoflegend/players", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
