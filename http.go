package udm

type Http struct {
	Method       string `json:"method,omitempty"`      // GET, POST, HEAD, UPDATE, DELETE
	ReferralURL  string `json:"referralUrl,omitempty"` // RFC 3986
	ResponseCode int32  `json:"responseCode,omitempty"`
	UserAgent    string `json:"userAgent,omitempty"`
}
