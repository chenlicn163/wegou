package admin

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

//查询永久素材
func ListMaterialServe(c *gin.Context) {

	materials := server.GetMaterial(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "success",
		"data":    materials,
	})

	/*c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(jsonStr)*/

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
func DeleteMaterialServe(c *gin.Context) {

	flag := server.DelMaterial(c)

	StatusJson := StatusJson{}
	if !flag {
		StatusJson.Code = "20001"
		StatusJson.Message = "删除失败"
	} else {
		StatusJson.Code = "0"
		StatusJson.Message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusJson.Code,
		"message": StatusJson.Message,
	})
}

//添加永久素材
func AddMaterialServe(c *gin.Context) {

	flag := server.AddMaterial(c)

	StatusJson := StatusJson{}
	if !flag {
		StatusJson.Code = "20002"
		StatusJson.Message = "新增永久图片素材失败"
	} else {
		StatusJson.Code = "0"
		StatusJson.Message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusJson.Code,
		"message": StatusJson.Message,
	})
}

//test 上传文件测试
func AddFileServe(c *gin.Context) {
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
	writer.WriteField("title", "test-title-wechat")
	writer.WriteField("author", "test-author-wechat")
	writer.WriteField("digest", "test-digest-wechat")
	writer.WriteField("content", "test-content-wechat")
	writer.WriteField("show_cover_pic", "0")
	writer.WriteField("material_type", "1")
	writer.WriteField("account_id", "0")
	writer.WriteField("status", "2")
	writer.WriteField("source_type", "voice")

	contentType := writer.FormDataContentType()

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8090/material", buf)
	if err != nil {
		log.Fatalf("New request failed: %s\n", err)
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Put failed: %s\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read failed: %s\n", err)
	}
	c.Writer.Write(body)

}
