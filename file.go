package udm

// File metadata
type File struct {
	FileMetadata string `json:"fileMetadata,omitempty"` // metadata associated with the file
	FullPath     string `json:"fullPath,omitempty"`     // full path location of the file
	MD5          string `json:"md5,omitempty"`          // MD5 hash - lowercase hexadecimal
	MimeType     string `json:"mimeType,omitempty"`     // Multipurpose Internet Mail Extensions (MIME) type for the file.
	SHA1         string `json:"sha1,omitempty"`         // SHA-1 hash value for the file
	SHA256       string `json:"sha256,omitempty"`       // SHA-256 hash value for the file
	Size         int64  `json:"size,omitempty"`         // 64 bit unsigned integer
}
