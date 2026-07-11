package model

import (
	"encoding/json"
	"testing"
)

func TestLocationList(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantItems int
	}{
		{
			name:      "paginated_response",
			body:      `{"items":[{"id":57000000}],"paging":{"cursors":{"after":"next"}}}`,
			wantItems: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var locations LocationList
			if err := json.Unmarshal([]byte(tt.body), &locations); err != nil {
				t.Fatal(err)
			}
			if len(locations.Items) != tt.wantItems {
				t.Fatalf("items: got %d, want %d", len(locations.Items), tt.wantItems)
			}
			if locations.Paging.Cursors.After != "next" {
				t.Fatalf("after cursor: got %q", locations.Paging.Cursors.After)
			}
		})
	}
}
