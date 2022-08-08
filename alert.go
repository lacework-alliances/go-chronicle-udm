package udm

// Alert metadata
type Alert struct {
	IsSignificant bool `json:"is_significant,omitempty"` // display alert in Enterprise Insights
	IsAlert       bool `json:"is_alert,omitempty"`       // is event an alert?
}
