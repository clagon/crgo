package client

import (
	"context"
	"net/url"

	model "github.com/clagon/crgo/model"
)

/*
GetLeaderboard

	  指定したリーダーボードのプレイヤー一覧を取得する。
	  params
		ctx: context.Context
		leaderboardId: リーダーボードID
		options: ページネーションオプション
	  return
		リーダーボード一覧
		error
*/
func (c *Client) GetLeaderboard(ctx context.Context, leaderboardId int, options *PaginationOptions) (*model.LeaderboardList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.LeaderboardList

	// APIリクエストを送信
	if err := c.do(ctx, "/leaderboard/"+escapedPath(leaderboardId), query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetLeaderboards

	  トロフィーロードごとのリーダーボード一覧を取得する。
	  params
		ctx: context.Context
	  return
		リーダーボード一覧
		error
*/
func (c *Client) GetLeaderboards(ctx context.Context) (*model.LeaderboardList, error) {

	// レスポンス用構造体
	var result model.LeaderboardList

	// APIリクエストを送信
	if err := c.do(ctx, "/leaderboards", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
