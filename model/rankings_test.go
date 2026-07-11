package model

import (
	"encoding/json"
	"testing"
)

func TestClanRankingList(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"rank":1}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rankings ClanRankingList
			if err := json.Unmarshal([]byte(tt.body), &rankings); err != nil {
				t.Fatal(err)
			}
			if len(rankings.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(rankings.Items), tt.wantItems)
			}
			if rankings.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", rankings.Paging.Cursors.After)
			}
		})
	}
}

func TestRankingLists(t *testing.T) {
	tests := []struct {
		name   string
		body   string
		decode func([]byte) (int, string, error)
	}{
		{
			name: "player_rankings",
			body: `{"items":[{"rank":1}],"paging":{"cursors":{"after":"next"}}}`,
			decode: func(body []byte) (int, string, error) {
				var rankings PlayerRankingList
				err := json.Unmarshal(body, &rankings)
				return len(rankings.Items), rankings.Paging.Cursors.After, err
			},
		},
		{
			name: "path_of_legend_rankings",
			body: `{"items":[{"rank":1}],"paging":{"cursors":{"after":"next"}}}`,
			decode: func(body []byte) (int, string, error) {
				var rankings PlayerPathOfLegendRankingList
				err := json.Unmarshal(body, &rankings)
				return len(rankings.Items), rankings.Paging.Cursors.After, err
			},
		},
		{
			name: "league_seasons",
			body: `{"items":[{"id":"2020-10"}],"paging":{"cursors":{"after":"next"}}}`,
			decode: func(body []byte) (int, string, error) {
				var seasons LeagueSeasonList
				err := json.Unmarshal(body, &seasons)
				return len(seasons.Items), seasons.Paging.Cursors.After, err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, after, err := tt.decode([]byte(tt.body))
			if err != nil {
				t.Fatal(err)
			}
			if items != 1 || after != "next" {
				t.Fatalf("unexpected response: items=%d after=%q", items, after)
			}
		})
	}
}
