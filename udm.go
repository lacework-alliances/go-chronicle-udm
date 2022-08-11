package udm

type UDM struct {
	Metadata       Metadata       `json:"metadata"`
	About          Noun           `json:"about,omitempty"`
	Alert          Alert          `json:"alert,omitempty"`
	Authentication Authentication `json:"authentication,omitempty"`
	Dhcp           Dhcp           `json:"dhcp,omitempty"`
	Dns            Dns            `json:"dns,omitempty"`
	Email          Email          `json:"email,omitempty"`
	Extensions     Extensions     `json:"extensions,omitempty"`
	File           File           `json:"file,omitempty"`
	FTP            FTP            `json:"ftp,omitempty"`
	Group          Group          `json:"group,omitempty"`
	Http           Http           `json:"http,omitempty"`
	Intermediary   Noun           `json:"intermediary,omitempty"`
	Location       Location       `json:"location,omitempty"`
	Network        Network        `json:"network,omitempty"`
	Observer       Noun           `json:"observer,omitempty"`
	Principal      Noun           `json:"principal,omitempty"`
	Process        Process        `json:"process,omitempty"`
	Registry       Registry       `json:"registry,omitempty"`
	SecurityResult SecurityResult `json:"security_result,omitempty"`
	Src            Noun           `json:"src,omitempty"`
	Target         Noun           `json:"target,omitempty"`
	User           User           `json:"user,omitempty"`
	Vulnerability  Vulnerability  `json:"vulnerability,omitempty"`
}
