package model

type Location struct {
	LocalizedName string `json:"localizedName,omitempty"`
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IsCountry     bool   `json:"isCountry,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
}

// LocationList is a paginated location response.
type LocationList struct {
	Items  []Location `json:"items,omitempty"`
	Paging Paging     `json:"paging,omitempty"`
}
