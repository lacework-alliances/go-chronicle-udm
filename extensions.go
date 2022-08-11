package udm

// Extensions - Event types with first-class metadata that are not already categorized by the Chronicle UDM.
type Extensions struct {
	Auth        string `json:"auth,omitempty"` // Extension to the authentication metadata.
	AuthDetails string `json:"authDetails,omitempty"`
	Vulns       string `json:"vulns,omitempty"` // Extension to the vulnerability metadata.
}
