package model

import (
	"encoding/json"
	"testing"
)

func TestRiverRaceLog(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"seasonId":1}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var log RiverRaceLog
			if err := json.Unmarshal([]byte(tt.body), &log); err != nil {
				t.Fatal(err)
			}
			if len(log.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(log.Items), tt.wantItems)
			}
			if log.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", log.Paging.Cursors.After)
			}
		})
	}
}
