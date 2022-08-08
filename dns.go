package udm

type Dns struct {
	Authoritative      bool             `json:"authoritative,omitempty"` // is an authoritative DNS server?
	ID                 int32            `json:"id,omitempty"`            // DNS query identifier
	Response           bool             `json:"response,omitempty"`      // is event a DNS response?
	OpCode             int32            `json:"opcode,omitempty"`        //Stores the DNS OpCode used to specify the type of DNS query (standard, inverse, server status, etc.).
	RecursionAvailable bool             `json:"recursion_available,omitempty"`
	RecursionDesired   bool             `json:"recursion_desired,omitempty"`
	ResponseCode       int32            `json:"response_code,omitempty"` // Stores the DNS response code as defined by RFC 1035, Domain Names - Implementation and Specification.
	Truncated          bool             `json:"truncated,omitempty"`
	Questions          []Question       `json:"questions,omitempty"`  // Stores the domain protocol message questions
	Answers            []ResourceRecord `json:"answers,omitempty"`    // Stores the answer to the domain name query.
	Authority          []ResourceRecord `json:"authority,omitempty"`  // stores the domain name servers which verified the answer
	Additional         []ResourceRecord `json:"additional,omitempty"` // Stores the additional domain name servers that can be used to verify the answer to the domain.
}

type Question struct {
	Name  string `json:"name,omitempty"`  // domain name
	Class int32  `json:"class,omitempty"` // query class code
	Type  int32  `json:"type,omitempty"`  // query type code
}

type ResourceRecord struct {
	BinaryData []byte `json:"binary_data,omitempty"` // Stores the raw bytes of any non-UTF8 strings that might be included as part of a DNS response. This field must only be used if the response data returned by the DNS server contains non-UTF8 data. Otherwise, place the DNS response in the data field below. This type of information must be stored here rather than in ResourceRecord.data.
	Class      int32  `json:"class,omitempty"`       // resource record class
	Data       string `json:"data,omitempty"`        // Stores the payload or response to the DNS question for all responses encoded in UTF-8 format. For example, the data field could return the IP address of the machine that the domain name refers to. If the resource record is for a different type or class, it might contain another domain name (when one domain name is redirected to another domain name). Data must be stored just as it is in the DNS response.
	Name       string `json:"name,omitempty"`        // resource record owner name
	Ttl        int32  `json:"ttl,omitempty"`         // Stores the time interval for which the resource record can be cached before the source of the information should again be queried.
	Type       int32  `json:"type,omitempty"`        // stores the resource record type
}
