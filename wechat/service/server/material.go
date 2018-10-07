package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"wegou/wechat/model"
	"wegou/wechat/utils"

	"github.com/gorilla/mux"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	availableStatus = 1
	addedStatus     = 2
	deletedStatus   = 3

	materialTemporary = 1
	materialForever   = 2

	materialTypeImage = "image"
	materialTypice    = "voice"
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
	size := 20
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
func GetMaterial(r *http.Request) []model.Material {

	query := r.URL.Query()
	pageStr := query.Get("page")
	MaterialType := query.Get("material")
	sourceType := query.Get("source")
	statusStr := query.Get("status")

	page, _ := strconv.Atoi(pageStr)
	status, _ := strconv.Atoi(statusStr)
	mat := model.Material{}
	materials := mat.GetMaterial(page, MaterialType, sourceType, status)

	return materials
}

//添加素材
func AddMaterial(r *http.Request) bool {

	fileName, err := utils.Upload(r, "upload")
	if err != nil {
		log.Println(err)
	}

	mat := model.Material{
		Title:        "test-title",
		Pic:          fileName,
		Author:       "test-author",
		Digest:       "test-digest",
		Content:      "test-content",
		ShowCoverPic: 0,
		MaterialType: materialForever,
		AccountId:    1,
		Status:       addedStatus,
		SourceType:   "image",
	}
	mat.AddMaterial()
	return true
}

//删除素材
func DelMaterial(r *http.Request) bool {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	mat := model.Material{}
	mat.Id = id
	//标记为删除状态
	mat.Status = deletedStatus
	mat.UpdateMaterial()

	//触发任务，删除微信服务器

	return true
}
