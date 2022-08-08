package udm

// Registry metadata
type Registry struct {
	RegistryKey       string `json:"registry_key,omitempty"`        // Stores the registry key associated with an application or system component
	RegistryValueName string `json:"registry_value_name,omitempty"` // Stores the name of the registry value associated with an application or system component.
	RegistryValueData string `json:"registry_value_data,omitempty"` // Stores the data associated with a registry value.
}
