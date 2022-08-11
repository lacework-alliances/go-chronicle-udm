package udm

// User metadata
type User struct {
	Email            string `json:"email,omitempty"`      // Repeated string
	EmployeeID       string `json:"employeeid,omitempty"` // Stores the human resources employee ID
	FirstName        string `json:"firstName,omitempty"`
	MiddleName       string `json:"middleName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	GroupIdentifiers string `json:"groupIdentifiers,omitempty"` //Stores the group ID(s) (a GUID, LDAP OID, or similar) associated with a user. Repeated string
	PhoneNumbers     string `json:"phoneNumbers,omitempty"`     // Repeated String
	Title            string `json:"title,omitempty"`            // Job Title
	UserDisplayName  string `json:"userDisplayName,omitempty"`
	UserID           string `json:"userid,omitempty"`
	WindowsSID       string `json:"windowsSid,omitempty"` // Stores the Microsoft Windows security identifier (SID) associated with a user.
}
