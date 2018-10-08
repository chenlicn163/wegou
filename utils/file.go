package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path"
)

func Upload(r *http.Request, UploadPath string) (string, error) {
	//上传文件
	formFile, header, err := r.FormFile("file")
	defer formFile.Close()
	if err != nil {
		return "", err
	}

	fileSuffix := path.Ext(path.Base(header.Filename))

	fileName := UniqueId() + fileSuffix
	fullFileName := UploadPath + "/" + fileName
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

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
