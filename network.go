package udm

const (
	// APPLICATION PROTOCOL
	UNKNOWN_APPLICATION_PROTOCOL = "UNKNOWN_APPLICATION_PROTOCOL"
	QUIC                         = "QUIC"
	HTTP                         = "HTTP"
	HTTPS                        = "HTTPS"
	DNS                          = "DNS"
	DHCP                         = "DHCP"

	// DIRECTION
	UNKNOWN_DIRECTION = "UNKNOWN_DIRECTION"
	INBOUND           = "INBOUND"
	OUTBOUND          = "OUTBOUND"
	BROADCAST         = "BROADCAST"

	// IP PROTOCOL
	UNKNOWN_IP_PROTOCOL = "UNKNOWN_IP_PROTOCOL"
	EIGRP               = "EIGRP"
	ESP                 = "ESP"
	ETHERIP             = "ETHERIP"
	GRE                 = "GRE"
	ICMP                = "ICMP"
	IGMP                = "IGMP"
	IP6IN4              = "IP6IN4"
	PIM                 = "PIM"
	TCP                 = "TCP"
	UDP                 = "UDP"
	VRRP                = "VRRP"
)

type Network struct {
	ApplicationProtocol string `json:"application_protocol,omitempty"` // Enumerated
	Direction           string `json:"direction,omitempty"`            // Enumerated
	Email               string `json:"email,omitempty"`
	IPProtocol          string `json:"ip_protocol,omitempty"`      // Enumerated
	ReceivedBytes       int64  `json:"received_bytes,omitempty"`   // 64-bit unsigned integer
	SentBytes           int64  `json:"sent_bytes,omitempty"`       // 64-bit unsigned integer
	SessionDuration     int64  `json:"session_duration,omitempty"` // Stores the network session duration. Choosing nanoseconds (int64)
	SessionID           string `json:"session_id,omitempty"`
}
