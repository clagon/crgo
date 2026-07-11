package model

type Item struct {
	IconUrls          map[string]any    `json:"iconUrls"`
	Name              JsonLocalizedName `json:"name"`
	Id                int               `json:"id"`
	Rarity            ItemRarity        `json:"rarity"`
	MaxLevel          int               `json:"maxLevel"`
	ElixirCost        *int              `json:"elixirCost"`
	MaxEvolutionLevel int               `json:"maxEvolutionLevel"`
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
	Id                int                   `json:"id"`
	Rarity            PlayerItemLevelRarity `json:"rarity"`
	Count             int                   `json:"count"`
	Level             int                   `json:"level"`
	StarLevel         int                   `json:"starLevel"`
	EvolutionLevel    int                   `json:"evolutionLevel"`
	Used              bool                  `json:"used"`
	Name              JsonLocalizedName     `json:"name"`
	MaxLevel          int                   `json:"maxLevel"`
	ElixirCost        *int                  `json:"elixirCost"`
	MaxEvolutionLevel int                   `json:"maxEvolutionLevel"`
	IconUrls          map[string]any        `json:"iconUrls"`
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
	Name              string `json:"name"`
	IndexHi           int    `json:"indexHi"`
	IndexLo           int    `json:"indexLo"`
	Available         bool   `json:"available"`
	DefaultOwned      bool   `json:"defaultOwned"`
	SfxFile           string `json:"sfxFile"`
	ScFile            string `json:"scFile"`
	AvailableForOffer bool   `json:"availableForOffer"`
	Exclusive         bool   `json:"exclusive"`
	DateAvailable     string `json:"dateAvailable"`
	HumanReadableName string `json:"humanReadableName"`
	Family            string `json:"family"`
}

type UpcomingChests struct {
	Items ChestList `json:"items"`
}

type ChestList []Chest

type Chest struct {
	Name     JsonLocalizedName `json:"name"`
	Index    int               `json:"index"`
	IconUrls map[string]any    `json:"iconUrls"`
}

type Items struct {
	Items        ItemList `json:"items"`
	SupportItems ItemList `json:"supportItems"`
}

type ItemList []Item
