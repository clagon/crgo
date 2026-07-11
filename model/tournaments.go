package model

type LadderTournamentRankingList []LadderTournamentRanking

type LadderTournamentRanking struct {
	Clan         PlayerRankingClan `json:"clan"`
	Wins         int               `json:"wins"`
	Losses       int               `json:"losses"`
	Tag          string            `json:"tag"`
	Name         string            `json:"name"`
	Rank         int               `json:"rank"`
	PreviousRank int               `json:"previousRank"`
}

// LadderTournamentList is a paginated global tournament response.
type LadderTournamentList struct {
	Items  []LadderTournament `json:"items"`
	Paging Paging             `json:"paging"`
}

type LadderTournament struct {
	GameMode         GameMode                    `json:"gameMode"`
	MaxLosses        int                         `json:"maxLosses"`
	MinExpLevel      int                         `json:"minExpLevel"`
	TournamentLevel  int                         `json:"tournamentLevel"`
	MilestoneRewards SurvivalMilestoneRewardList `json:"milestoneRewards"`
	FreeTierRewards  SurvivalMilestoneRewardList `json:"freeTierRewards"`
	Tag              string                      `json:"tag"`
	Title            string                      `json:"title"`
	StartTime        string                      `json:"startTime"`
	EndTime          string                      `json:"endTime"`
	TopRankReward    SurvivalMilestoneRewardList `json:"topRankReward"`
	MaxTopRewardRank int                         `json:"maxTopRewardRank"`
}

type SurvivalMilestoneRewardList []SurvivalMilestoneReward

type SurvivalMilestoneReward struct {
	Rarity         SurvivalMilestoneRewardRarity   `json:"rarity"`
	Chest          string                          `json:"chest"`
	Resource       SurvivalMilestoneRewardResource `json:"resource"`
	Type           SurvivalMilestoneRewardType     `json:"type"`
	Amount         int                             `json:"amount"`
	Card           Item                            `json:"card"`
	ConsumableName string                          `json:"consumableName"`
	Wins           int                             `json:"wins"`
}

type SurvivalMilestoneRewardRarity string

const (
	SurvivalMilestoneRewardRarityCommon    SurvivalMilestoneRewardRarity = "COMMON"
	SurvivalMilestoneRewardRarityRare      SurvivalMilestoneRewardRarity = "RARE"
	SurvivalMilestoneRewardRarityEpic      SurvivalMilestoneRewardRarity = "EPIC"
	SurvivalMilestoneRewardRarityLegendary SurvivalMilestoneRewardRarity = "LEGENDARY"
	SurvivalMilestoneRewardRarityChampion  SurvivalMilestoneRewardRarity = "CHAMPION"
)

type SurvivalMilestoneRewardResource string

const (
	SurvivalMilestoneRewardResourceGold    SurvivalMilestoneRewardResource = "GOLD"
	SurvivalMilestoneRewardResourceUnknown SurvivalMilestoneRewardResource = "UNKNOWN"
)

type SurvivalMilestoneRewardType string

const (
	SurvivalMilestoneRewardTypeNone            SurvivalMilestoneRewardType = "NONE"
	SurvivalMilestoneRewardTypeCardStack       SurvivalMilestoneRewardType = "CARD_STACK"
	SurvivalMilestoneRewardTypeChest           SurvivalMilestoneRewardType = "CHEST"
	SurvivalMilestoneRewardTypeCardStackRandom SurvivalMilestoneRewardType = "CARD_STACK_RANDOM"
	SurvivalMilestoneRewardTypeResource        SurvivalMilestoneRewardType = "RESOURCE"
	SurvivalMilestoneRewardTypeTradeToken      SurvivalMilestoneRewardType = "TRADE_TOKEN"
	SurvivalMilestoneRewardTypeConsumable      SurvivalMilestoneRewardType = "CONSUMABLE"
)

type Tournament struct {
	GameMode            GameMode             `json:"gameMode"`
	MembersList         TournamentMemberList `json:"membersList"`
	Status              TournamentStatus     `json:"status"`
	PreparationDuration int                  `json:"preparationDuration"`
	CreatedTime         string               `json:"createdTime"`
	StartedTime         string               `json:"startedTime"`
	EndedTime           string               `json:"endedTime"`
	FirstPlaceCardPrize int                  `json:"firstPlaceCardPrize"`
	Duration            int                  `json:"duration"`
	Type                TournamentType       `json:"type"`
	Tag                 string               `json:"tag"`
	CreatorTag          string               `json:"creatorTag"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	Capacity            int                  `json:"capacity"`
	MaxCapacity         int                  `json:"maxCapacity"`
	LevelCap            int                  `json:"levelCap"`
}

type TournamentStatus string

const (
	TournamentStatusInPreparation TournamentStatus = "IN_PREPARATION"
	TournamentStatusInProgress    TournamentStatus = "IN_PROGRESS"
	TournamentStatusEnded         TournamentStatus = "ENDED"
	TournamentStatusUnknown       TournamentStatus = "UNKNOWN"
)

type TournamentType string

const (
	TournamentTypeOpen              TournamentType = "OPEN"
	TournamentTypePasswordProtected TournamentType = "PASSWORD_PROTECTED"
	TournamentTypeUnknown           TournamentType = "UNKNOWN"
)

type TournamentMemberList []TournamentMember

type TournamentMember struct {
	Clan         PlayerClan `json:"clan"`
	PreviousRank int        `json:"previousRank"`
	Rank         int        `json:"rank"`
	Tag          string     `json:"tag"`
	Name         string     `json:"name"`
	Score        int        `json:"score"`
}

// TournamentHeaderList is a paginated tournament search response.
type TournamentHeaderList struct {
	Items  []TournamentHeader `json:"items"`
	Paging Paging             `json:"paging"`
}

type TournamentHeader struct {
	GameMode            GameMode               `json:"gameMode"`
	Status              TournamentHeaderStatus `json:"status"`
	PreparationDuration int                    `json:"preparationDuration"`
	CreatedTime         string                 `json:"createdTime"`
	FirstPlaceCardPrize int                    `json:"firstPlaceCardPrize"`
	Duration            int                    `json:"duration"`
	Type                TournamentHeaderType   `json:"type"`
	Tag                 string                 `json:"tag"`
	CreatorTag          string                 `json:"creatorTag"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	Capacity            int                    `json:"capacity"`
	MaxCapacity         int                    `json:"maxCapacity"`
	LevelCap            int                    `json:"levelCap"`
}

type TournamentHeaderStatus string

const (
	TournamentHeaderStatusInPreparation TournamentHeaderStatus = "IN_PREPARATION"
	TournamentHeaderStatusInProgress    TournamentHeaderStatus = "IN_PROGRESS"
	TournamentHeaderStatusEnded         TournamentHeaderStatus = "ENDED"
	TournamentHeaderStatusUnknown       TournamentHeaderStatus = "UNKNOWN"
)

type TournamentHeaderType string

const (
	TournamentHeaderTypeOpen              TournamentHeaderType = "OPEN"
	TournamentHeaderTypePasswordProtected TournamentHeaderType = "PASSWORD_PROTECTED"
	TournamentHeaderTypeUnknown           TournamentHeaderType = "UNKNOWN"
)
