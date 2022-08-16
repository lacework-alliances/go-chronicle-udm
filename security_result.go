package udm

const (
	// ACTION
	ALLOW                  = "ALLOW"
	ALLOW_WITH_MODIFIATION = "ALLOW_WITH_MODIFICATION"
	BLOCK                  = "BLOCK"
	QUARANTINE             = "QUARANTINE"
	UNKNOWN_ACTION         = "UNKNOWN_ACTION"

	// CATEGORY
	ACL_VIOLATION               = "ACL_VIOLATION"
	AUTH_VIOLATION              = "AUTH_VIOLATION"
	DATA_AT_REST                = "DATA_AT_REST"
	DATA_DESTRUCTION            = "DATA_DESTRUCTION"
	DATA_EXFILTRATION           = "DATA_EXFILTRATION"
	EXPLOIT                     = "EXPLOIT"
	MAIL_PHISHING               = "MAIL_PHISHING"
	MAIL_SPAM                   = "MAIL_SPAM"
	MAIL_SPOOFING               = "MAIL_SPOOFING"
	NETWORK_CATEGORIZED_CONTENT = "NETWORK_CATEGORIZED_CONTENT"
	NETWORK_COMMAND_AND_CONTROL = "NETWORK_COMMAND_AND_CONTROL"
	NETWORK_DENIAL_OF_SERVICE   = "NETWORK_DENIAL_OF_SERVICE"
	NETWORK_MALICIOUS           = "NETWORK_MALICIOUS"
	NETWORK_SUSPICIOUS          = "NETWORK_SUSPICIOUS"
	NETWORK_RECON               = "NETWORK_RECON"
	POLICY_VIOLATION            = "POLICY_VIOLATION"
	SOFTWARE_MALICIOUS          = "SOFTWARE_MALICIOUS"
	SOFTWARE_PUA                = "SOFTWARE_PUA"
	SOFTWARE_SUSPICIOUS         = "SOFTWARE_SUSPICIOUS"
	UNKNOWN_CATEGORY            = "UNKNOWN_CATEGORY"

	// CONFIDENCE
	UNKNOWN_CONFIDENCE = "UNKNOWN_CONFIDENCE"
	LOW_CONFIDENCE     = "LOW_CONFIDENCE"
	MEDIUM_CONFIDENCE  = "MEDIUM_CONFIDENCE"
	HIGH_CONFIDENCE    = "HIGH_CONFIDENCE"

	// PRIORITY
	UNKNOWN_PRIORITY = "UNKNOWN_PRIORITY"
	LOW_PRIORITY     = "LOW_PRIORITY"
	MEDIUM_PRIORITY  = "MEDIUM_PRIORITY"
	HIGH_PRIORITY    = "HIGH_PRIORITY"

	// SEVERITY
	INFORMATIONAL = "INFORMATIONAL"
	ERROR         = "ERROR"
)

// SecurityResult metadata
type SecurityResult struct {
	About             string   `json:"about,omitempty"`             // String description
	Action            []string `json:"action,omitempty"`            // Enumerated
	ActionDetails     string   `json:"actionDetails,omitempty"`     // Vendor-provided details of the action taken as a result of the security incident. Examples: drop, block, decrypt, encrypt
	Category          string   `json:"category,omitempty"`          // Enumerated
	Confidence        string   `json:"confidence,omitempty"`        // Enumerated
	ConfidenceDetails string   `json:"confidenceDetails,omitempty"` // Additional detail with regards to the confidence of a security event as estimated by the product vendor.
	Priority          string   `json:"priority,omitempty"`          // Enumerated
	RuleID            string   `json:"ruleid,omitempty"`
	RuleName          string   `json:"ruleName,omitempty"`
	Severity          string   `json:"severity,omitempty"`        // Enumerated
	SeverityDetails   string   `json:"severityDetails,omitempty"` // Severity for a security event as estimated by the product vendor.
	ThreatName        string   `json:"threatName,omitempty"`
	URLBackToProduct  string   `json:"URLBackToProduct,omitempty"`
}
