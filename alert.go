package udm

// Alert metadata
type Alert struct {
	IsSignificant bool `json:"isSignificant,omitempty"` // display alert in Enterprise Insights
	IsAlert       bool `json:"isAlert,omitempty"`       // is event an alert?
}
