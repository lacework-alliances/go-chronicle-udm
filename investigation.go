package udm

const (
	PRIORITY_CRITICAL    = "PRIORITY_CRITICAL"
	PRIORITY_HIGH        = "PRIORITY_HIGH"
	PRIORITY_MEDIUM      = "PRIORITY_MEDIUM"
	PRIORITY_INFO        = "PRIORITY_INFO"
	PRIORITY_LOW         = "PRIORITY_LOW"
	PRIORITY_UNSPECIFIED = "PRIORITY_UNSPECIFIED"

	REASON_MAINTENANCE   = "REASON_MAINTENANCE"
	REASON_MALICIOUS     = "REASON_MALICIOUS"
	REASON_NOT_MALICIOUS = "REASON_NOT_MALICIOUS"
	REASON_UNSPECIFIED   = "REASON_UNSPECIFIED"

	NOT_USEFUL             = "NOT_USEFUL"
	REPUTATION_UNSPECIFIED = "REPUTATION_UNSPECIFIED"
	USERFUL                = "USEFUL"

	CLOSED              = "CLOSED"
	NEW                 = "NEW"
	OPEN                = "OPEN"
	REVIEWED            = "REVIEWED"
	STATUS_UNSPECIFIFED = "STATUS_UNSPECIFIED"

	FALSE_POSITIVE      = "FALSE_POSITIVE"
	TRUE_POSITIVE       = "TRUE_POSITIVE"
	VERDICT_UNSPECIFIED = "VERDICT_UNSPECIFIED"
)

// Investigation - https://cloud.google.com/chronicle/docs/reference/udm-field-list#investigation
type Investigation struct {
	Comments      []string `json:"comments,omitempty"`
	Priority      string   `json:"priority,omitempty"` // enumerated constant
	Reason        string   `json:"reason,omitempty"`
	Reputation    string   `json:"reputation,omitempty"`
	RootCause     string   `json:"root_cause,omitempty"`
	SeverityScore uint32   `json:"severity_score,omitempty"`
	Status        string   `json:"status,omitempty"`
	Verdict       string   `json:"verdict,omitempty"`
}
