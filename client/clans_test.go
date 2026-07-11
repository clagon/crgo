package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchClansOptions(t *testing.T) {

	tests := []struct {
		name    string
		options *SearchClansOptions
		want    map[string]string
	}{
		{
			name: "search_filters",
			options: &SearchClansOptions{
				Name: "Royals", LocationId: 57000000, MinMembers: 10, MaxMembers: 40, MinScore: 1234,
				Pagination: PaginationOptions{Limit: 20, After: "cursor"},
			},
			want: map[string]string{"name": "Royals", "locationId": "57000000", "minMembers": "10", "maxMembers": "40", "minScore": "1234", "limit": "20", "after": "cursor"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for key, want := range tt.want {
					if got := r.URL.Query().Get(key); got != want {
						t.Errorf("%s: got %q, want %q", key, got, want)
					}
				}
				_, _ = w.Write([]byte(`{"items":[]}`))
			}))
			defer server.Close()
			client, err := NewClient("secret", WithBaseURL(server.URL))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := client.SearchClans(context.Background(), tt.options); err != nil {
				t.Fatal(err)
			}
		})
	}
}
