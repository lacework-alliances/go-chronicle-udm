package udm

const (
	AWS   = "AWS"
	AZURE = "AZURE"
	GCP   = "GCP"
	// LINUX = "LINUX"
	// MAC = "MAC"
	// UNKNOWN_PLATFORM = "UNKNOWN_PLATFORM"
	// WINDOWS = "WINDOWS"
)

// PlatformSoftware - https://cloud.google.com/chronicle/docs/reference/udm-field-list#platformsoftware
type PlatformSoftware struct {
	Platform           string `json:"platform,omitempty"` // enumerated constant
	PlatformPatchLevel string `json:"platform_patch_level,omitempty"`
	PlatformVersion    string `json:"platform_version,omitempty"`
}
