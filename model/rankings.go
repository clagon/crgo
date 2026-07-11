package model

// PlayerPathOfLegendRankingList is a paginated Path of Legend ranking response.
type PlayerPathOfLegendRankingList struct {
	Items  []PlayerPathOfLegendRanking `json:"items"`
	Paging Paging                      `json:"paging"`
}

type PlayerPathOfLegendRanking struct {
	Clan      PlayerRankingClan `json:"clan"`
	Tag       string            `json:"tag"`
	Name      string            `json:"name"`
	ExpLevel  int               `json:"expLevel"`
	Rank      int               `json:"rank"`
	EloRating int               `json:"eloRating"`
}

type PlayerRankingClan struct {
	BadgeId   int            `json:"badgeId"`
	Tag       string         `json:"tag"`
	Name      string         `json:"name"`
	BadgeUrls map[string]any `json:"badgeUrls"`
}

// LeagueSeasonList is a paginated league season response.
type LeagueSeasonList struct {
	Items  []LeagueSeason `json:"items"`
	Paging Paging         `json:"paging"`
}

type LeagueSeason struct {
	Id string `json:"id"`
}

// PlayerRankingList is a paginated player ranking response.
type PlayerRankingList struct {
	Items  []PlayerRanking `json:"items"`
	Paging Paging          `json:"paging"`
}

type PlayerRanking struct {
	Clan         PlayerRankingClan `json:"clan"`
	Arena        Arena             `json:"arena"`
	Tag          string            `json:"tag"`
	Name         string            `json:"name"`
	ExpLevel     int               `json:"expLevel"`
	Rank         int               `json:"rank"`
	PreviousRank int               `json:"previousRank"`
	Trophies     int               `json:"trophies"`
}

// ClanRankingList is a paginated clan ranking response.
type ClanRankingList struct {
	Items  []ClanRanking `json:"items"`
	Paging Paging        `json:"paging"`
}

type ClanRanking struct {
	ClanScore    int            `json:"clanScore"`
	BadgeId      int            `json:"badgeId"`
	Location     Location       `json:"location"`
	Members      int            `json:"members"`
	Tag          string         `json:"tag"`
	Name         string         `json:"name"`
	Rank         int            `json:"rank"`
	PreviousRank int            `json:"previousRank"`
	BadgeUrls    map[string]any `json:"badgeUrls"`
}

type PathOfLegendSeasonResult struct {
	Trophies     int `json:"trophies"`
	LeagueNumber int `json:"leagueNumber"`
	Rank         int `json:"rank"`
}

type PlayerLeagueStatistics struct {
	PreviousSeason LeagueSeasonResult `json:"previousSeason"`
	CurrentSeason  LeagueSeasonResult `json:"currentSeason"`
	BestSeason     LeagueSeasonResult `json:"bestSeason"`
}

type LeagueSeasonResult struct {
	Trophies     int    `json:"trophies"`
	Rank         int    `json:"rank"`
	BestTrophies int    `json:"bestTrophies"`
	Id           string `json:"id"`
}
