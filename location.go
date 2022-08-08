package udm

type Location struct {
	City            string `json:"city,omitempty"`
	CountryOrRegion string `json:"country_or_region,omitempty"`
	Name            string `json:"name,omitempty"`
	State           string `json:"state,omitempty"`
}
