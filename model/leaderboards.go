package model

type LeaderboardList []Leaderboard

type Leaderboard struct {
	Name JsonLocalizedName `json:"name,omitempty"`
	Id   int               `json:"id,omitempty"`
}
