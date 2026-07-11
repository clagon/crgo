package crgo_test

import (
	"testing"

	crgo "github.com/clagon/crgo"
)

func TestSearchOptions(t *testing.T) {
	tests := []struct {
		name  string
		check func(t *testing.T)
	}{
		{
			name: "clans",
			check: func(t *testing.T) {
				options := crgo.SearchClansOptions{
					Name:       "Royals",
					MinMembers: 10,
					Pagination: crgo.PaginationOptions{Limit: 20},
				}
				if options.Name != "Royals" || options.MinMembers != 10 || options.Pagination.Limit != 20 {
					t.Fatalf("unexpected options: %+v", options)
				}
			},
		},
		{
			name: "tournaments",
			check: func(t *testing.T) {
				options := crgo.SearchTournamentsOptions{
					Name:       "Global",
					Pagination: crgo.PaginationOptions{Limit: 10},
				}
				if options.Name != "Global" || options.Pagination.Limit != 10 {
					t.Fatalf("unexpected options: %+v", options)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.check)
	}
}
