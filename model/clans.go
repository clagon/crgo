package model

type ClanMemberList []ClanMember

// ClanMemberPage is a paginated clan member response.
type ClanMemberPage struct {
	Items  ClanMemberList `json:"items"`
	Paging Paging         `json:"paging"`
}

type ClanMember struct {
	Arena             Arena          `json:"arena"`
	ClanChestPoints   int            `json:"clanChestPoints"`
	LastSeen          string         `json:"lastSeen"`
	Tag               string         `json:"tag"`
	Name              string         `json:"name"`
	Role              ClanMemberRole `json:"role"`
	ExpLevel          int            `json:"expLevel"`
	Trophies          int            `json:"trophies"`
	ClanRank          int            `json:"clanRank"`
	PreviousClanRank  int            `json:"previousClanRank"`
	Donations         int            `json:"donations"`
	DonationsReceived int            `json:"donationsReceived"`
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
	MemberList        ClanMemberList      `json:"memberList"`
	Tag               string              `json:"tag"`
	ClanChestStatus   ClanClanChestStatus `json:"clanChestStatus"`
	ClanChestLevel    int                 `json:"clanChestLevel"`
	ClanChestMaxLevel int                 `json:"clanChestMaxLevel"`
	ClanScore         int                 `json:"clanScore"`
	ClanWarTrophies   int                 `json:"clanWarTrophies"`
	RequiredTrophies  int                 `json:"requiredTrophies"`
	DonationsPerWeek  int                 `json:"donationsPerWeek"`
	BadgeId           int                 `json:"badgeId"`
	Name              string              `json:"name"`
	Location          Location            `json:"location"`
	Type              ClanType            `json:"type"`
	Members           int                 `json:"members"`
	Description       string              `json:"description"`
	ClanChestPoints   int                 `json:"clanChestPoints"`
	BadgeUrls         map[string]any      `json:"badgeUrls"`
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
	State             CurrentClanWarState    `json:"state"`
	Clan              ClanWarClan            `json:"clan"`
	Participants      ClanWarParticipantList `json:"participants"`
	Clans             ClanWarClanList        `json:"clans"`
	CollectionEndTime string                 `json:"collectionEndTime"`
	WarEndTime        string                 `json:"warEndTime"`
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
	Tag           string `json:"tag"`
	ClanScore     int    `json:"clanScore"`
	Crowns        int    `json:"crowns"`
	BadgeId       int    `json:"badgeId"`
	Name          string `json:"name"`
	Participants  int    `json:"participants"`
	BattlesPlayed int    `json:"battlesPlayed"`
	Wins          int    `json:"wins"`
}

type ClanWarParticipantList []ClanWarParticipant

type ClanWarParticipant struct {
	Tag                        string `json:"tag"`
	Name                       string `json:"name"`
	CardsEarned                int    `json:"cardsEarned"`
	BattlesPlayed              int    `json:"battlesPlayed"`
	Wins                       int    `json:"wins"`
	CollectionDayBattlesPlayed int    `json:"collectionDayBattlesPlayed"`
	NumberOfBattles            int    `json:"numberOfBattles"`
}

// ClanList is a paginated clan search response.
type ClanList struct {
	Items  []Clan `json:"items"`
	Paging Paging `json:"paging"`
}

type ClanWarLog []ClanWarLogEntry

type ClanWarLogEntry struct {
	Standings    ClanWarStandingList    `json:"standings"`
	SeasonId     int                    `json:"seasonId"`
	Participants ClanWarParticipantList `json:"participants"`
	CreatedDate  string                 `json:"createdDate"`
}

type ClanWarStandingList []ClanWarStanding

type ClanWarStanding struct {
	TrophyChange int         `json:"trophyChange"`
	Clan         ClanWarClan `json:"clan"`
}
