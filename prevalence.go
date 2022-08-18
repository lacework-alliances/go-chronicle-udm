package udm

// Prevalence - https://cloud.google.com/chronicle/docs/reference/udm-field-list#prevalence
type Prevalence struct {
	DayCount             int32 `json:"day_count,omitempty"`
	DayMax               int32 `json:"day_max,omitempty"`
	RollingMax           int32 `json:"rolling_max,omitempty"`
	RollingMaxSubDomains int32 `json:"rolling_max_sub_domains,omitempty"`
}
