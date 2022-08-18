package udm

const (
	ADMINISTRATOR    = "ADMINISTRATOR"
	SERVICE_ACCOUNT  = "SERVICE_ACCOUNT"
	TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
)

type Role struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"` // Enumerated constant
}
