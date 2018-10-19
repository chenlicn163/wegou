package server

import (
	"strconv"
	"time"
	"wegou/model"
	"wegou/utils"

	"github.com/gin-gonic/gin"

	"wegou/types"
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

type MaterialDto struct {
	Code    string
	Message string
	Data    interface{}
}

//获取素材
func (result *MaterialDto) GetMaterial(c *gin.Context) {

	web := c.Param("web")
	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Message = types.AccountParamErrorMsg
		return
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
		return
	}

	statusStr := c.Query("status")
	statusValues := []string{strconv.Itoa(availableStatus), strconv.Itoa(addedStatus), strconv.Itoa(deletedStatus)}
	if exists, _ := utils.InArray(statusStr, statusValues); statusStr != "" && !exists {
		result.Code = types.SourceStatusErrorCode
		result.Message = types.SourceStatusErrorMsg
		return
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
		return
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
	return
}

//添加素材
func (result *MaterialDto) AddMaterial(c *gin.Context) {

	web := c.Param("web")
	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Message = types.AccountParamErrorMsg
		return
	}

	wechat, err := (&WechatCache{Web: web}).Get()
	if err != nil {
		result.Code = types.AccountNotExistCode
		result.Code = types.AccountNotExistMsg
		return
	}

	title := c.PostForm("title")
	if title == "" {
		result.Code = types.MaterialTitleAddFailedCode
		result.Message = types.MaterialTitleAddFailedMsg
		return
	}

	uploadPath := "upload/" + web + "/" +
		strconv.Itoa(time.Now().Year()) + "/" +
		strconv.Itoa(int(time.Now().Month())) + "/" +
		strconv.Itoa(time.Now().Day())
	fileName, err := utils.GetUpload(uploadPath).UploadFile(c.Request)
	if err != nil {
		result.Code = types.MaterialFileAddFailedCode
		result.Message = types.MaterialFileAddFailedMsg
		return
	}

	showCoverPic, err := strconv.Atoi(c.PostForm("show_cover_pic"))
	if err != nil {
		showCoverPic = hideCoverPic
	}

	materialType, err := strconv.Atoi(c.PostForm("material_type"))
	if err != nil {
		materialType = materialTemporary
	}

	sourceType := c.PostForm("source_type")
	sourceTypeValues := []string{materialTypeImage, materialTypeVoice, materialTypeVideo, materialTypeThumb,
		materialTypeNews}
	if exists, _ := utils.InArray(sourceType, sourceTypeValues); sourceType != "" && !exists {
		result.Code = types.MaterialSourceTypeAddFailedCode
		result.Message = types.MaterialSourceTypeAddFailedMsg
		return
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
		AccountId:    wechat.Id,
		Status:       addedStatus,
		SourceType:   sourceType,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}

	mat.AddMaterial(web)
	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	return
}

//删除素材
func (result *MaterialDto) DelMaterial(c *gin.Context) {

	web := c.Param("web")
	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Message = types.AccountParamErrorMsg
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Code = types.MaterialIdDeleteFailedCode
		result.Message = types.MaterialIdDeleteFailedMsg
		return
	}
	mat := model.Material{}
	mat.Id = id
	//标记为删除状态
	mat.Status = deletedStatus
	mat.UpdateMaterial(web)

	//触发任务，删除微信服务器
	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	return
}
