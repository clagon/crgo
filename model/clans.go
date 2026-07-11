package model

type ClanMemberList []ClanMember

// ClanMemberPage is a paginated clan member response.
type ClanMemberPage struct {
	Items  ClanMemberList `json:"items,omitempty"`
	Paging Paging         `json:"paging,omitempty"`
}

type ClanMember struct {
	Arena             Arena          `json:"arena,omitempty"`
	ClanChestPoints   int            `json:"clanChestPoints,omitempty"`
	LastSeen          string         `json:"lastSeen,omitempty"`
	Tag               string         `json:"tag,omitempty"`
	Name              string         `json:"name,omitempty"`
	Role              ClanMemberRole `json:"role,omitempty"`
	ExpLevel          int            `json:"expLevel,omitempty"`
	Trophies          int            `json:"trophies,omitempty"`
	ClanRank          int            `json:"clanRank,omitempty"`
	PreviousClanRank  int            `json:"previousClanRank,omitempty"`
	Donations         int            `json:"donations,omitempty"`
	DonationsReceived int            `json:"donationsReceived,omitempty"`
}

type ClanMemberRole string

const (
	ClanMemberRoleNotMember ClanMemberRole = "NOT_MEMBER"
	ClanMemberRoleMember    ClanMemberRole = "MEMBER"
	ClanMemberRoleLeader    ClanMemberRole = "LEADER"
	ClanMemberRoleAdmin     ClanMemberRole = "ADMIN"
	ClanMemberRoleColeader  ClanMemberRole = "COLEADER"
)

type Clan struct {
	MemberList        ClanMemberList      `json:"memberList,omitempty"`
	Tag               string              `json:"tag,omitempty"`
	ClanChestStatus   ClanClanChestStatus `json:"clanChestStatus,omitempty"`
	ClanChestLevel    int                 `json:"clanChestLevel,omitempty"`
	ClanChestMaxLevel int                 `json:"clanChestMaxLevel,omitempty"`
	ClanScore         int                 `json:"clanScore,omitempty"`
	ClanWarTrophies   int                 `json:"clanWarTrophies,omitempty"`
	RequiredTrophies  int                 `json:"requiredTrophies,omitempty"`
	DonationsPerWeek  int                 `json:"donationsPerWeek,omitempty"`
	BadgeId           int                 `json:"badgeId,omitempty"`
	Name              string              `json:"name,omitempty"`
	Location          Location            `json:"location,omitempty"`
	Type              ClanType            `json:"type,omitempty"`
	Members           int                 `json:"members,omitempty"`
	Description       string              `json:"description,omitempty"`
	ClanChestPoints   int                 `json:"clanChestPoints,omitempty"`
	BadgeUrls         map[string]any      `json:"badgeUrls,omitempty"`
}

type ClanClanChestStatus string

const (
	ClanClanChestStatusInactive  ClanClanChestStatus = "INACTIVE"
	ClanClanChestStatusActive    ClanClanChestStatus = "ACTIVE"
	ClanClanChestStatusCompleted ClanClanChestStatus = "COMPLETED"
	ClanClanChestStatusUnknown   ClanClanChestStatus = "UNKNOWN"
)

type ClanType string

const (
	ClanTypeOpen       ClanType = "OPEN"
	ClanTypeInviteOnly ClanType = "INVITE_ONLY"
	ClanTypeClosed     ClanType = "CLOSED"
)

type CurrentClanWar struct {
	State             CurrentClanWarState    `json:"state,omitempty"`
	Clan              ClanWarClan            `json:"clan,omitempty"`
	Participants      ClanWarParticipantList `json:"participants,omitempty"`
	Clans             ClanWarClanList        `json:"clans,omitempty"`
	CollectionEndTime string                 `json:"collectionEndTime,omitempty"`
	WarEndTime        string                 `json:"warEndTime,omitempty"`
}

type CurrentClanWarState string

const (
	CurrentClanWarStateClanNotFound  CurrentClanWarState = "CLAN_NOT_FOUND"
	CurrentClanWarStateAccessDenied  CurrentClanWarState = "ACCESS_DENIED"
	CurrentClanWarStateNotInWar      CurrentClanWarState = "NOT_IN_WAR"
	CurrentClanWarStateCollectionDay CurrentClanWarState = "COLLECTION_DAY"
	CurrentClanWarStateMatchmaking   CurrentClanWarState = "MATCHMAKING"
	CurrentClanWarStateWarDay        CurrentClanWarState = "WAR_DAY"
	CurrentClanWarStateEnded         CurrentClanWarState = "ENDED"
)

type ClanWarClanList []ClanWarClan

type ClanWarClan struct {
	Tag           string `json:"tag,omitempty"`
	ClanScore     int    `json:"clanScore,omitempty"`
	Crowns        int    `json:"crowns,omitempty"`
	BadgeId       int    `json:"badgeId,omitempty"`
	Name          string `json:"name,omitempty"`
	Participants  int    `json:"participants,omitempty"`
	BattlesPlayed int    `json:"battlesPlayed,omitempty"`
	Wins          int    `json:"wins,omitempty"`
}

type ClanWarParticipantList []ClanWarParticipant

type ClanWarParticipant struct {
	Tag                        string `json:"tag,omitempty"`
	Name                       string `json:"name,omitempty"`
	CardsEarned                int    `json:"cardsEarned,omitempty"`
	BattlesPlayed              int    `json:"battlesPlayed,omitempty"`
	Wins                       int    `json:"wins,omitempty"`
	CollectionDayBattlesPlayed int    `json:"collectionDayBattlesPlayed,omitempty"`
	NumberOfBattles            int    `json:"numberOfBattles,omitempty"`
}

// ClanList is a paginated clan search response.
type ClanList struct {
	Items  []Clan `json:"items,omitempty"`
	Paging Paging `json:"paging,omitempty"`
}

type ClanWarLog []ClanWarLogEntry

type ClanWarLogEntry struct {
	Standings    ClanWarStandingList    `json:"standings,omitempty"`
	SeasonId     int                    `json:"seasonId,omitempty"`
	Participants ClanWarParticipantList `json:"participants,omitempty"`
	CreatedDate  string                 `json:"createdDate,omitempty"`
}

type ClanWarStandingList []ClanWarStanding

type ClanWarStanding struct {
	TrophyChange int         `json:"trophyChange,omitempty"`
	Clan         ClanWarClan `json:"clan,omitempty"`
}
