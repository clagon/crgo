package model

// PlayerPathOfLegendRankingList is a paginated Path of Legend ranking response.
type PlayerPathOfLegendRankingList struct {
	Items  []PlayerPathOfLegendRanking `json:"items,omitempty"`
	Paging Paging                      `json:"paging,omitempty"`
}

type PlayerPathOfLegendRanking struct {
	Clan      PlayerRankingClan `json:"clan,omitempty"`
	Tag       string            `json:"tag,omitempty"`
	Name      string            `json:"name,omitempty"`
	ExpLevel  int               `json:"expLevel,omitempty"`
	Rank      int               `json:"rank,omitempty"`
	EloRating int               `json:"eloRating,omitempty"`
}

type PlayerRankingClan struct {
	BadgeId   int            `json:"badgeId,omitempty"`
	Tag       string         `json:"tag,omitempty"`
	Name      string         `json:"name,omitempty"`
	BadgeUrls map[string]any `json:"badgeUrls,omitempty"`
}

// LeagueSeasonList is a paginated league season response.
type LeagueSeasonList struct {
	Items  []LeagueSeason `json:"items,omitempty"`
	Paging Paging         `json:"paging,omitempty"`
}

type LeagueSeason struct {
	Id string `json:"id,omitempty"`
}

// PlayerRankingList is a paginated player ranking response.
type PlayerRankingList struct {
	Items  []PlayerRanking `json:"items,omitempty"`
	Paging Paging          `json:"paging,omitempty"`
}

type PlayerRanking struct {
	Clan         PlayerRankingClan `json:"clan,omitempty"`
	Arena        Arena             `json:"arena,omitempty"`
	Tag          string            `json:"tag,omitempty"`
	Name         string            `json:"name,omitempty"`
	ExpLevel     int               `json:"expLevel,omitempty"`
	Rank         int               `json:"rank,omitempty"`
	PreviousRank int               `json:"previousRank,omitempty"`
	Trophies     int               `json:"trophies,omitempty"`
}

// ClanRankingList is a paginated clan ranking response.
type ClanRankingList struct {
	Items  []ClanRanking `json:"items,omitempty"`
	Paging Paging        `json:"paging,omitempty"`
}

type ClanRanking struct {
	ClanScore    int            `json:"clanScore,omitempty"`
	BadgeId      int            `json:"badgeId,omitempty"`
	Location     Location       `json:"location,omitempty"`
	Members      int            `json:"members,omitempty"`
	Tag          string         `json:"tag,omitempty"`
	Name         string         `json:"name,omitempty"`
	Rank         int            `json:"rank,omitempty"`
	PreviousRank int            `json:"previousRank,omitempty"`
	BadgeUrls    map[string]any `json:"badgeUrls,omitempty"`
}

type PathOfLegendSeasonResult struct {
	Trophies     int `json:"trophies,omitempty"`
	LeagueNumber int `json:"leagueNumber,omitempty"`
	Rank         int `json:"rank,omitempty"`
}

type PlayerLeagueStatistics struct {
	PreviousSeason LeagueSeasonResult `json:"previousSeason,omitempty"`
	CurrentSeason  LeagueSeasonResult `json:"currentSeason,omitempty"`
	BestSeason     LeagueSeasonResult `json:"bestSeason,omitempty"`
}

type LeagueSeasonResult struct {
	Trophies     int    `json:"trophies,omitempty"`
	Rank         int    `json:"rank,omitempty"`
	BestTrophies int    `json:"bestTrophies,omitempty"`
	Id           string `json:"id,omitempty"`
}
