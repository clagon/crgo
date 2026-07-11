package crgo

import (
	"context"
	"net/url"
	"strconv"
)

func (c *Client) GetClanWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*ClanWarLog, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result ClanWarLog

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

func (c *Client) SearchClans(ctx context.Context, options *SearchClansOptions) (*ClanList, error) {

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
	var result ClanList

	// APIリクエストを送信
	if err := c.do(ctx, "/clans", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRiverRaceWarLog(ctx context.Context, clanTag string, options *PaginationOptions) (*RiverRaceLog, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result RiverRaceLog

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/riverracelog", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCurrentWar(ctx context.Context, clanTag string) (*CurrentClanWar, error) {

	// レスポンス用構造体
	var result CurrentClanWar

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentwar", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetClan(ctx context.Context, clanTag string) (*Clan, error) {

	// レスポンス用構造体
	var result Clan

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetClanMembers(ctx context.Context, clanTag string, options *PaginationOptions) (*ClanMemberList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result ClanMemberList

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/members", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetCurrentRiverRace(ctx context.Context, clanTag string) (*CurrentRiverRace, error) {

	// レスポンス用構造体
	var result CurrentRiverRace

	// APIリクエストを送信
	if err := c.do(ctx, "/clans/"+escapedPath(clanTag)+"/currentriverrace", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
