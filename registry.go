package udm

// Registry metadata
type Registry struct {
	RegistryKey       string `json:"registryKey,omitempty"`       // Stores the registry key associated with an application or system component
	RegistryValueName string `json:"registryValueName,omitempty"` // Stores the name of the registry value associated with an application or system component.
	RegistryValueData string `json:"registryValueData,omitempty"` // Stores the data associated with a registry value.
}
