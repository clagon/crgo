package client

import (
	"context"
	"net/url"
	"strconv"

	model "github.com/clagon/crgo/model"
)

/*
GetClanWarLog

	  指定したクランのクラン対戦ログを取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
		options: ページネーションオプション
	  return
		クラン対戦ログ
		error
*/
func (c *Client) GetClanWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*model.ClanWarLog, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.ClanWarLog

	// APIリクエストを送信
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

/*
SearchClans

	  名前や各種条件でクランを検索する。
	  params
		ctx: context.Context
		options: 検索オプション
			name: クラン名の一部
			locationId: クランのロケーションID
			minMembers: クランのメンバー数の最小値
			maxMembers: クランのメンバー数の最大値
			minScore: クランスコアの最小値
			pagination: ページネーションオプション
	  return
		クラン一覧
		error
*/
func (c *Client) SearchClans(ctx context.Context, options *SearchClansOptions) (*model.ClanList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// 検索オプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
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

	// レスポンス用構造体
	var result model.ClanList

	// APIリクエストを送信
	if err := c.do(ctx, "/clans", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetRiverRaceWarLog

	  指定したクランのリバーレースログを取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
		options: ページネーションオプション
	  return
		リバーレースログ
		error
*/
func (c *Client) GetRiverRaceWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*model.RiverRaceLog, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.RiverRaceLog

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/riverracelog", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetCurrentWar

	  指定したクランの現在のクラン対戦情報を取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
	  return
		現在のクラン対戦情報
		error
*/
func (c *Client) GetCurrentWar(ctx context.Context, clanTag string) (*model.CurrentClanWar, error) {

	// レスポンス用構造体
	var result model.CurrentClanWar

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentwar", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetClan

	  指定したクランのクラン情報を取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
	  return
		クラン情報
		error
*/
func (c *Client) GetClan(ctx context.Context, clanTag string) (*model.Clan, error) {

	// レスポンス用構造体
	var result model.Clan

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetClanMembers

	  指定したクランのメンバー一覧を取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
		options: ページネーションオプション
	  return
		クランメンバー一覧
		error
*/
func (c *Client) GetClanMembers(ctx context.Context, clanTag string, options *PaginationOptions) (*model.ClanMemberPage, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.ClanMemberPage

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/members", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetCurrentRiverRace

	  指定したクランの現在のリバーレース情報を取得する。
	  params
		ctx: context.Context
		clanTag: クランタグ
	  return
		現在のリバーレース情報
		error
*/
func (c *Client) GetCurrentRiverRace(ctx context.Context, clanTag string) (*model.CurrentRiverRace, error) {

	// レスポンス用構造体
	var result model.CurrentRiverRace

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentriverrace", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
