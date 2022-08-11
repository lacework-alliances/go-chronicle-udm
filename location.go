package udm

type Location struct {
	City            string `json:"city,omitempty"`
	CountryOrRegion string `json:"countryOrRegion,omitempty"`
	Name            string `json:"name,omitempty"`
	State           string `json:"state,omitempty"`
}
