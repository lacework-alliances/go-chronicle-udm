package udm

const (
	// Platform Options
	LINUX            = "LINUX"
	MAC              = "MAC"
	WINDOWS          = "WINDOWS"
	UNKNOWN_PLATFORM = "UNKNOWN_PLATFORM"
)

// Noun - https://cloud.google.com/chronicle/docs/reference/udm-field-list#noun
type Noun struct {
	AdministrativeDomain string         `json:"administrative_domain,omitempty"` // Domain which the device belongs to (for example, the Windows domain). Valid domain name string (128 characters maximum).
	Application          string         `json:"application,omitempty"`           // The name of an application or service. Some SSO solutions only capture the name of a target application such as "Atlassian" or "Chronicle".
	Artifact             *Artifact      `json:"artifact,omitempty"`              // Information about an artifact.
	Asset                *Asset         `json:"asset,omitempty"`                 // Information about an asset
	AssetID              string         `json:"asset_id,omitempty"`              // Vendor-specific unique device identifier. Lacework.PolygraphDataPlatform.<ASSET>
	Domain               *Domain        `json:"domain,omitempty"`                // Information about the domain
	Email                string         `json:"email,omitempty"`                 // email address
	File                 *File          `json:"file,omitempty"`                  // file metadata
	Group                *Group         `json:"group,omitempty"`                 // group data
	Hostname             string         `json:"hostname,omitempty"`              // Client hostname or domain name field. Do not include if a URL is present. RFC 1123
	Investigation        *Investigation `json:"investigation,omitempty"`
	IP                   []string       `json:"ip,omitempty"` // Single IP address associated with a network connection. Valid IPv4 or IPv6 address (RFC 5942) encoded in ASCII.
	Labels               []*Label       `json:"labels,omitempty"`
	Location             *Location      `json:"location,omitempty"`
	MAC                  []string       `json:"mac,omitempty"` // One or more MAC addresses associated with a device. Valid MAC address (EUI-48) in ASCII.
	Namespace            string         `json:"namespace,omitempty"`
	NatIp                []string       `json:"nat_ip,omitempty"`
	NatPort              int32          `json:"nat_port,omitempty"`
	ObjectReference      string         `json:"object_reference,omitempty"`
	Platform             string         `json:"platform,omitempty"` // Platform operating system.
	PlatformPatchLevel   string         `json:"platform_patch_level,omitempty"`
	PlatformVersion      string         `json:"platform_version,omitempty"`
	Port                 string         `json:"port,omitempty"`                  // Source or destination network port number when a specific network connection is described within an event.
	Process              *Process       `json:"process,omitempty"`               // Detailed process metadata
	ProcessAncestors     []*Process     `json:"process_ancestors,omitempty"`     // Information about the process's ancestors ordered from immediate ancestor (parent process) to root. Note: process_ancestors is only populated when data is exported to BigQuery since recursive fields (e.g. process.parent_process) are not supported by BigQuery.
	Registry             *Registry      `json:"registry,omitempty"`              // Detailed registry metadata
	Resource             *Resource      `json:"resource,omitempty"`              // Information about the resource (e.g. scheduled task, calendar entry). This field should not be used for files, registry, processes, etc. since these objects are already part of Noun.
	ResourceAncestors    []*Resource    `json:"resource_ancestors,omitempty"`    // Information about the resource's ancestors ordered from immediate ancestor (starting with parent resource).
	URL                  string         `json:"url,omitempty"`                   // Standard URL. RFC 3986
	User                 *User          `json:"user,omitempty"`                  // Detailed user metadata.
	UserManagementChain  []*User        `json:"user_management_chain,omitempty"` // Information about the user's management chain (reporting hierarchy). Note: user_management_chain is only populated when data is exported to BigQuery since recursive fields (e.g. user.managers) are not supported by BigQuery.
}
