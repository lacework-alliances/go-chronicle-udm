package udm

import "time"

type Group struct {
	CreationTime    time.Time `json:"creationTime,omitempty"`
	EmailAddresses  string    `json:"emailAddresses,omitempty"`
	DisplayName     string    `json:"displayName,omitempty"`
	ProductObjectID string    `json:"productObjectid,omitempty"` // Globally unique user object identifier for the product, such as an LDAP object identifier.
	WindowsSID      string    `json:"windowsSid,omitempty"`      // Microsoft Windows Security Identifier (SID) group attribute field.
}
