package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
