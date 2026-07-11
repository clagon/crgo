package crgo

import (
	"context"
	"net/url"
)

type SearchTournamentsOptions struct {
	Name       string
	Pagination PaginationOptions
}

func (c *Client) SearchTournaments(ctx context.Context, options *SearchTournamentsOptions) (*TournamentHeaderList, error) {
	query := make(url.Values)
	if options != nil {
		if options.Name != "" {
			query.Set("name", options.Name)
		}
		addPagination(query, options.Pagination)
	}
	var result TournamentHeaderList
	if err := c.do(ctx, "/tournaments", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTournament(ctx context.Context, tournamentTag string) (*Tournament, error) {
	query := make(url.Values)
	var result Tournament
	if err := c.do(ctx, "/tournaments/"+escapedPath(tournamentTag), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
