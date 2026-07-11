package crgo

import (
	"context"
)

func (c *Client) GetGlobalTournaments(ctx context.Context) (*LadderTournamentList, error) {

	// レスポンス用構造体
	var result LadderTournamentList

	// APIリクエストを送信
	if err := c.do(ctx, "/globaltournaments", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
