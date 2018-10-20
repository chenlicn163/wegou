package upload

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
)

//上传
type FileUpload struct {
	UploadPath string
}

//上传
func (up *FileUpload) UploadFile(r *http.Request) (string, error) {
	uploadPath := up.UploadPath
	//判断目录是否存在
	exists, _ := up.pathExists()
	if !exists {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			return "", errors.New("目录创建失败")
		}
	}

	//上传文件
	formFile, header, err := r.FormFile("file")
	defer formFile.Close()
	if err != nil {
		return "", err
	}

	fileSuffix := path.Ext(path.Base(header.Filename))

	fileName := up.uniqueId() + fileSuffix
	fullFileName := uploadPath + "/" + fileName
	destFile, err := os.Create(fullFileName)
	defer destFile.Close()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(destFile, formFile)
	if err != nil {
		return "", err
	}

	fileName = "/" + fullFileName
	return fileName, nil
}

//md5
func (up *FileUpload) getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成唯一标识
func (up *FileUpload) uniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return up.getMd5String(base64.URLEncoding.EncodeToString(b))
}

//判断路径是否存在
func (up *FileUpload) pathExists() (bool, error) {
	_, err := os.Stat(up.UploadPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
