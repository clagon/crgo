package model

type Location struct {
	LocalizedName string `json:"localizedName,omitempty"`
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IsCountry     bool   `json:"isCountry,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
}

type LocationList []Location
