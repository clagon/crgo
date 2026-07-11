package model

type CurrentRiverRace struct {
	State             CurrentRiverRaceState      `json:"state,omitempty"`
	Clan              RiverRaceClan              `json:"clan,omitempty"`
	Clans             RiverRaceClanList          `json:"clans,omitempty"`
	CollectionEndTime string                     `json:"collectionEndTime,omitempty"`
	WarEndTime        string                     `json:"warEndTime,omitempty"`
	SectionIndex      int                        `json:"sectionIndex,omitempty"`
	PeriodIndex       int                        `json:"periodIndex,omitempty"`
	PeriodType        CurrentRiverRacePeriodType `json:"periodType,omitempty"`
	PeriodLogs        PeriodLogList              `json:"periodLogs,omitempty"`
}

type CurrentRiverRaceState string

const (
	CurrentRiverRaceStateClanNotFound CurrentRiverRaceState = "CLAN_NOT_FOUND"
	CurrentRiverRaceStateAccessDenied CurrentRiverRaceState = "ACCESS_DENIED"
	CurrentRiverRaceStateMatchmaking  CurrentRiverRaceState = "MATCHMAKING"
	CurrentRiverRaceStateMatched      CurrentRiverRaceState = "MATCHED"
	CurrentRiverRaceStateFull         CurrentRiverRaceState = "FULL"
	CurrentRiverRaceStateEnded        CurrentRiverRaceState = "ENDED"
)

type CurrentRiverRacePeriodType string

const (
	CurrentRiverRacePeriodTypeTraining  CurrentRiverRacePeriodType = "TRAINING"
	CurrentRiverRacePeriodTypeWarDay    CurrentRiverRacePeriodType = "WAR_DAY"
	CurrentRiverRacePeriodTypeColosseum CurrentRiverRacePeriodType = "COLOSSEUM"
)

type PeriodLogList []PeriodLog

type PeriodLog struct {
	Items       PeriodLogEntryList `json:"items,omitempty"`
	PeriodIndex int                `json:"periodIndex,omitempty"`
}

type PeriodLogEntryList []PeriodLogEntry

type PeriodLogEntry struct {
	Clan                       PeriodLogEntryClan `json:"clan,omitempty"`
	PointsEarned               int                `json:"pointsEarned,omitempty"`
	ProgressStartOfDay         int                `json:"progressStartOfDay,omitempty"`
	ProgressEndOfDay           int                `json:"progressEndOfDay,omitempty"`
	EndOfDayRank               int                `json:"endOfDayRank,omitempty"`
	ProgressEarned             int                `json:"progressEarned,omitempty"`
	NumOfDefensesRemaining     int                `json:"numOfDefensesRemaining,omitempty"`
	ProgressEarnedFromDefenses int                `json:"progressEarnedFromDefenses,omitempty"`
}

type PeriodLogEntryClan struct {
	Tag string `json:"tag,omitempty"`
}

type RiverRaceClanList []RiverRaceClan

type RiverRaceClan struct {
	Tag          string                   `json:"tag,omitempty"`
	ClanScore    int                      `json:"clanScore,omitempty"`
	BadgeId      int                      `json:"badgeId,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Fame         int                      `json:"fame,omitempty"`
	RepairPoints int                      `json:"repairPoints,omitempty"`
	FinishTime   string                   `json:"finishTime,omitempty"`
	Participants RiverRaceParticipantList `json:"participants,omitempty"`
	PeriodPoints int                      `json:"periodPoints,omitempty"`
}

type RiverRaceParticipantList []RiverRaceParticipant

type RiverRaceParticipant struct {
	Tag            string `json:"tag,omitempty"`
	Name           string `json:"name,omitempty"`
	Fame           int    `json:"fame,omitempty"`
	RepairPoints   int    `json:"repairPoints,omitempty"`
	BoatAttacks    int    `json:"boatAttacks,omitempty"`
	DecksUsed      int    `json:"decksUsed,omitempty"`
	DecksUsedToday int    `json:"decksUsedToday,omitempty"`
}

type RiverRaceLog []RiverRaceLogEntry

type RiverRaceLogEntry struct {
	Standings    RiverRaceStandingList `json:"standings,omitempty"`
	SeasonId     int                   `json:"seasonId,omitempty"`
	CreatedDate  string                `json:"createdDate,omitempty"`
	SectionIndex int                   `json:"sectionIndex,omitempty"`
}

type RiverRaceStandingList []RiverRaceStanding

type RiverRaceStanding struct {
	Rank         int           `json:"rank,omitempty"`
	TrophyChange int           `json:"trophyChange,omitempty"`
	Clan         RiverRaceClan `json:"clan,omitempty"`
}
