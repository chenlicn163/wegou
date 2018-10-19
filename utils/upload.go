package utils

import (
	"wegou/config"
	"wegou/utils/upload"
)

func GetUpload(uploadPath string) (up upload.Upload) {
	toolsConfig := config.GetToolsConfig()
	switch toolsConfig.Upload {
	case "file":
		up = &upload.FileUpload{UploadPath: uploadPath}
	case "qiniu":

	}
	return up
}
