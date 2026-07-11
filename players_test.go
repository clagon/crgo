package crgo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	model "github.com/clagon/crgo/model"
)

func TestGetPlayer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/v1/players/%23ABC123" {
			t.Errorf("unexpected URI: %s", r.RequestURI)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer secret" {
			t.Errorf("unexpected authorization: %s", got)
		}
		if got := r.Header.Get("User-Agent"); got != "test-agent" {
			t.Errorf("unexpected user agent: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tag":"#ABC123","name":"Example","expLevel":42}`))
	}))
	defer server.Close()
	c, err := NewClient("secret", WithBaseURL(server.URL+"/v1"), WithUserAgent("test-agent"))
	if err != nil {
		t.Fatal(err)
	}
	player, err := c.GetPlayer(context.Background(), "#ABC123")
	if err != nil {
		t.Fatal(err)
	}
	if player.Tag != "#ABC123" || player.Name != "Example" || player.ExpLevel != 42 {
		t.Fatalf("unexpected player: %+v", player)
	}
}

func TestPlayerFixture(t *testing.T) {
	body, err := os.ReadFile("testdata/player.json")
	if err != nil {
		t.Fatal(err)
	}
	var player model.Player
	if err := json.Unmarshal(body, &player); err != nil {
		t.Fatal(err)
	}
	if player.Clan.Name != "Example Clan" {
		t.Fatalf("unexpected nested clan: %+v", player.Clan)
	}
}
