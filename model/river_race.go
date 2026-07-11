package model

type CurrentRiverRace struct {
	State             CurrentRiverRaceState      `json:"state"`
	Clan              RiverRaceClan              `json:"clan"`
	Clans             RiverRaceClanList          `json:"clans"`
	CollectionEndTime string                     `json:"collectionEndTime"`
	WarEndTime        string                     `json:"warEndTime"`
	SectionIndex      int                        `json:"sectionIndex"`
	PeriodIndex       int                        `json:"periodIndex"`
	PeriodType        CurrentRiverRacePeriodType `json:"periodType"`
	PeriodLogs        PeriodLogList              `json:"periodLogs"`
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
	Items       PeriodLogEntryList `json:"items"`
	PeriodIndex int                `json:"periodIndex"`
}

type PeriodLogEntryList []PeriodLogEntry

type PeriodLogEntry struct {
	Clan                       PeriodLogEntryClan `json:"clan"`
	PointsEarned               int                `json:"pointsEarned"`
	ProgressStartOfDay         int                `json:"progressStartOfDay"`
	ProgressEndOfDay           int                `json:"progressEndOfDay"`
	EndOfDayRank               int                `json:"endOfDayRank"`
	ProgressEarned             int                `json:"progressEarned"`
	NumOfDefensesRemaining     int                `json:"numOfDefensesRemaining"`
	ProgressEarnedFromDefenses int                `json:"progressEarnedFromDefenses"`
}

type PeriodLogEntryClan struct {
	Tag string `json:"tag"`
}

type RiverRaceClanList []RiverRaceClan

type RiverRaceClan struct {
	Tag          string                   `json:"tag"`
	ClanScore    int                      `json:"clanScore"`
	BadgeId      int                      `json:"badgeId"`
	Name         string                   `json:"name"`
	Fame         int                      `json:"fame"`
	RepairPoints int                      `json:"repairPoints"`
	FinishTime   string                   `json:"finishTime"`
	Participants RiverRaceParticipantList `json:"participants"`
	PeriodPoints int                      `json:"periodPoints"`
}

type RiverRaceParticipantList []RiverRaceParticipant

type RiverRaceParticipant struct {
	Tag            string `json:"tag"`
	Name           string `json:"name"`
	Fame           int    `json:"fame"`
	RepairPoints   int    `json:"repairPoints"`
	BoatAttacks    int    `json:"boatAttacks"`
	DecksUsed      int    `json:"decksUsed"`
	DecksUsedToday int    `json:"decksUsedToday"`
}

// RiverRaceLog is a paginated river race log response.
type RiverRaceLog struct {
	Items  []RiverRaceLogEntry `json:"items"`
	Paging Paging              `json:"paging"`
}

type RiverRaceLogEntry struct {
	Standings    RiverRaceStandingList `json:"standings"`
	SeasonId     int                   `json:"seasonId"`
	CreatedDate  string                `json:"createdDate"`
	SectionIndex int                   `json:"sectionIndex"`
}

type RiverRaceStandingList []RiverRaceStanding

type RiverRaceStanding struct {
	Rank         int           `json:"rank"`
	TrophyChange int           `json:"trophyChange"`
	Clan         RiverRaceClan `json:"clan"`
}
