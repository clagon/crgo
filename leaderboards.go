package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetLeaderboard(ctx context.Context, leaderboardId int, options *PaginationOptions) (*LeaderboardList, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result LeaderboardList

	// APIリクエストを送信
	if err := c.do(ctx, "/leaderboard/"+escapedPath(leaderboardId), query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetLeaderboards(ctx context.Context) (*LeaderboardList, error) {

	// レスポンス用構造体
	var result LeaderboardList

	// APIリクエストを送信
	if err := c.do(ctx, "/leaderboards", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
