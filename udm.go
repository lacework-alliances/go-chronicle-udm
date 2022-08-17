package udm

type UDM struct {
	About *Noun `json:"about,omitempty"` // Repeatable
	// Additional 	protobuf	`json:"additional"`
	Extensions     *Extensions       `json:"extensions,omitempty"`
	Intermediary   *Noun             `json:"intermediary,omitempty"` // Repeatable
	Metadata       Metadata          `json:"metadata"`
	Network        *Network          `json:"network,omitempty"`
	Observer       *Noun             `json:"observer,omitempty"`
	Principal      *Noun             `json:"principal,omitempty"`
	SecurityResult []*SecurityResult `json:"security_result,omitempty"` // Repeatable
	Src            *Noun             `json:"src,omitempty"`
	Target         *Noun             `json:"target,omitempty"`
	Vulnerability  *Vulnerability    `json:"vulnerability,omitempty"`
}
