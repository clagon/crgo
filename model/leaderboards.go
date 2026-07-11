package model

// LeaderboardList is a paginated leaderboard response.
type LeaderboardList struct {
	Items  []Leaderboard `json:"items"`
	Paging Paging        `json:"paging"`
}

type Leaderboard struct {
	Name JsonLocalizedName `json:"name"`
	Id   int               `json:"id"`
}
