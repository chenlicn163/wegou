package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"wegou/service/server"
	"wegou/types"

	"github.com/gin-gonic/gin"
)

//查询永久素材
func ListMaterialServe(c *gin.Context) {

	materials := server.GetMaterial(c)
	count, pageSize, pageNum := server.GetMaterialCount(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "success",
		"data": map[string]interface{}{
			"materials": materials,
			"page": map[string]int{
				"count":     count,
				"page_size": pageSize,
				"page_num":  pageNum,
			},
		},
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

	StatusJson := types.StatusJson{}
	if !flag {
		StatusJson.Code = types.MaterialDelFailedCode
		StatusJson.Message = types.MaterialDelFailedMsg
	} else {
		StatusJson.Code = types.WechatSuccessCode
		StatusJson.Message = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusJson.Code,
		"message": StatusJson.Message,
	})
}

//添加永久素材
func AddMaterialServe(c *gin.Context) {

	_, err := server.AddMaterial(c)

	StatusJson := types.StatusJson{}
	if err != nil {
		StatusJson.Code = types.MaterialAddFailedCode
		StatusJson.Message = types.MaterialAddFailedMsg + ",error:" + err.Error()
	} else {
		StatusJson.Code = types.WechatSuccessCode
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
	writer.WriteField("title", c.PostForm("title"))
	writer.WriteField("author", c.PostForm("author"))
	writer.WriteField("digest", c.PostForm("digest"))
	writer.WriteField("content", c.PostForm("content"))
	writer.WriteField("show_cover_pic", c.PostForm("show_cover_pic"))
	writer.WriteField("material_type", c.PostForm("material_type"))
	writer.WriteField("account_id", c.PostForm("account_id"))
	writer.WriteField("source_type", c.PostForm("source_type"))

	contentType := writer.FormDataContentType()

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8090/controller/material/test1", buf)
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
