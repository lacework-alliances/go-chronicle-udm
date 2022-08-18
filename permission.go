package udm

const (
	ADMIN_READ              = "ADMIN_READ"
	ADMIN_WRITE             = "ADMIN_WRITE"
	DATA_READ               = "DATA_READ"
	DATA_WRITE              = "DATA_WRITE"
	UNKNOWN_PERMISSION_TYPE = "UNKNOWN_PERMISSION_TYPE"
)

type Permission struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"` // Enumerated constant
}
