package udm

type Email struct {
	From    string `json:"from,omitempty"`
	ReplyTo string `json:"reply_to,omitempty"`
	To      string `json:"to,omitempty"`
	CC      string `json:"cc,omitempty"`
	BCC     string `json:"bcc,omitempty"`
	MailID  string `json:"mail_id,omitempty"`
	Subject string `json:"subject,omitempty"`
}