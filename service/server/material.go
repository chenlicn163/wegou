package server

import (
	"fmt"
	"strconv"
	"time"
	"wegou/model"
	"wegou/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"

	"wegou/types"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	availableStatus = 1 //可用
	addedStatus     = 2 //添加中
	deletedStatus   = 3 //删除中

	materialTemporary = 1 //临时
	materialForever   = 2 //永久

	hideCoverPic = 0 //不显示图
	showCoverPic = 1 //显示图

	materialTypeImage = "image"
	materialTypeVoice = "voice"
	materialTypeVideo = "video"
	materialTypeThumb = "thumb"
	materialTypeNews  = "news"
)

//获取永久素材数量
func FetchMaterialCount(clt *core.Client) *material.MaterialCountInfo {

	info, err := material.GetMaterialCount(clt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	return info
}

//批量获取永久素材
func BatchFetchMaterial(clt *core.Client, materialType string, pageStr string) *material.BatchGetResult {
	size := types.MaterialPageSize
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	offset := (page - 1) * size
	rslt, err := material.BatchGet(clt, materialType, offset, size)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rslt)

	return rslt
}

//永久素材，同步微信服务器数据
func SyncMaterial() {
	//1 获取总数
	//2 分页获取,入库
}

//永久素材，上传到微信服务器
func SyncAddMaterial() {

}

//永久素材，从微信服务器上删除
func SyncDelMaterial() {

}

//---------------------------------------------------------------------------------------------

//获取素材
func GetMaterial(c *gin.Context) types.Dto {

	result := types.Dto{}

	web := c.Param("web")
	if web == "" {
		result.Code = types.WebFiledCode
		result.Code = types.WebFiledMsg
		return result
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	MaterialType := c.Query("material")
	MaterialTypeValues := []string{strconv.Itoa(materialTemporary), strconv.Itoa(materialForever)}
	if exists, _ := utils.InArray(MaterialType, MaterialTypeValues); MaterialType != "" && !exists {
		result.Code = types.MaterialTypeErrorCode
		result.Message = types.MaterialTypeErrorMsg
		return result
	}

	statusStr := c.Query("status")
	statusValues := []string{strconv.Itoa(availableStatus), strconv.Itoa(addedStatus), strconv.Itoa(deletedStatus)}
	if exists, _ := utils.InArray(statusStr, statusValues); statusStr != "" && !exists {
		result.Code = types.SourceStatusErrorCode
		result.Message = types.SourceStatusErrorMsg
		return result
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		status = 0
	}

	sourceType := c.Query("source")
	sourceTypeValues := []string{materialTypeImage, materialTypeVoice, materialTypeVideo, materialTypeThumb,
		materialTypeNews}
	if exists, _ := utils.InArray(sourceType, sourceTypeValues); sourceType != "" && !exists {
		result.Code = types.SourceTypeErrorCode
		result.Message = types.SourceTypeErrorMsg
		return result
	}

	mat := model.Material{}
	pageCount := mat.GetMaterialCount(web, MaterialType, sourceType, status)
	pageSize := types.MaterialPageSize
	var pageNum int
	if pageCount%pageSize == 0 {
		pageNum = pageCount / pageSize
	} else {
		pageNum = pageCount/pageSize + 1
	}
	materials := mat.GetMaterial(web, page, MaterialType, sourceType, status)

	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	result.Data = map[string]interface{}{
		"materials": materials,
		"page": map[string]int{
			"page_count": pageCount,
			"page_size":  pageSize,
			"page_num":   pageNum,
		},
	}
	return result
}

//添加素材
func AddMaterial(c *gin.Context) types.Dto {

	result := types.Dto{}

	web := c.Param("web")
	if web == "" {
		result.Code = types.WebFiledCode
		result.Code = types.WebFiledMsg
		return result
	}

	title := c.PostForm("title")
	if title == "" {
		result.Code = types.MaterialTitleAddFailedCode
		result.Message = types.MaterialTitleAddFailedMsg
		return result
	}

	uploadPath := "upload/" + web + "/" +
		strconv.Itoa(time.Now().Year()) + "/" +
		strconv.Itoa(int(time.Now().Month())) + "/" +
		strconv.Itoa(time.Now().Day())
	fileName, err := utils.Upload(c.Request, uploadPath)
	if err != nil {
		result.Code = types.MaterialFileAddFailedCode
		result.Message = types.MaterialFileAddFailedMsg
		return result
	}

	showCoverPic, err := strconv.Atoi(c.PostForm("show_cover_pic"))
	if err != nil {
		showCoverPic = hideCoverPic
	}

	materialType, err := strconv.Atoi(c.PostForm("material_type"))
	if err != nil {
		materialType = materialTemporary
	}

	accountId, err := strconv.Atoi(c.PostForm("account_id"))
	if err != nil {
		result.Code = types.MaterialFileAddFailedCode
		result.Message = types.MaterialFileAddFailedMsg
		return result
	}

	sourceType := c.PostForm("source_type")
	sourceTypeValues := []string{materialTypeImage, materialTypeVoice, materialTypeVideo, materialTypeThumb,
		materialTypeNews}
	if exists, _ := utils.InArray(sourceType, sourceTypeValues); sourceType != "" && !exists {
		result.Code = types.MaterialSourceTypeAddFailedCode
		result.Message = types.MaterialSourceTypeAddFailedMsg
		return result
	}

	author := c.PostForm("author")
	digest := c.PostForm("digest")
	content := c.PostForm("content")

	createdAt := time.Now().Unix()
	mat := model.Material{
		Title:        title,
		Pic:          fileName,
		Author:       author,
		Digest:       digest,
		Content:      content,
		ShowCoverPic: showCoverPic,
		MaterialType: materialType,
		AccountId:    accountId,
		Status:       addedStatus,
		SourceType:   sourceType,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}

	mat.AddMaterial(web)
	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	return result
}

//删除素材
func DelMaterial(c *gin.Context) types.Dto {
	result := types.Dto{}

	web := c.Param("web")
	if web == "" {
		result.Code = types.WebFiledCode
		result.Code = types.WebFiledMsg
		return result
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Code = types.MaterialIdDeleteFailedCode
		result.Message = types.MaterialIdDeleteFailedMsg
		return result
	}
	mat := model.Material{}
	mat.Id = id
	//标记为删除状态
	mat.Status = deletedStatus
	mat.UpdateMaterial(web)

	//触发任务，删除微信服务器
	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	return result
}
