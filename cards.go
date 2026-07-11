package crgo

import (
	"context"
	"net/url"
)

func (c *Client) GetCards(ctx context.Context, options *PaginationOptions) (*Items, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result Items

	// APIリクエストを送信
	if err := c.do(ctx, "/cards", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
