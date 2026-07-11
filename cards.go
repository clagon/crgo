package crgo

import (
	"context"
	"net/url"

	model "github.com/clagon/crgo/model"
)

/*
GetCards

	  カード一覧を取得する。
	  params
		ctx: context.Context
		options: ページネーションオプション
	  return
		カード一覧
		error
*/
func (c *Client) GetCards(ctx context.Context, options *PaginationOptions) (*model.Items, error) {

	// クエリパラメータ初期化
	query := make(url.Values)

	// ページネーションオプションが指定されている場合
	if options != nil {
		// クエリパラメータに追加
		addPagination(query, *options)
	}

	// レスポンス用構造体
	var result model.Items

	// APIリクエストを送信
	if err := c.do(ctx, "/cards", query, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
