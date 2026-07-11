package model

import (
	"encoding/json"
	"testing"
)

func TestLeaderboardList(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"id":1}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var leaderboards LeaderboardList
			if err := json.Unmarshal([]byte(tt.body), &leaderboards); err != nil {
				t.Fatal(err)
			}
			if len(leaderboards.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(leaderboards.Items), tt.wantItems)
			}
			if leaderboards.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", leaderboards.Paging.Cursors.After)
			}
		})
	}
}
