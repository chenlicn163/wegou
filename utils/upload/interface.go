package upload

import "net/http"

type Upload interface {
	UploadFile(r *http.Request) (string, error)
}
