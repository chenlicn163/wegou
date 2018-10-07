package engine

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"wegou/wechat/service/server"
)

//查询永久素材
func ListMaterialServe(w http.ResponseWriter, r *http.Request) {

	materials := server.GetMaterial(r)
	DataJson := DataJson{
		Code:    "0",
		Message: "success",
		Data:    materials,
	}

	jsonStr, _ := json.Marshal(DataJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)

	/*wechatConfig := GetWechatConfig()
	srv := core.NewDefaultAccessTokenServer(wechatConfig.AppId, wechatConfig.AppSecret, nil)
	clt := core.NewClient(srv, nil)
	r.ParseForm()
	materialType := r.Form.Get("type")
	page := r.Form.Get("page")
	if materialType != "" {
		rslt := server.BatchFetch(clt, materialType, page)
		fmt.Println(rslt)
		//materialCount := GetMaterialCount(clt)
	}*/

}

//删除永久素材
func DeleteMaterialServe(w http.ResponseWriter, r *http.Request) {

	flag := server.DelMaterial(r)

	StatusJson := StatusJson{}
	if !flag {
		StatusJson.Code = "20001"
		StatusJson.Message = "删除失败"
	} else {
		StatusJson.Code = "0"
		StatusJson.Message = "success"
	}

	jsonStr, _ := json.Marshal(StatusJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonStr))
}

//添加永久素材
func AddMaterialServe(w http.ResponseWriter, r *http.Request) {

	flag := server.AddMaterial(r)

	StatusJson := StatusJson{}
	if !flag {
		StatusJson.Code = "20002"
		StatusJson.Message = "新增永久图片素材失败"
	} else {
		StatusJson.Code = "0"
		StatusJson.Message = "success"
	}

	jsonStr, _ := json.Marshal(StatusJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonStr))
}

//test 上传文件测试
func addFileServe(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("file", "test.jpg")
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	srcFile, err := os.Open("test.jpg")
	defer srcFile.Close()
	if err != nil {
		log.Fatalf("Open source file failed: %s\n", err)
	}
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file failed: %s\n", err)
	}

	contentType := writer.FormDataContentType()
	writer.Close()
	resp, err := http.Post("http://127.0.0.1:8090/material", contentType, buf)

	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)

}
