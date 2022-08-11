package udm

type Email struct {
	From    string `json:"from,omitempty"`
	ReplyTo string `json:"replyTo,omitempty"`
	To      string `json:"to,omitempty"`
	CC      string `json:"cc,omitempty"`
	BCC     string `json:"bcc,omitempty"`
	MailID  string `json:"mailid,omitempty"`
	Subject string `json:"subject,omitempty"`
}
