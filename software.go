package udm

// Software - https://cloud.google.com/chronicle/docs/reference/udm-field-list#software
type Software struct {
	Name        string        `json:"name,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty"`
	Version     string        `json:"version,omitempty"`
}
