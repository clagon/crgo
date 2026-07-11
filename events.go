package crgo

import (
	"context"

	model "github.com/clagon/crgo/model"
)

/*
GetTrailEvents

	  開催中のイベント一覧を取得する。
	  params
		ctx: context.Context
	  return
		イベント一覧
		error
*/
func (c *Client) GetTrailEvents(ctx context.Context) (*model.TrailEventList, error) {

	// レスポンス用構造体
	var result model.TrailEventList

	// APIリクエストを送信
	if err := c.do(ctx, "/events", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
