package udm

const (
	AMAZON_WEB_SERVICES           = "AMAZON_WEB_SERVICES"
	GOOGLE_CLOUD_PLATFORM         = "GOOGLE_CLOUD_PLATFORM"
	MICROSOFT_AZURE               = "MICROSOFT_AZURE"
	UNSPECIFIED_CLOUD_ENVIRONMENT = "UNSPECIFIED_CLOUD_ENVIRONMENT"
)

type Cloud struct {
	AvailabilityZone string    `json:"availability_zone,omitempty"`
	Environment      string    `json:"environment,omitempty"` // enumerated constant
	Project          *Resource `json:"project,omitempty"`
	VPC              *Resource `json:"vpc,omitempty"`
}
