package crgo

type Replay struct {
	ReplayInfo string   `json:"replayInfo,omitempty"`
	ReplayData JsonNode `json:"replayData,omitempty"`
	Version    Version  `json:"version,omitempty"`
	Tag        string   `json:"tag,omitempty"`
	BattleTime string   `json:"battleTime,omitempty"`
	ViewCount  int      `json:"viewCount,omitempty"`
	ShareCount int      `json:"shareCount,omitempty"`
}

type Match struct {
	Invites   map[string]any `json:"invites,omitempty"`
	StartTime string         `json:"startTime,omitempty"`
	State     MatchState     `json:"state,omitempty"`
	Battle    Battle         `json:"battle,omitempty"`
	Tag       string         `json:"tag,omitempty"`
}

type MatchState string

const (
	MatchStatePending    MatchState = "PENDING"
	MatchStateCancelled  MatchState = "CANCELLED"
	MatchStateTimedOut   MatchState = "TIMED_OUT"
	MatchStateInProgress MatchState = "IN_PROGRESS"
	MatchStateComplete   MatchState = "COMPLETE"
	MatchStateFailed     MatchState = "FAILED"
)

type Battle struct {
	GameMode                GameMode                `json:"gameMode,omitempty"`
	Arena                   Arena                   `json:"arena,omitempty"`
	Type                    BattleType              `json:"type,omitempty"`
	DeckSelection           BattleDeckSelection     `json:"deckSelection,omitempty"`
	Team                    PlayerBattleDataList    `json:"team,omitempty"`
	Opponent                PlayerBattleDataList    `json:"opponent,omitempty"`
	ChallengeWinCountBefore int                     `json:"challengeWinCountBefore,omitempty"`
	BoatBattleSide          string                  `json:"boatBattleSide,omitempty"`
	BoatBattleWon           bool                    `json:"boatBattleWon,omitempty"`
	NewTowersDestroyed      int                     `json:"newTowersDestroyed,omitempty"`
	PrevTowersDestroyed     int                     `json:"prevTowersDestroyed,omitempty"`
	RemainingTowers         int                     `json:"remainingTowers,omitempty"`
	LeagueNumber            int                     `json:"leagueNumber,omitempty"`
	BattleTime              string                  `json:"battleTime,omitempty"`
	ChallengeId             int                     `json:"challengeId,omitempty"`
	TournamentTag           string                  `json:"tournamentTag,omitempty"`
	ChallengeTitle          string                  `json:"challengeTitle,omitempty"`
	IsLadderTournament      bool                    `json:"isLadderTournament,omitempty"`
	IsHostedMatch           bool                    `json:"isHostedMatch,omitempty"`
	EventTag                string                  `json:"eventTag,omitempty"`
	Modifiers               PlayerBattleAugmentList `json:"modifiers,omitempty"`
}

type BattleType string

const (
	BattleTypePvp                    BattleType = "PVP"
	BattleTypePve                    BattleType = "PVE"
	BattleTypeClanmate               BattleType = "CLANMATE"
	BattleTypeTournament             BattleType = "TOURNAMENT"
	BattleTypeFriendly               BattleType = "FRIENDLY"
	BattleTypeSurvival               BattleType = "SURVIVAL"
	BattleTypePvp2V2                 BattleType = "PVP2v2"
	BattleTypeClanmate2V2            BattleType = "CLANMATE2v2"
	BattleTypeChallenge2V2           BattleType = "CHALLENGE2v2"
	BattleTypeClanwarCollectionDay   BattleType = "CLANWAR_COLLECTION_DAY"
	BattleTypeClanwarWarDay          BattleType = "CLANWAR_WAR_DAY"
	BattleTypeCasual1V1              BattleType = "CASUAL_1V1"
	BattleTypeCasual2V2              BattleType = "CASUAL_2V2"
	BattleTypeBoatBattle             BattleType = "BOAT_BATTLE"
	BattleTypeBoatBattlePractice     BattleType = "BOAT_BATTLE_PRACTICE"
	BattleTypeRiverRacePvp           BattleType = "RIVER_RACE_PVP"
	BattleTypeRiverRaceDuel          BattleType = "RIVER_RACE_DUEL"
	BattleTypeRiverRaceDuelColosseum BattleType = "RIVER_RACE_DUEL_COLOSSEUM"
	BattleTypeTutorial               BattleType = "TUTORIAL"
	BattleTypePathOfLegend           BattleType = "PATH_OF_LEGEND"
	BattleTypeSeasonalBattle         BattleType = "SEASONAL_BATTLE"
	BattleTypePractice               BattleType = "PRACTICE"
	BattleTypeTrail                  BattleType = "TRAIL"
	BattleTypeUnknown                BattleType = "UNKNOWN"
)

type BattleDeckSelection string

const (
	BattleDeckSelectionCollection       BattleDeckSelection = "COLLECTION"
	BattleDeckSelectionDraft            BattleDeckSelection = "DRAFT"
	BattleDeckSelectionDraftCompetitive BattleDeckSelection = "DRAFT_COMPETITIVE"
	BattleDeckSelectionPredefined       BattleDeckSelection = "PREDEFINED"
	BattleDeckSelectionEventDeck        BattleDeckSelection = "EVENT_DECK"
	BattleDeckSelectionPick             BattleDeckSelection = "PICK"
	BattleDeckSelectionWardeckPick      BattleDeckSelection = "WARDECK_PICK"
	BattleDeckSelectionQuaddeckPick     BattleDeckSelection = "QUADDECK_PICK"
	BattleDeckSelectionUnknown          BattleDeckSelection = "UNKNOWN"
)

type PlayerBattleAugmentList []PlayerBattleAugment

type PlayerBattleAugment struct {
	Tag       string     `json:"tag,omitempty"`
	Modifiers StringList `json:"modifiers,omitempty"`
}

type PlayerBattleDataList []PlayerBattleData

type PlayerBattleData struct {
	Clan                    PlayerClan            `json:"clan,omitempty"`
	Cards                   PlayerItemLevelList   `json:"cards,omitempty"`
	SupportCards            PlayerItemLevelList   `json:"supportCards,omitempty"`
	GlobalRank              int                   `json:"globalRank,omitempty"`
	Crowns                  int                   `json:"crowns,omitempty"`
	PrincessTowersHitPoints IntegerList           `json:"princessTowersHitPoints,omitempty"`
	ElixirLeaked            Float                 `json:"elixirLeaked,omitempty"`
	Rounds                  PlayerBattleRoundList `json:"rounds,omitempty"`
	Tag                     string                `json:"tag,omitempty"`
	Name                    string                `json:"name,omitempty"`
	StartingTrophies        int                   `json:"startingTrophies,omitempty"`
	TrophyChange            int                   `json:"trophyChange,omitempty"`
	KingTowerHitPoints      int                   `json:"kingTowerHitPoints,omitempty"`
}

type PlayerBattleRoundList []PlayerBattleRound

type PlayerBattleRound struct {
	Cards                   PlayerItemLevelList `json:"cards,omitempty"`
	ElixirLeaked            Float               `json:"elixirLeaked,omitempty"`
	Crowns                  int                 `json:"crowns,omitempty"`
	KingTowerHitPoints      int                 `json:"kingTowerHitPoints,omitempty"`
	PrincessTowersHitPoints IntegerList         `json:"princessTowersHitPoints,omitempty"`
}

type CancelMatchResponse struct {
	Success bool `json:"success,omitempty"`
}

type BattleList []Battle

type RegisterMatchRequest struct {
	PlayerTags StringList                   `json:"playerTags,omitempty"`
	GameMode   RegisterMatchRequestGameMode `json:"gameMode,omitempty"`
}

type RegisterMatchRequestGameMode string

const (
	RegisterMatchRequestGameModeRegular      RegisterMatchRequestGameMode = "REGULAR"
	RegisterMatchRequestGameModeTeamVsTeam   RegisterMatchRequestGameMode = "TEAM_VS_TEAM"
	RegisterMatchRequestGameModeDoubleElixir RegisterMatchRequestGameMode = "DOUBLE_ELIXIR"
	RegisterMatchRequestGameModeTripleElixir RegisterMatchRequestGameMode = "TRIPLE_ELIXIR"
	RegisterMatchRequestGameModeRage         RegisterMatchRequestGameMode = "RAGE"
	RegisterMatchRequestGameModeSuddenDeath  RegisterMatchRequestGameMode = "SUDDEN_DEATH"
	RegisterMatchRequestGameModeTouchdown    RegisterMatchRequestGameMode = "TOUCHDOWN"
	RegisterMatchRequestGameModeRampUp       RegisterMatchRequestGameMode = "RAMP_UP"
	RegisterMatchRequestGameModeDraft        RegisterMatchRequestGameMode = "DRAFT"
	RegisterMatchRequestGameModeMirror       RegisterMatchRequestGameMode = "MIRROR"
	RegisterMatchRequestGameModeDragonHunt   RegisterMatchRequestGameMode = "DRAGON_HUNT"
	RegisterMatchRequestGameModeTripleDraft  RegisterMatchRequestGameMode = "TRIPLE_DRAFT"
	RegisterMatchRequestGameModeBestOf3      RegisterMatchRequestGameMode = "BEST_OF_3"
	RegisterMatchRequestGameModeMegaDraft    RegisterMatchRequestGameMode = "MEGA_DRAFT"
	RegisterMatchRequestGameModeHeist        RegisterMatchRequestGameMode = "HEIST"
)

type RegisterMatchResponse struct {
	Tag string `json:"tag,omitempty"`
}
