package udm

// User metadata
type User struct {
	Email            string `json:"email,omitempty"`       // Repeated string
	EmployeeID       string `json:"employee_id,omitempty"` // Stores the human resources employee ID
	FirstName        string `json:"first_name,omitempty"`
	MiddleName       string `json:"middle_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	GroupIdentifiers string `json:"group_identifiers,omitempty"` //Stores the group ID(s) (a GUID, LDAP OID, or similar) associated with a user. Repeated string
	PhoneNumbers     string `json:"phone_numbers,omitempty"`     // Repeated String
	Title            string `json:"title,omitempty"`             // Job Title
	UserDisplayName  string `json:"user_display_name,omitempty"`
	UserID           string `json:"user_id,omitempty"`
	WindowsSID       string `json:"windows_sid,omitempty"` // Stores the Microsoft Windows security identifier (SID) associated with a user.
}
