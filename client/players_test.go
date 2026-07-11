package client

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
	tests := []struct {
		name      string
		playerTag string
		wantURI   string
		wantTag   string
		wantName  string
		wantLevel int
	}{
		{
			name:      "escaped_tag",
			playerTag: "#ABC123",
			wantURI:   "/v1/players/%23ABC123",
			wantTag:   "#ABC123",
			wantName:  "Example",
			wantLevel: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.RequestURI != tt.wantURI {
					t.Errorf("URI: got %s, want %s", r.RequestURI, tt.wantURI)
				}
				if got := r.Header.Get("Authorization"); got != "Bearer secret" {
					t.Errorf("authorization: got %s", got)
				}
				if got := r.Header.Get("User-Agent"); got != "test-agent" {
					t.Errorf("user agent: got %s", got)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"tag":"#ABC123","name":"Example","expLevel":42}`))
			}))
			defer server.Close()
			client, err := NewClient("secret", WithBaseURL(server.URL+"/v1"), WithUserAgent("test-agent"))
			if err != nil {
				t.Fatal(err)
			}
			player, err := client.GetPlayer(context.Background(), tt.playerTag)
			if err != nil {
				t.Fatal(err)
			}
			if player.Tag != tt.wantTag || player.Name != tt.wantName || player.ExpLevel != tt.wantLevel {
				t.Fatalf("unexpected player: %+v", player)
			}
		})
	}
}

func TestPlayerFixture(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		wantClan string
	}{
		{name: "nested_clan", file: "testdata/player.json", wantClan: "Example Clan"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			var player model.Player
			if err := json.Unmarshal(body, &player); err != nil {
				t.Fatal(err)
			}
			if player.Clan.Name != tt.wantClan {
				t.Fatalf("clan name: got %q, want %q", player.Clan.Name, tt.wantClan)
			}
		})
	}
}
