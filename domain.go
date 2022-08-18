package udm

import "time"

// Domain - https://cloud.google.com/chronicle/docs/reference/udm-field-list#domain
type Domain struct {
	Admin               *User       `json:"admin,omitempty"`
	AuditUpdateTime     *time.Time  `json:"audit_update_time,omitempty"`
	Billing             *User       `json:"billing,omitempty"`
	ContactEmail        string      `json:"contact_email,omitempty"`
	CreationTime        *time.Time  `json:"creation_time,omitempty"`
	ExpirationTime      *time.Time  `json:"expiration_time,omitempty"`
	FirstSeenTime       *time.Time  `json:"first_seen_time,omitempty"`
	IanaRegistrarId     int32       `json:"iana_registrar_id,omitempty"`
	LastSeenTime        *time.Time  `json:"last_seen_time,omitempty"`
	Name                string      `json:"name,omitempty"`
	NameServer          []string    `json:"name_server,omitempty"`
	Prevalence          *Prevalence `json:"prevalence,omitempty"`
	PrivateRegistration bool        `json:"private_registration,omitempty"`
	Registrant          *User       `json:"registrant,omitempty"`
	Registrar           string      `json:"registrar,omitempty"`
	RegistryDataRawText []byte      `json:"registry_data_raw_text,omitempty"`
	Status              string      `json:"status,omitempty"`
	Tech                *User       `json:"tech,omitempty"`
	UpdateTime          *time.Time  `json:"update_time,omitempty"`
	WhoisRecordRawText  []byte      `json:"whois_record_raw_text,omitempty"`
	WhoisServer         string      `json:"whois_server,omitempty"`
	Zone                *User       `json:"zone,omitempty"`
}
