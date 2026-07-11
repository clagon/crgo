package model

import (
	"encoding/json"
	"testing"
)

func TestClanLists(t *testing.T) {
	tests := []struct {
		name   string
		body   string
		decode func([]byte) (int, string, error)
	}{
		{
			name: "search_response",
			body: `{"items":[{"name":"Royals"}],"paging":{"cursors":{"after":"next"}}}`,
			decode: func(body []byte) (int, string, error) {
				var clans ClanList
				err := json.Unmarshal(body, &clans)
				return len(clans.Items), clans.Paging.Cursors.After, err
			},
		},
		{
			name: "members_response",
			body: `{"items":[{"name":"Member"}],"paging":{"cursors":{"after":"next"}}}`,
			decode: func(body []byte) (int, string, error) {
				var members ClanMemberPage
				err := json.Unmarshal(body, &members)
				return len(members.Items), members.Paging.Cursors.After, err
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
