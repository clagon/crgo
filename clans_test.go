package crgo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchClansOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		for key, want := range map[string]string{"name": "Royals", "locationId": "57000000", "minMembers": "10", "maxMembers": "40", "minScore": "1234", "limit": "20", "after": "cursor"} {
			if got := q.Get(key); got != want {
				t.Errorf("%s: got %q want %q", key, got, want)
			}
		}
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()
	c, _ := NewClient("secret", WithBaseURL(server.URL))
	_, err := c.SearchClans(context.Background(), &SearchClansOptions{Name: "Royals", LocationId: 57000000, MinMembers: 10, MaxMembers: 40, MinScore: 1234, Pagination: PaginationOptions{Limit: 20, After: "cursor"}})
	if err != nil {
		t.Fatal(err)
	}
}
