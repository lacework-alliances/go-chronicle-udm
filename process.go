package udm

// Process metadata
type Process struct {
	CommandLine                           string        `json:"command_line,omitempty"`                               // Command line string
	ProductSpecificProcessID              string        `json:"product_specific_process_id,omitempty"`                // Stores the product specific process ID.
	ParentProcessProductSpecificProcessID string        `json:"parent_process_product_specific_process_id,omitempty"` // Stores the product specific process ID for the parent process.
	File                                  string        `json:"file,omitempty"`                                       // file name of the file used by the process
	ParentProcess                         ParentProcess `json:"parent_process,omitempty"`                             // details of the parent process
	PID                                   string        `json:"pid,omitempty"`                                        // Process ID
}

type ParentProcess struct {
	CommandLine              string `json:"command_line,omitempty"`                // Command line string
	ProductSpecificProcessID string `json:"product_specific_process_id,omitempty"` // Stores the product specific process ID.
	File                     string `json:"file,omitempty"`                        // file name of the file used by the process
	PID                      string `json:"pid,omitempty"`                         // process id
}
