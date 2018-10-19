package utils

import (
	"net/http"
	"wegou/config"
	"wegou/utils/upload"
)

type Upload interface {
	UploadFile(r *http.Request) (string, error)
}

func GetUpload(uploadPath string) (up Upload) {
	toolsConfig := config.GetToolsConfig()
	switch toolsConfig.Upload {
	case "file":
		up = &upload.FileUpload{UploadPath: uploadPath}
	case "qiniu":

	}
	return up
}
