package crgo

import (
	"context"
)

/*
GetPlayer

	  指定したプレイヤーのプレイヤー情報を取得する。
	  params
		ctx: context.Context
		playerTag: プレイヤータグ
	  return
		プレイヤー情報
		error
*/
func (c *Client) GetPlayer(ctx context.Context, playerTag string) (*Player, error) {

	// レスポンス用構造体
	var result Player

	// APIリクエストを送信
	if err := c.do(ctx, "/players/"+escapedPath(playerTag), nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetPlayerUpcomingChests

	  指定したプレイヤーの宝箱情報を取得する。
	  params
		ctx: context.Context
		playerTag: プレイヤータグ
	  return
		宝箱情報
		error
*/
func (c *Client) GetPlayerUpcomingChests(ctx context.Context, playerTag string) (*UpcomingChests, error) {

	// レスポンス用構造体
	var result UpcomingChests

	// APIリクエストを送信
	if err := c.do(ctx, "/players/"+escapedPath(playerTag)+"/upcomingchests", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
GetPlayerBattles

	  指定したプレイヤーのバトルログ情報を取得する。
	  params
		ctx: context.Context
		playerTag: プレイヤータグ
	  return
		バトルログ情報
		error
*/
func (c *Client) GetPlayerBattles(ctx context.Context, playerTag string) (*BattleList, error) {

	// レスポンス用構造体
	var result BattleList

	// APIリクエストを送信
	if err := c.do(ctx, "/players/"+escapedPath(playerTag)+"/battlelog", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
