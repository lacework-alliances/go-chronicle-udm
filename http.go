package udm

type Http struct {
	Method       string `json:"method,omitempty"`       // GET, POST, HEAD, UPDATE, DELETE
	ReferralURL  string `json:"referral_url,omitempty"` // RFC 3986
	ResponseCode int32  `json:"response_code,omitempty"`
	UserAgent    string `json:"user_agent,omitempty"`
}
