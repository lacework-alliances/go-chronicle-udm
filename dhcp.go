package udm

const (
	// OPCODE
	UNKNOWN_OPCODE = "UNKNOWN_OPCODE"
	BOOTREQUEST    = "BOOTREQUEST"
	BOOTREPLY      = "BOOTREPLY"

	// TYPE
	UNKNOWN_MESSAGE_TYPE = "UNKNOWN_MESSAGE_TYPE"
	DISCOVER             = "DISCOVER"
	OFFER                = "OFFER"
	REQUEST              = "REQUEST"
	DECLINE              = "DECLINE"
	ACK                  = "ACK"
	NAK                  = "NAK"
	RELEASE              = "RELEASE"
	INFORM               = "INFORM"
	WIN_DELETED          = "WIN_DELETED"
	WIN_EXPIRED          = "WIN_EXPIRED"
)

type Dhcp struct {
	ClientHostname   string `json:"clientHostname,omitempty"`
	ClientIdentifier []byte `json:"clientIdentifier,omitempty"`
	File             string `json:"file,omitempty"` // file name for the boot image
	Flags            int32  `json:"flags,omitempty"`
	Hlen             int32  `json:"hlen,omitempty"`  // hardware address length
	Hops             int32  `json:"hops,omitempty"`  // HDCP hop count
	HType            int32  `json:"htype,omitempty"` // hardware address type
	LeaseTimeSeconds int32  `json:"leaseTimeSeconds,omitempty"`
	OpCode           string `json:"opcode,omitempty"`           // Enumerated
	RequestedAddress string `json:"requestedAddress,omitempty"` // valid ipv4 or ipv6 address encoded in ASCII
	Seconds          int32  `json:"seconds,omitempty"`          // seconds elapses since the client began the acquisition/renewal process
	SName            string `json:"sname,omitempty"`            // name of the server with the client requested boot from
	TransactionID    int32  `json:"transactionid"`
	Type             string `json:"type,omitempty"`   // Enumerated
	ChAddr           string `json:"chaddr,omitempty"` // ip address for client hardware. valid ipv4 or ipv6 address
	CiAddr           string `json:"ciaddr,omitempty"` // ip address for the client
	GiAddr           string `json:"giaddr,omitempty"` // ip address for the relay agent
	SiAddr           string `json:"siaddr,omitempty"` // ip address for the next bootstrap server
	YiAddr           string `json:"yiaddr,omitempty"` // Your ip address
}

type Option struct {
	Code int32  `json:"code,omitempty"` // dhcp option code
	Data []byte `json:"data,omitempty"` // dhcp option data
}
