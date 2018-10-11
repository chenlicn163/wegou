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

	result := server.GetMaterial(c)

	var data map[string]interface{}
	if result.Code == types.WechatSuccessCode {
		rslt := result.Data.(map[string]interface{})
		page := rslt["page"].(map[string]int)

		data = map[string]interface{}{
			"materials": rslt["materials"],
			"page": map[string]int{
				"page_count": page["page_count"],
				"page_size":  page["page_size"],
				"page_num":   page["page_num"],
			},
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    result.Code,
			"message": result.Message,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    result.Code,
			"message": result.Message,
		})
	}

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

	result := server.DelMaterial(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    result.Code,
		"message": result.Message,
	})
}

//添加永久素材
func AddMaterialServe(c *gin.Context) {

	result := server.AddMaterial(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    result.Code,
		"message": result.Message,
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
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8090/wegou/material/"+c.PostForm("web"), buf)
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
