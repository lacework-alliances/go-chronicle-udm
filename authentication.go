package udm

const (
	// AUTH TYPE
	AUTHTYPE_UNSPECIFIED = "AUTHTYPE_UNSPECIFIED"
	MACHINE              = "MACHINE"
	PHYSICAL             = "PHYSICAL"
	SSO                  = "SSO"
	TACACS               = "TACACS"
	VPN                  = "VPN"

	// AUTHENTICATION STATUS
	UNKNOWN_AUTHENTICATION_STATUS = "UNKNOWN_AUTHENTICATION_STATUS"
	ACTIVE                        = "ACTIVE"
	SUSPENDED                     = "SUSPENDED"
	DELETED                       = "DELETED"
	NO_ACTIVE_CREDENTIALS         = "NO_ACTIVE_CREDENTIALS"

	// MECHANISM
	MECHANISM_UNSPECIFIED = "MECHANISM_UNSPECIFIED"
	BADGE_READER          = "BADGE_READER"
	BATCH                 = "BATCH"
	CACHED_INTERACTIVE    = "CACHED_INTERACTIVE"
	HARDWARE_KEY          = "HARDWARE_KEY"
	LOCAL                 = "LOCAL"
	MECHANISM_OTHER       = "MECHANISM_OTHER"
	NETWORK               = "NETWORK"
	NETWORK_CLEAR_TEXT    = "NETWORK_CLEAR_TEXT"
	NEW_CREDENTIALS       = "NEW_CREDENTIALS"
	OTP                   = "OTP"
	REMOTE                = "REMOTE"
	REMOTE_INTERACTIVE    = "REMOTE_INTERACTIVE"
	SERVICE               = "SERVICE"
	UNLOCK                = "UNLOCK"
	USERNAME_PASSWORD     = "USERNAME_PASSWORD"
)

type Authentication struct {
	AuthType             string   `json:"type,omitempty"`      // Type of system an authentication event is associated. Enumerated
	AuthenticationStatus string   `json:"status,omitempty"`    // Status of user or credential. Enumerated
	Mechanism            []string `json:"mechanism,omitempty"` // Enumerated
}
