package udm

// Extensions - Event types with first-class metadata that are not already categorized by the Chronicle UDM.
type Extensions struct {
	Auth  *Authentication `json:"auth,omitempty"`  // Extension to the authentication metadata.
	Vulns *Vulnerability  `json:"vulns,omitempty"` // Extension to the vulnerability metadata.
}
