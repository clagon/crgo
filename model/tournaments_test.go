package model

import (
	"encoding/json"
	"testing"
)

func TestLadderTournamentList(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"tag":"#TAG"}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tournaments LadderTournamentList
			if err := json.Unmarshal([]byte(tt.body), &tournaments); err != nil {
				t.Fatal(err)
			}
			if len(tournaments.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(tournaments.Items), tt.wantItems)
			}
			if tournaments.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", tournaments.Paging.Cursors.After)
			}
		})
	}
}

func TestTournamentHeaderList(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"tag":"#TAG"}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tournaments TournamentHeaderList
			if err := json.Unmarshal([]byte(tt.body), &tournaments); err != nil {
				t.Fatal(err)
			}
			if len(tournaments.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(tournaments.Items), tt.wantItems)
			}
			if tournaments.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", tournaments.Paging.Cursors.After)
			}
		})
	}
}
