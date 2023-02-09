package request

import "github.com/google/uuid"

type FileRequestModel struct {
	Filetype string
	Filesize int
	Buffer   []byte
}

type FileDownloadRequestModel struct {
	Filename uuid.UUID `json:"filename"`
	Filetype string    `json:"filetype"`
}
