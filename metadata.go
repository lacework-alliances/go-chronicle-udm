package udm

import "time"

const (
	// EVENT TYPES

	EMAIL_TRANSACTION = "EMAIL_TRANSACTION"

	FILE_CREATION     = "FILE_CREATION"
	FILE_DELETION     = "FILE_DELETION"
	FILE_MODIFICATION = "FILE_MODIFICATION"
	FILE_READ         = "FILE_READ"
	FILE_OPEN         = "FILE_OPEN"
	FILE_COPY         = "FILE_COPY"

	MUTEX_CREATION = "MUTEX_CREATION"

	NETWORK_CONNECTION = "NETWORK_CONNECTION"
	NETWORK_HTTP       = "NETWORK_HTTP"

	PROCESS_INJECTION            = "PROCESS_INJECTION"
	PROCESS_LAUNCH               = "PROCESS_LAUNCH"
	PROCESS_OPEN                 = "PROCESS_OPEN"
	PROCESS_TERMINATION          = "PROCESS_TERMINATION"
	PROCESS_UNCATEGORIZED        = "PROCESS_UNCATEGORIZED"
	PROCESS_MODULE_LOAD          = "PROCESS_MODULE_LOAD"
	PROCESS_PRIVILEGE_ESCALATION = "PROCESS_PRIVILEGE_ESCALATION"

	REGISTRY_CREATION     = "REGISTRY_CREATION"
	REGISTRY_MODIFICATION = "REGISTRY_MODIFIATION"
	REGISTRY_DELETION     = "REGISTRY_DELETION"

	SCAN_FILE         = "SCAN_FILE"
	SCAN_HOST         = "SCAN_HOST"
	SCAN_PROCESS      = "SCAN_PROCESS"
	SCAN_VULN_HOST    = "SCAN_VULN_HOST"
	SCAN_VULN_NETWORK = "SCAN_VULN_NETWORK"

	SCHEDULED_TASK_CREATION      = "SCHEDULED_TASK_CREATION"
	SCHEDULED_TASK_DELETION      = "SCHEDULED_TASK_DELETION"
	SCHEDULED_TASK_DISABLE       = "SCHEDULED_TASK_DISABLE"
	SCHEDULED_TASK_ENABLE        = "SCHEDULED_TASK_ENABLE"
	SCHEDULED_TASK_MODIFICATION  = "SCHEDULED_TASK_MODIFICATION"
	SCHEDULED_TASK_UNCATEGORIZED = "SCHEDULED_TASK_UNCATEGORIZED"

	SETTING_UNCATEGORIZED = "SETTING_UNCATEGORIZED"
	SETTING_CREATION      = "SETTING_CREATION"
	SETTING_MODIFICATION  = "SETTING_MODIFACTION"
	SETTING_DELETION      = "SETTING_DELETION"

	SERVICE_UNSPECIFIED = "SERVICE_UNSPECIFIED"
	SERVICE_CREATION    = "SERVICE_CREATION"
	SERVICE_DELETION    = "SERVICE_DELETION"
	SERVICE_START       = "SERVICE_START"
	SERVICE_STOP        = "SERVICE_STOP"

	STATUS_HEARTBEAT = "STATUS_HEARTBEAT"
	STATUS_STARTUP   = "STATUS_STARTUP"
	STATUS_SHUTDOWN  = "STATUS_SHUTDOWN"
	STATUS_UPDATE    = "STATUS_UPDATE"

	SYSTEM_AUDIT_LOG_UNCATEGORIZED = "SYSTEM_AUDIT_LOG_UNCATEGORIZED"
	SYSTEM_AUDIT_LOG_WIPE          = "SYSTEM_AUDIT_LOG_WIPE"

	USER_CHANGE_PASSWORD    = "USER_CHANGE_PASSWORD"
	USER_CHANGE_PERMISSIONS = "USER_CHANGE_PERMISSIONS"

	USER_CREATION = "USER_CREATION"
	USER_DELETION = "USER_DELETION"

	USER_LOGIN  = "USER_LOGIN"
	USER_LOGOUT = "USER_LOGOUT"

	USER_RESOURCE_ACCESS = "USER_RESOURCE_ACCESS"

	USER_RESOURCE_CREATION = "USER_RESOURCE_CREATION"
	USER_RESOURCE_DELETION = "USER_RESOURCE_DELETION"

	USER_RESOURCE_UPDATE_CONTENT = "USER_RESOURCE_UPDATE_CONTENT"

	USER_RESOURCE_UPDATE_PERMISSIONS = "USER_RESOURCE_UPDATE_PERMISSIONS"

	USER_UNCATEGORIZED = "USER_UNCATEGORIZED"

	// END EVENT TYPES

)

// Metadata - The event metadata section for UDM events stores general information about each event.
type Metadata struct {
	CollectedTimestamp time.Time `json:"collected_timestamp,omitempty"` // Collection time of the event RFC 3339
	Description        string    `json:"description,omitempty"`         // Human-readable description of the event 1024-byte max
	EventTimestamp     time.Time `json:"event_timestamp"`               // Event generation time RFC 3339
	EventType          string    `json:"event_type"`                    // Specifies the primary type of the event
	ID                 []byte    `json:"id,omitempty"`
	IngestedTimestamp  time.Time `json:"ingested_timestamp,omitempty"`
	IngestionLabels    Label     `json:"ingestion_labels,omitempty"`
	ProductEventType   string    `json:"product_event_type,omitempty"` // Short, descriptive, human-readable, and product-specific event name or type. 64-byte max
	ProductLogID       string    `json:"product_log_id,omitempty"`     // Encodes a vendor-specific event identifier to uniquely identify the event (a GUID). case-sensitive 256-byte max
	ProductName        string    `json:"product_name,omitempty"`       // Specifies the name of the product. 256-byte max
	ProductVersion     string    `json:"product_version,omitempty"`    // Specifies the version of the product. 32-byte max
	Tags               Tag       `json:"tags,omitempty"`
	URLBackToProduct   string    `json:"url_back_to_product,omitempty"` // URL linking to a relevant website where you can view more information - Valid RFC 3986 URL
	VendorName         string    `json:"vendor_name,omitempty"`         // Specifies the product vendor's name - 256-byte max
}
