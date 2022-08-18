package udm

import "time"

// Artifact - https://cloud.google.com/chronicle/docs/reference/udm-field-list#artifact
type Artifact struct {
	FirstSeenTime *time.Time  `json:"first_seen_time,omitempty"`
	IP            string      `json:"ip,omitempty"`
	LastSeenTime  *time.Time  `json:"last_seen_time,omitempty"`
	Prevalence    *Prevalence `json:"prevalence,omitempty"`
}
