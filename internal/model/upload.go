package model

type FileUploadInput struct {
	FileName string
	TempPath string
}

type FileUploadOutput struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}
