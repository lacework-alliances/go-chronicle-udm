package udm

// Hardware - https://cloud.google.com/chronicle/docs/reference/udm-field-list#hardware
type Hardware struct {
	CpuClockSpeed    uint64 `json:"cpu_clock_speed,omitempty"`
	CpuMaxClockSpeed uint64 `json:"cpu_max_clock_speed,omitempty"`
	CpuModel         string `json:"cpu_model,omitempty"`
	CpuNumberCores   uint64 `json:"cpu_number_cores,omitempty"`
	CpuPlatform      string `json:"cpu_platform,omitempty"`
	Manufacturer     string `json:"manufacturer,omitempty"`
	Model            string `json:"model,omitempty"`
	RAM              uint64 `json:"ram,omitempty"`
	SerialNumber     string `json:"serial_number,omitempty"`
}
