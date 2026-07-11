package model

// LeaderboardList is a paginated leaderboard response.
type LeaderboardList struct {
	Items  []Leaderboard `json:"items,omitempty"`
	Paging Paging        `json:"paging,omitempty"`
}

type Leaderboard struct {
	Name JsonLocalizedName `json:"name,omitempty"`
	Id   int               `json:"id,omitempty"`
}
