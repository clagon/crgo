package crgo

import (
	"context"
	"net/url"
)

type SearchTournamentsOptions struct {
	Name       string
	Pagination PaginationOptions
}

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

func (c *Client) GetTournament(ctx context.Context, tournamentTag string) (*Tournament, error) {

	// レスポンス用構造体
	var result Tournament

	// APIリクエストを送信
	if err := c.do(ctx, "/tournaments/"+escapedPath(tournamentTag), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
