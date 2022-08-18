package udm

import "time"

const (
	// ASSET_TYPE
	IOT                      = "IOT"
	LAPTOP                   = "LAPTOP"
	MOBILE                   = "MOBILE"
	NETWORK_ATTACHED_STORAGE = "NETWORK_ATTACHED_STORAGE"
	PRINTER                  = "PRINTER"
	ROLE_UNSPECIFIED         = "ROLE_UNSPECIFIED"
	SCANNER                  = "SCANNER"
	SERVER                   = "SERVER"
	TAPE_LIBRARY             = "TAPE_LIBRARY"
	WORKSTATION              = "WORKSTATION"

	// ASSET DEPLOYMENT STATUS
	DECOMMISSIONED                = "DECOMMISSIONED"
	DEPLOYMENT_STATUS_UNSPECIFIED = "DEPLOYMENT_STATUS_UNSPECIFIED"
	PENDING_DECOMISSION           = "PENDING_DECOMISSION"
)

// Asset - https://cloud.google.com/chronicle/docs/reference/udm-field-list#asset
type Asset struct {
	AssetId              string            `json:"asset_id,omitempty"`
	Attribute            *Attribute        `json:"attribute,omitempty"`
	Category             string            `json:"category,omitempty"`
	CreationTime         *time.Time        `json:"creation_time,omitempty"`
	DeploymentStatus     string            `json:"deployment_status,omitempty"`
	FirstDiscoverTime    *time.Time        `json:"first_discover_time,omitempty"`
	FirstSeenTime        *time.Time        `json:"first_seen_time,omitempty"`
	Hardware             []*Hardware       `json:"hardware,omitempty"`
	Hostname             string            `json:"hostname,omitempty"`
	IP                   []string          `json:"ip,omitempty"`
	Labels               []*Label          `json:"labels,omitempty"`
	LastBootTime         *time.Time        `json:"last_boot_time,omitempty"`
	LastDiscoverTime     *time.Time        `json:"last_discover_time,omitempty"`
	Location             *Location         `json:"location,omitempty"`
	MAC                  []string          `json:"mac,omitempty"`
	NatIp                []string          `json:"nat_ip,omitempty"`
	NetworkDomain        string            `json:"network_domain,omitempty"`
	PlatformSoftware     *PlatformSoftware `json:"platform_software,omitempty"`
	ProductObjectId      string            `json:"product_object_id,omitempty"`
	Software             []*Software       `json:"software,omitempty"`
	SystemLastUpdateTime *time.Time        `json:"system_last_update_time,omitempty"`
	Type                 string            `json:"type,omitempty"` // enumerated constant ASSET_TYPE
	Vulnerabilities      []*Vulnerability  `json:"vulnerabilities,omitempty"`
}
