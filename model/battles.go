package model

type Replay struct {
	ReplayInfo string   `json:"replayInfo"`
	ReplayData JsonNode `json:"replayData"`
	Version    Version  `json:"version"`
	Tag        string   `json:"tag"`
	BattleTime string   `json:"battleTime"`
	ViewCount  int      `json:"viewCount"`
	ShareCount int      `json:"shareCount"`
}

type Match struct {
	Invites   map[string]any `json:"invites"`
	StartTime string         `json:"startTime"`
	State     MatchState     `json:"state"`
	Battle    Battle         `json:"battle"`
	Tag       string         `json:"tag"`
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
	GameMode                GameMode                `json:"gameMode"`
	Arena                   Arena                   `json:"arena"`
	Type                    BattleType              `json:"type"`
	DeckSelection           BattleDeckSelection     `json:"deckSelection"`
	Team                    PlayerBattleDataList    `json:"team"`
	Opponent                PlayerBattleDataList    `json:"opponent"`
	ChallengeWinCountBefore int                     `json:"challengeWinCountBefore"`
	BoatBattleSide          string                  `json:"boatBattleSide"`
	BoatBattleWon           bool                    `json:"boatBattleWon"`
	NewTowersDestroyed      int                     `json:"newTowersDestroyed"`
	PrevTowersDestroyed     int                     `json:"prevTowersDestroyed"`
	RemainingTowers         int                     `json:"remainingTowers"`
	LeagueNumber            int                     `json:"leagueNumber"`
	BattleTime              string                  `json:"battleTime"`
	ChallengeId             int                     `json:"challengeId"`
	TournamentTag           string                  `json:"tournamentTag"`
	ChallengeTitle          string                  `json:"challengeTitle"`
	IsLadderTournament      bool                    `json:"isLadderTournament"`
	IsHostedMatch           bool                    `json:"isHostedMatch"`
	EventTag                string                  `json:"eventTag"`
	Modifiers               PlayerBattleAugmentList `json:"modifiers"`
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
	Tag       string     `json:"tag"`
	Modifiers StringList `json:"modifiers"`
}

type PlayerBattleDataList []PlayerBattleData

type PlayerBattleData struct {
	Clan                    PlayerClan            `json:"clan"`
	Cards                   PlayerItemLevelList   `json:"cards"`
	SupportCards            PlayerItemLevelList   `json:"supportCards"`
	GlobalRank              int                   `json:"globalRank"`
	Crowns                  int                   `json:"crowns"`
	PrincessTowersHitPoints IntegerList           `json:"princessTowersHitPoints"`
	ElixirLeaked            Float                 `json:"elixirLeaked"`
	Rounds                  PlayerBattleRoundList `json:"rounds"`
	Tag                     string                `json:"tag"`
	Name                    string                `json:"name"`
	StartingTrophies        int                   `json:"startingTrophies"`
	TrophyChange            int                   `json:"trophyChange"`
	KingTowerHitPoints      int                   `json:"kingTowerHitPoints"`
}

type PlayerBattleRoundList []PlayerBattleRound

type PlayerBattleRound struct {
	Cards                   PlayerItemLevelList `json:"cards"`
	ElixirLeaked            Float               `json:"elixirLeaked"`
	Crowns                  int                 `json:"crowns"`
	KingTowerHitPoints      int                 `json:"kingTowerHitPoints"`
	PrincessTowersHitPoints IntegerList         `json:"princessTowersHitPoints"`
}

type CancelMatchResponse struct {
	Success bool `json:"success"`
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
	Tag string `json:"tag"`
}
