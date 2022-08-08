package udm

import "time"

type Group struct {
	CreationTime    time.Time `json:"creation_time,omitempty"`
	EmailAddresses  string    `json:"email_addresses,omitempty"`
	DisplayName     string    `json:"display_name,omitempty"`
	ProductObjectID string    `json:"product_object_id,omitempty"` // Globally unique user object identifier for the product, such as an LDAP object identifier.
	WindowsSID      string    `json:"windows_sid,omitempty"`       // Microsoft Windows Security Identifier (SID) group attribute field.
}
