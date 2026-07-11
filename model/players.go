package model

type PlayerClan struct {
	BadgeId   int            `json:"badgeId"`
	Tag       string         `json:"tag"`
	Name      string         `json:"name"`
	BadgeUrls map[string]any `json:"badgeUrls"`
}

type Player struct {
	Clan                            PlayerClan                    `json:"clan"`
	LegacyTrophyRoadHighScore       int                           `json:"legacyTrophyRoadHighScore"`
	CurrentDeck                     PlayerItemLevelList           `json:"currentDeck"`
	CurrentDeckSupportCards         PlayerItemLevelList           `json:"currentDeckSupportCards"`
	Arena                           Arena                         `json:"arena"`
	TotalDonations                  int                           `json:"totalDonations"`
	Role                            PlayerRole                    `json:"role"`
	Wins                            int                           `json:"wins"`
	Losses                          int                           `json:"losses"`
	LeagueStatistics                PlayerLeagueStatistics        `json:"leagueStatistics"`
	Cards                           PlayerItemLevelList           `json:"cards"`
	SupportCards                    PlayerItemLevelList           `json:"supportCards"`
	CurrentFavouriteCard            Item                          `json:"currentFavouriteCard"`
	Badges                          PlayerAchievementBadgeList    `json:"badges"`
	Tag                             string                        `json:"tag"`
	Name                            string                        `json:"name"`
	ExpLevel                        int                           `json:"expLevel"`
	Trophies                        int                           `json:"trophies"`
	BestTrophies                    int                           `json:"bestTrophies"`
	Donations                       int                           `json:"donations"`
	DonationsReceived               int                           `json:"donationsReceived"`
	Achievements                    PlayerAchievementProgressList `json:"achievements"`
	BattleCount                     int                           `json:"battleCount"`
	ThreeCrownWins                  int                           `json:"threeCrownWins"`
	ChallengeCardsWon               int                           `json:"challengeCardsWon"`
	ChallengeMaxWins                int                           `json:"challengeMaxWins"`
	TournamentCardsWon              int                           `json:"tournamentCardsWon"`
	TournamentBattleCount           int                           `json:"tournamentBattleCount"`
	CurrentWinLoseStreak            int                           `json:"currentWinLoseStreak"`
	WarDayWins                      int                           `json:"warDayWins"`
	ClanCardsCollected              int                           `json:"clanCardsCollected"`
	StarPoints                      int                           `json:"starPoints"`
	ExpPoints                       int                           `json:"expPoints"`
	TotalExpPoints                  int                           `json:"totalExpPoints"`
	CurrentPathOfLegendSeasonResult PathOfLegendSeasonResult      `json:"currentPathOfLegendSeasonResult"`
	LastPathOfLegendSeasonResult    PathOfLegendSeasonResult      `json:"lastPathOfLegendSeasonResult"`
	BestPathOfLegendSeasonResult    PathOfLegendSeasonResult      `json:"bestPathOfLegendSeasonResult"`
	Progress                        map[string]any                `json:"progress"`
}

type PlayerRole string

const (
	PlayerRoleNotMember PlayerRole = "NOT_MEMBER"
	PlayerRoleMember    PlayerRole = "MEMBER"
	PlayerRoleLeader    PlayerRole = "LEADER"
	PlayerRoleAdmin     PlayerRole = "ADMIN"
	PlayerRoleColeader  PlayerRole = "COLEADER"
)

type PlayerAchievementProgressList []PlayerAchievementProgress

type PlayerAchievementProgress struct {
	Stars          int               `json:"stars"`
	Value          int               `json:"value"`
	Name           JsonLocalizedName `json:"name"`
	Target         int               `json:"target"`
	Info           JsonLocalizedName `json:"info"`
	CompletionInfo JsonLocalizedName `json:"completionInfo"`
}

type PlayerAchievementBadgeList []PlayerAchievementBadge

type PlayerAchievementBadge struct {
	IconUrls map[string]any `json:"iconUrls"`
	MaxLevel int            `json:"maxLevel"`
	Progress int            `json:"progress"`
	Level    int            `json:"level"`
	Target   int            `json:"target"`
	Name     string         `json:"name"`
}
