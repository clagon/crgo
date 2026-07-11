package model

type Location struct {
	LocalizedName string `json:"localizedName"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}

// LocationList is a paginated location response.
type LocationList struct {
	Items  []Location `json:"items"`
	Paging Paging     `json:"paging"`
}
