package model

type LadderTournamentRankingList []LadderTournamentRanking

type LadderTournamentRanking struct {
	Clan         PlayerRankingClan `json:"clan,omitempty"`
	Wins         int               `json:"wins,omitempty"`
	Losses       int               `json:"losses,omitempty"`
	Tag          string            `json:"tag,omitempty"`
	Name         string            `json:"name,omitempty"`
	Rank         int               `json:"rank,omitempty"`
	PreviousRank int               `json:"previousRank,omitempty"`
}

// LadderTournamentList is a paginated global tournament response.
type LadderTournamentList struct {
	Items  []LadderTournament `json:"items,omitempty"`
	Paging Paging             `json:"paging,omitempty"`
}

type LadderTournament struct {
	GameMode         GameMode                    `json:"gameMode,omitempty"`
	MaxLosses        int                         `json:"maxLosses,omitempty"`
	MinExpLevel      int                         `json:"minExpLevel,omitempty"`
	TournamentLevel  int                         `json:"tournamentLevel,omitempty"`
	MilestoneRewards SurvivalMilestoneRewardList `json:"milestoneRewards,omitempty"`
	FreeTierRewards  SurvivalMilestoneRewardList `json:"freeTierRewards,omitempty"`
	Tag              string                      `json:"tag,omitempty"`
	Title            string                      `json:"title,omitempty"`
	StartTime        string                      `json:"startTime,omitempty"`
	EndTime          string                      `json:"endTime,omitempty"`
	TopRankReward    SurvivalMilestoneRewardList `json:"topRankReward,omitempty"`
	MaxTopRewardRank int                         `json:"maxTopRewardRank,omitempty"`
}

type SurvivalMilestoneRewardList []SurvivalMilestoneReward

type SurvivalMilestoneReward struct {
	Rarity         SurvivalMilestoneRewardRarity   `json:"rarity,omitempty"`
	Chest          string                          `json:"chest,omitempty"`
	Resource       SurvivalMilestoneRewardResource `json:"resource,omitempty"`
	Type           SurvivalMilestoneRewardType     `json:"type,omitempty"`
	Amount         int                             `json:"amount,omitempty"`
	Card           Item                            `json:"card,omitempty"`
	ConsumableName string                          `json:"consumableName,omitempty"`
	Wins           int                             `json:"wins,omitempty"`
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
	GameMode            GameMode             `json:"gameMode,omitempty"`
	MembersList         TournamentMemberList `json:"membersList,omitempty"`
	Status              TournamentStatus     `json:"status,omitempty"`
	PreparationDuration int                  `json:"preparationDuration,omitempty"`
	CreatedTime         string               `json:"createdTime,omitempty"`
	StartedTime         string               `json:"startedTime,omitempty"`
	EndedTime           string               `json:"endedTime,omitempty"`
	FirstPlaceCardPrize int                  `json:"firstPlaceCardPrize,omitempty"`
	Duration            int                  `json:"duration,omitempty"`
	Type                TournamentType       `json:"type,omitempty"`
	Tag                 string               `json:"tag,omitempty"`
	CreatorTag          string               `json:"creatorTag,omitempty"`
	Name                string               `json:"name,omitempty"`
	Description         string               `json:"description,omitempty"`
	Capacity            int                  `json:"capacity,omitempty"`
	MaxCapacity         int                  `json:"maxCapacity,omitempty"`
	LevelCap            int                  `json:"levelCap,omitempty"`
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
	Clan         PlayerClan `json:"clan,omitempty"`
	PreviousRank int        `json:"previousRank,omitempty"`
	Rank         int        `json:"rank,omitempty"`
	Tag          string     `json:"tag,omitempty"`
	Name         string     `json:"name,omitempty"`
	Score        int        `json:"score,omitempty"`
}

// TournamentHeaderList is a paginated tournament search response.
type TournamentHeaderList struct {
	Items  []TournamentHeader `json:"items,omitempty"`
	Paging Paging             `json:"paging,omitempty"`
}

type TournamentHeader struct {
	GameMode            GameMode               `json:"gameMode,omitempty"`
	Status              TournamentHeaderStatus `json:"status,omitempty"`
	PreparationDuration int                    `json:"preparationDuration,omitempty"`
	CreatedTime         string                 `json:"createdTime,omitempty"`
	FirstPlaceCardPrize int                    `json:"firstPlaceCardPrize,omitempty"`
	Duration            int                    `json:"duration,omitempty"`
	Type                TournamentHeaderType   `json:"type,omitempty"`
	Tag                 string                 `json:"tag,omitempty"`
	CreatorTag          string                 `json:"creatorTag,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Capacity            int                    `json:"capacity,omitempty"`
	MaxCapacity         int                    `json:"maxCapacity,omitempty"`
	LevelCap            int                    `json:"levelCap,omitempty"`
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
