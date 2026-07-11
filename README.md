# crgo

日本語 | [English](README.en.md)

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

> [!NOTE]
> APIトークンはClashRoyaleAPIから取得する必要があります。  
> https://developer.clashroyale.com

## モデル

```go
import "github.com/clagon/crgo/model"

var player *model.Player
```

APIレスポンスの型は `model` パッケージに定義されています。

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

任意引数を指定しない場合は `nil` を渡せます。

## APIエラー

```go
var apiErr *crgo.APIError
if errors.As(err, &apiErr) {
	fmt.Printf("status=%d reason=%s message=%s\n", apiErr.StatusCode, apiErr.Reason, apiErr.Message)
}
```

非2xxレスポンスは `*crgo.APIError` として返されます。  
429や5xxに対する自動リトライは行いません。

## クライアント設定

```go
client, err := crgo.NewClient(token,
	crgo.WithBaseURL("https://proxy.example.com/v1"),
	crgo.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	crgo.WithUserAgent("my-service/1.0"),
)
```

初期化時にAPIのBase URLなどを指定して差し替えることができます。

## 開発

```sh
go test ./...
go vet ./...
```

実APIへ接続するintegration testは、通常のテストでは実行されません。  
実行する場合は、有効なAPIトークンを設定して `integration` ビルドタグを指定します。

```sh
CLASH_ROYALE_API_TOKEN=... go test -tags=integration ./...
```
