package crgo

import (
	"context"
	"net/url"
)

type SearchTournamentsOptions struct {
	Name       string
	Pagination PaginationOptions
}

/*
SearchTournaments

	  名前で大会を検索する。
	  params
		ctx: context.Context
		options: 検索オプション
	  return
		大会一覧
		error
*/
func (c *Client) SearchTournaments(ctx context.Context, options *SearchTournamentsOptions) (*TournamentHeaderList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// 検索オプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		if options.Name != "" {
			query.Set("name", options.Name)
		}
		addPagination(query, options.Pagination)
	}

	// レスポンス用構造体
	var result TournamentHeaderList

	// APIリクエストを送信
	if err := c.do(ctx, "/tournaments", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetTournament

	  指定した大会の情報を取得する。
	  params
		ctx: context.Context
		tournamentTag: 大会タグ
	  return
		大会情報
		error
*/
func (c *Client) GetTournament(ctx context.Context, tournamentTag string) (*Tournament, error) {

	// レスポンス用構造体
	var result Tournament

	// APIリクエストを送信
	if err := c.do(ctx, "/tournaments/"+escapedPath(tournamentTag), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
