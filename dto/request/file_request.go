package request

import "github.com/google/uuid"

type FileGroupRequestModel struct {
	Filename string
	Data     []byte
}

type FileRequestModel struct {
	Filetype    string
	Filesize    int
	Buffer      []byte
	AssetTypeId uint16
}

type FileDownloadRequestModel struct {
	Filename uuid.UUID `json:"filename"`
	Filetype string    `json:"filetype"`
}
