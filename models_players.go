package crgo

type PlayerClan struct {
	BadgeId   int            `json:"badgeId,omitempty"`
	Tag       string         `json:"tag,omitempty"`
	Name      string         `json:"name,omitempty"`
	BadgeUrls map[string]any `json:"badgeUrls,omitempty"`
}

type Player struct {
	Clan                            PlayerClan                    `json:"clan,omitempty"`
	LegacyTrophyRoadHighScore       int                           `json:"legacyTrophyRoadHighScore,omitempty"`
	CurrentDeck                     PlayerItemLevelList           `json:"currentDeck,omitempty"`
	CurrentDeckSupportCards         PlayerItemLevelList           `json:"currentDeckSupportCards,omitempty"`
	Arena                           Arena                         `json:"arena,omitempty"`
	TotalDonations                  int                           `json:"totalDonations,omitempty"`
	Role                            PlayerRole                    `json:"role,omitempty"`
	Wins                            int                           `json:"wins,omitempty"`
	Losses                          int                           `json:"losses,omitempty"`
	LeagueStatistics                PlayerLeagueStatistics        `json:"leagueStatistics,omitempty"`
	Cards                           PlayerItemLevelList           `json:"cards,omitempty"`
	SupportCards                    PlayerItemLevelList           `json:"supportCards,omitempty"`
	CurrentFavouriteCard            Item                          `json:"currentFavouriteCard,omitempty"`
	Badges                          PlayerAchievementBadgeList    `json:"badges,omitempty"`
	Tag                             string                        `json:"tag,omitempty"`
	Name                            string                        `json:"name,omitempty"`
	ExpLevel                        int                           `json:"expLevel,omitempty"`
	Trophies                        int                           `json:"trophies,omitempty"`
	BestTrophies                    int                           `json:"bestTrophies,omitempty"`
	Donations                       int                           `json:"donations,omitempty"`
	DonationsReceived               int                           `json:"donationsReceived,omitempty"`
	Achievements                    PlayerAchievementProgressList `json:"achievements,omitempty"`
	BattleCount                     int                           `json:"battleCount,omitempty"`
	ThreeCrownWins                  int                           `json:"threeCrownWins,omitempty"`
	ChallengeCardsWon               int                           `json:"challengeCardsWon,omitempty"`
	ChallengeMaxWins                int                           `json:"challengeMaxWins,omitempty"`
	TournamentCardsWon              int                           `json:"tournamentCardsWon,omitempty"`
	TournamentBattleCount           int                           `json:"tournamentBattleCount,omitempty"`
	CurrentWinLoseStreak            int                           `json:"currentWinLoseStreak,omitempty"`
	WarDayWins                      int                           `json:"warDayWins,omitempty"`
	ClanCardsCollected              int                           `json:"clanCardsCollected,omitempty"`
	StarPoints                      int                           `json:"starPoints,omitempty"`
	ExpPoints                       int                           `json:"expPoints,omitempty"`
	TotalExpPoints                  int                           `json:"totalExpPoints,omitempty"`
	CurrentPathOfLegendSeasonResult PathOfLegendSeasonResult      `json:"currentPathOfLegendSeasonResult,omitempty"`
	LastPathOfLegendSeasonResult    PathOfLegendSeasonResult      `json:"lastPathOfLegendSeasonResult,omitempty"`
	BestPathOfLegendSeasonResult    PathOfLegendSeasonResult      `json:"bestPathOfLegendSeasonResult,omitempty"`
	Progress                        map[string]any                `json:"progress,omitempty"`
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
	Stars          int               `json:"stars,omitempty"`
	Value          int               `json:"value,omitempty"`
	Name           JsonLocalizedName `json:"name,omitempty"`
	Target         int               `json:"target,omitempty"`
	Info           JsonLocalizedName `json:"info,omitempty"`
	CompletionInfo JsonLocalizedName `json:"completionInfo,omitempty"`
}

type PlayerAchievementBadgeList []PlayerAchievementBadge

type PlayerAchievementBadge struct {
	IconUrls map[string]any `json:"iconUrls,omitempty"`
	MaxLevel int            `json:"maxLevel,omitempty"`
	Progress int            `json:"progress,omitempty"`
	Level    int            `json:"level,omitempty"`
	Target   int            `json:"target,omitempty"`
	Name     string         `json:"name,omitempty"`
}
