package model

type Item struct {
	IconUrls          map[string]any    `json:"iconUrls,omitempty"`
	Name              JsonLocalizedName `json:"name,omitempty"`
	Id                int               `json:"id,omitempty"`
	Rarity            ItemRarity        `json:"rarity,omitempty"`
	MaxLevel          int               `json:"maxLevel,omitempty"`
	ElixirCost        int               `json:"elixirCost,omitempty"`
	MaxEvolutionLevel int               `json:"maxEvolutionLevel,omitempty"`
}

type ItemRarity string

const (
	ItemRarityCommon    ItemRarity = "COMMON"
	ItemRarityRare      ItemRarity = "RARE"
	ItemRarityEpic      ItemRarity = "EPIC"
	ItemRarityLegendary ItemRarity = "LEGENDARY"
	ItemRarityChampion  ItemRarity = "CHAMPION"
)

type PlayerItemLevelList []PlayerItemLevel

type PlayerItemLevel struct {
	Id                int                   `json:"id,omitempty"`
	Rarity            PlayerItemLevelRarity `json:"rarity,omitempty"`
	Count             int                   `json:"count,omitempty"`
	Level             int                   `json:"level,omitempty"`
	StarLevel         int                   `json:"starLevel,omitempty"`
	EvolutionLevel    int                   `json:"evolutionLevel,omitempty"`
	Used              bool                  `json:"used,omitempty"`
	Name              JsonLocalizedName     `json:"name,omitempty"`
	MaxLevel          int                   `json:"maxLevel,omitempty"`
	ElixirCost        int                   `json:"elixirCost,omitempty"`
	MaxEvolutionLevel int                   `json:"maxEvolutionLevel,omitempty"`
	IconUrls          map[string]any        `json:"iconUrls,omitempty"`
}

type PlayerItemLevelRarity string

const (
	PlayerItemLevelRarityCommon    PlayerItemLevelRarity = "COMMON"
	PlayerItemLevelRarityRare      PlayerItemLevelRarity = "RARE"
	PlayerItemLevelRarityEpic      PlayerItemLevelRarity = "EPIC"
	PlayerItemLevelRarityLegendary PlayerItemLevelRarity = "LEGENDARY"
	PlayerItemLevelRarityChampion  PlayerItemLevelRarity = "CHAMPION"
)

type EmoteList []Emote

type Emote struct {
	Name              string `json:"name,omitempty"`
	IndexHi           int    `json:"indexHi,omitempty"`
	IndexLo           int    `json:"indexLo,omitempty"`
	Available         bool   `json:"available,omitempty"`
	DefaultOwned      bool   `json:"defaultOwned,omitempty"`
	SfxFile           string `json:"sfxFile,omitempty"`
	ScFile            string `json:"scFile,omitempty"`
	AvailableForOffer bool   `json:"availableForOffer,omitempty"`
	Exclusive         bool   `json:"exclusive,omitempty"`
	DateAvailable     string `json:"dateAvailable,omitempty"`
	HumanReadableName string `json:"humanReadableName,omitempty"`
	Family            string `json:"family,omitempty"`
}

type UpcomingChests struct {
	Items ChestList `json:"items,omitempty"`
}

type ChestList []Chest

type Chest struct {
	Name     JsonLocalizedName `json:"name,omitempty"`
	Index    int               `json:"index,omitempty"`
	IconUrls map[string]any    `json:"iconUrls,omitempty"`
}

type Items struct {
	Items        ItemList `json:"items,omitempty"`
	SupportItems ItemList `json:"supportItems,omitempty"`
}

type ItemList []Item
