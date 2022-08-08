package udm

// File metadata
type File struct {
	FileMetadata string `json:"file_metadata,omitempty"` // metadata associated with the file
	FullPath     string `json:"full_path,omitempty"`     // full path location of the file
	MD5          string `json:"md_5,omitempty"`          // MD5 hash - lowercase hexadecimal
	MimeType     string `json:"mime_type,omitempty"`     // Multipurpose Internet Mail Extensions (MIME) type for the file.
	SHA1         string `json:"sha_1,omitempty"`         // SHA-1 hash value for the file
	SHA256       string `json:"sha_256,omitempty"`       // SHA-256 hash value for the file
	Size         int64  `json:"size,omitempty"`          // 64 bit unsigned integer
}
