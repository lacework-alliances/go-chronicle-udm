package udm

const (
	// Platform Options
	LINUX            = "LINUX"
	MAC              = "MAC"
	WINDOWS          = "WINDOWS"
	UNKNOWN_PLATFORM = "UNKNOWN_PLATFORM"
)

// Noun - In this section, the word Noun is a overarching term used to represent the entities; principal, src, target,
// intermediary, observer, and about. These entities have common attributes, but represent different objects in an event.
type Noun struct {
	AssetID              string   `json:"assetid,omitempty"`              // Vendor-specific unique device identifier. Lacework.PolygraphDataPlatform.<ASSET>
	Email                string   `json:"email,omitempty"`                // email address
	File                 File     `json:"file,omitempty"`                 // file metadata
	Hostname             string   `json:"hostname,omitempty"`             // Client hostname or domain name field. Do not include if a URL is present. RFC 1123
	Platform             string   `json:"platform,omitempty"`             // Platform operating system.
	Process              Process  `json:"process,omitempty"`              // Detailed process metadata
	IP                   string   `json:"ip,omitempty"`                   // Single IP address associated with a network connection. Valid IPv4 or IPv6 address (RFC 5942) encoded in ASCII.
	Port                 string   `json:"port,omitempty"`                 // Source or destination network port number when a specific network connection is described within an event.
	MAC                  string   `json:"mac"`                            // One or more MAC addresses associated with a device. Valid MAC address (EUI-48) in ASCII.
	AdministrativeDomain string   `json:"administrativeDomain,omitempty"` // Domain which the device belongs to (for example, the Windows domain). Valid domain name string (128 characters maximum).
	Registry             Registry `json:"registry,omitempty"`             // Detailed registry metadata
	URL                  string   `json:"url,omitempty"`                  // Standard URL. RFC 3986
	User                 User     `json:"user,omitempty"`                 // Detailed user metadata.
}
