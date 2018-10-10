package server

import (
	"errors"
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
func GetMaterial(c *gin.Context) []model.Material {

	pageStr := c.Query("page")
	MaterialType := c.Query("material")
	sourceType := c.Query("source")
	statusStr := c.Query("status")
	web := c.Param("web")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		status = 0
	}

	mat := model.Material{}
	materials := mat.GetMaterial(web, page, MaterialType, sourceType, status)

	return materials
}

//获取素材数量
func GetMaterialCount(c *gin.Context) (int, int, int) {

	MaterialType := c.Query("material")
	sourceType := c.Query("source")
	statusStr := c.Query("status")
	web := c.Param("web")

	status, _ := strconv.Atoi(statusStr)

	mat := model.Material{}
	count := mat.GetMaterialCount(web, MaterialType, sourceType, status)
	pageSize := types.MaterialPageSize

	var pageNum int
	if count%pageSize == 0 {
		pageNum = count / pageSize
	} else {
		pageNum = count/pageSize + 1
	}

	return count, pageSize, pageNum

}

//添加素材
func AddMaterial(c *gin.Context) (bool, error) {

	fileName, err := utils.Upload(c.Request, "upload")
	if err != nil {
		return false, errors.New("field upload :" + err.Error())
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
		return false, errors.New("field account_id :" + err.Error())
	}

	sourceType := c.PostForm("source_type")
	if sourceType == "" {
		return false, err
	}
	createdAt := time.Now().Unix()
	mat := model.Material{
		Title:        c.PostForm("title"),
		Pic:          fileName,
		Author:       c.PostForm("author"),
		Digest:       c.PostForm("digest"),
		Content:      c.PostForm("content"),
		ShowCoverPic: showCoverPic,
		MaterialType: materialType,
		AccountId:    accountId,
		Status:       addedStatus,
		SourceType:   sourceType,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}

	web := c.Param("web")
	mat.AddMaterial(web)
	return true, nil
}

//删除素材
func DelMaterial(c *gin.Context) bool {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	mat := model.Material{}
	mat.Id = id
	//标记为删除状态
	mat.Status = deletedStatus
	web := c.Param("web")
	mat.UpdateMaterial(web)

	//触发任务，删除微信服务器

	return true
}
