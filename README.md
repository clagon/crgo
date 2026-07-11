# crgo
Clash Royale APIのGoクライアントです。

## インストール

```sh
go get github.com/clagon/crgo
```

## 基本的な使い方

```go
client, err := crgo.NewClient(os.Getenv("CLASH_ROYALE_API_TOKEN"))
if err != nil {
	log.Fatal(err)
}

player, err := client.GetPlayer(context.Background(), "#PLAYER_TAG")
if err != nil {
	log.Fatal(err)
}
fmt.Println(player.Name)
```

トークンはすべてのリクエストへ `Authorization: Bearer ...` として設定されます。ライブラリがトークンをログやエラーメッセージへ出力することはありません。

## モデル

APIレスポンスの型は `model` パッケージで提供します。APIメソッドの戻り値は、たとえば `*model.Player` です。

```go
import "github.com/clagon/crgo/model"

var player *model.Player
```

## 検索とページング

```go
clans, err := client.SearchClans(ctx, &crgo.SearchClansOptions{
	Name:       "Royals",
	MinMembers: 10,
	Pagination: crgo.PaginationOptions{Limit: 20},
})

players, err := client.GetPlayerRanking(ctx, "global", &crgo.PaginationOptions{
	Limit: 50,
	After: "cursor",
})
```

任意引数を指定しない場合は `nil` を渡せます。APIの仕様上、`After` と `Before` は同時に指定しないでください。

## APIエラー

```go
var apiErr *crgo.APIError
if errors.As(err, &apiErr) {
	fmt.Printf("status=%d reason=%s message=%s\n", apiErr.StatusCode, apiErr.Reason, apiErr.Message)
}
```

非2xxレスポンスは `*crgo.APIError` として返されます。JSONとして解釈できないレスポンスを含め、生の本文は `Body` に保持されます。429や5xxに対する自動リトライは行いません。

## クライアント設定

```go
client, err := crgo.NewClient(token,
	crgo.WithBaseURL("https://proxy.example.com/v1"),
	crgo.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	crgo.WithUserAgent("my-service/1.0"),
)
```

Base URLとHTTPクライアントを差し替えられるため、プロキシやテスト環境でも利用できます。

## 開発

```sh
go test ./...
go vet ./...
```

実APIへ接続するintegration testは、通常のテストやCIでは実行されません。実行する場合は、有効なAPIトークンを設定して `integration` ビルドタグを指定します。

```sh
CLASH_ROYALE_API_TOKEN=... go test -tags=integration ./...
```

integration testはClash Royale公式APIの `/cards` を呼び出します。トークンが未設定の場合は設定漏れとしてテストが失敗します。トークンはリポジトリやテストfixtureへ保存しません。

API仕様の基準はリポジトリ直下の `swagger.json` です。
