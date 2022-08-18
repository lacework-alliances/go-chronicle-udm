package udm

import "time"

// Certificate - https://cloud.google.com/chronicle/docs/reference/udm-field-list#certificate
type Certificate struct {
	Issuer    string     `json:"issuer,omitempty"`
	MD5       string     `json:"md5,omitempty"`
	NotAfter  *time.Time `json:"not_after,omitempty"`
	NotBefore *time.Time `json:"not_before,omitempty"`
	Serial    string     `json:"serial,omitempty"`
	SHA1      string     `json:"sha1,omitempty"`
	SHA256    string     `json:"sha256,omitempty"`
	Subject   string     `json:"subject,omitempty"`
	Version   string     `json:"version,omitempty"`
}
