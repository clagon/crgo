package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
do

	APIリクエストを送信する。
	エンドポイントごとに異なる型のレスポンスを共通して処理できるよう、
	レスポンスは戻り値ではなく、デコード先のresultを引数として受け取る。
	成功レスポンスのJSONを、resultで指定されたデコード先へ格納する。
	params
		ctx: context.Context
		path: APIパス
		query: JSONのデコード先。json.Unmarshalが受け付けるポインタを指定する。
		result: デコード対象の構造体ポインタ
	return
		error
*/
func (c *Client) do(ctx context.Context, path string, query url.Values, result any) error {

	// -- リクエストURLを構築 --

	// baseURLをコピー
	u := *c.baseURL

	// パスを結合
	rawPath := strings.TrimRight(c.baseURL.EscapedPath(), "/") + "/" + strings.TrimLeft(path, "/")

	// パスをURLデコード
	decodedPath, err := url.PathUnescape(rawPath)
	if err != nil {
		// デコードできない不正な%文字が含まれている場合はエラー
		return fmt.Errorf("crgo: build request path: %w", err)
	}

	// パスを設定
	// エンコード済みのタグが二重エンコードされないよう、デコード済みのPathとエンコード済みのRawPathを設定
	u.Path = decodedPath
	u.RawPath = rawPath

	// クエリパラメータを設定
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("crgo: create request: %w", err)
	}

	// -- リクエストヘッダーを設定 --

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// -- リクエストを送信 --

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("crgo: send request: %w", err)
	}
	defer resp.Body.Close()

	// -- レスポンスを処理 --

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("crgo: read response: %w", err)
	}

	// ステータスコードが2xx以外の場合はエラー
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {

		apiErr := &APIError{StatusCode: resp.StatusCode, Body: body}

		// JSONでないエラーレスポンスでも、ステータスコードと本文を保持して返す
		_ = json.Unmarshal(body, apiErr)

		return apiErr
	}

	// デコード対象がない または レスポンスボディが空の場合は処理を終了
	if result == nil || len(body) == 0 {
		return nil
	}

	// レスポンスボディをJSONデコード
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("crgo: decode response: %w", err)
	}
	return nil
}

func escapedPath(value any) string { return url.PathEscape(fmt.Sprint(value)) }

func addPagination(query url.Values, options PaginationOptions) {
	if options.Limit != 0 {
		query.Set("limit", fmt.Sprint(options.Limit))
	}
	if options.After != "" {
		query.Set("after", options.After)
	}
	if options.Before != "" {
		query.Set("before", options.Before)
	}
}
