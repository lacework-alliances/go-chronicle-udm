package udm

const (
	ACCESS_POLICY      = "ACCESS_POLICY"
	BACKEND_SERVICE    = "BACKEND_SERVICE"
	CLOUD_ORGANIZATION = "CLOUD_ORGANIZATION"
	CLOUD_PROJECT      = "CLOUD_PROJECT"
	CLUSTER            = "CLUSTER"
	DATABASE           = "DATABASE"
	DATASET            = "DATASET"
	DEVICE             = "DEVICE"
	FIREWALL_RULE      = "FIREWALL_RULE"
	MAILBOX_FOLDER     = "MAILBOX_FOLDER"
	MUTEX              = "MUTEX"
	PIPE               = "PIP"
	SERVICE_ACOUNT     = "SERVICE_ACCOUNT"
	STORAGE_BUCKET     = "STORAGE_BUCKET"
	STORAGE_OBJECT     = "STORAGE_OBJECT"
	TABLE              = "TABLE"
	TASK               = "TASK"
	UNSPECIFIED        = "UNSPECIFIED"
	VIRTUAL_MACHINE    = "VIRTUAL_MACHINE"
	VPC_NETWORK        = "VPC_NETWORK"
)

type Resource struct {
	Attribute       *Attribute `json:"attribute,omitempty"`
	Name            string     `json:"name,omitempty"`
	Parent          string     `json:"parent,omitempty"`
	ProductObjectId string     `json:"product_object_id,omitempty"`
	ResourceSubtype string     `json:"resource_subtype,omitempty"`
	ResourceType    string     `json:"resource_type,omitempty"` // Enumerated constant
	Type            string     `json:"type,omitempty"`          // Enumerated and Depreciated
}
