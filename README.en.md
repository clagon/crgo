# crgo

[日本語](README.md) | English

A Go client for the Clash Royale API.

## Installation

```sh
go get github.com/clagon/crgo
```

## Basic Usage

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
> You need to obtain an API token from the Clash Royale API.  
> https://developer.clashroyale.com

## Models

```go
import "github.com/clagon/crgo/model"

var player *model.Player
```

API response types are defined in the `model` package.

Response model JSON tags do not use `omitempty`, so values explicitly returned by the API, including `0`, `false`, and empty strings, are preserved when the model is marshaled back to JSON.   
Values whose presence is meaningful use pointers. For example, a nil `Item.ElixirCost` means the API response did not contain an elixir cost; otherwise, the pointed-to value is the cost.

## Search and Pagination

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

Pass `nil` when no optional parameters are needed.

## API Errors

```go
var apiErr *crgo.APIError
if errors.As(err, &apiErr) {
	fmt.Printf("status=%d reason=%s message=%s\n", apiErr.StatusCode, apiErr.Reason, apiErr.Message)
}
```

Non-2xx responses are returned as `*crgo.APIError`.  
The client does not automatically retry 429 or 5xx responses.

## Client Configuration

```go
client, err := crgo.NewClient(token,
	crgo.WithBaseURL("https://proxy.example.com/v1"),
	crgo.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	crgo.WithUserAgent("my-service/1.0"),
)
```

You can override settings such as the API base URL when initializing the client.

## Development

```sh
go test ./...
go vet ./...
```

Integration tests that connect to the live API are not run as part of the regular test suite.  
To run them, set a valid API token and specify the `integration` build tag.

```sh
CLASH_ROYALE_API_TOKEN=... go test -tags=integration ./...
```
