package task

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
	"wegou/model"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

const (
	availableStatus = 1 //可用
	addedStatus     = 2 //添加中
	updatedStatus   = 4 //更新中
	deletedStatus   = 5 //删除中

)

type MaterialDto struct {
	Code    string
	Message string
	Web     string
	Data    model.Material
	clt     *core.Client
}

func (dto *MaterialDto) Material(msg string) {
	topic := gjson.Get(msg, "kafka.topic").String()
	web := gjson.Get(msg, "web").String()
	if web == "" {
		return
	}
	dto.Web = web

	//数据验证
	materialId := gjson.Get(msg, "material_id").String()
	id, err := strconv.Atoi(materialId)
	if err != nil {
		return
	}
	//读取material
	entity := model.Material{Id: id}
	entity.GetMaterialById(dto.Web)
	if entity.Title == "" {
		return
	}
	dto.Data = entity

	//微信client
	wechatCache := WechatCache{Web: dto.Web}
	wechat, _ := wechatCache.Get()
	srv := core.NewDefaultAccessTokenServer(wechat.Appid, wechat.Appsecret, nil)
	dto.clt = core.NewClient(srv, nil)

	switch topic {
	//永久素材添加
	case "material-add":
		dto.addMaterial(msg)
	//永久素材修改
	case "material-edit":
		dto.editMaterial(msg)
	//永久素材删除
	case "material-delete":
		dto.deleteMaterial(msg)
	//临时素材添加
	case "media-add":
	default:
		return
	}
}

//更新素材
func (dto *MaterialDto) editMaterial(msg string) {
	entity := dto.Data
	if entity.Status != updatedStatus {
		return
	}
}

//删除素材
func (dto *MaterialDto) deleteMaterial(msg string) {
	entity := dto.Data
	if entity.Status != deletedStatus {
		return
	}

	switch entity.SourceType {
	case material.MaterialTypeImage:
		//上传到微信服务器
		dto.deleteToWx()
	case material.MaterialTypeVoice:
	case material.MaterialTypeThumb:
	case material.MaterialTypeNews:
	default:
		return

	}

}

//添加素材
func (dto *MaterialDto) addMaterial(msg string) {

	entity := dto.Data
	if entity.Status != addedStatus {
		return
	}

	switch entity.SourceType {
	case material.MaterialTypeImage:
		//上传到微信服务器
		dto.uploadImageToWx()
	case material.MaterialTypeVoice:
	case material.MaterialTypeThumb:
	case material.MaterialTypeNews:
	default:
		return
	}

	//更新数据库
	entity = dto.Data
	unix := time.Now().Unix()
	entity.UpdatedAt = unix
	entity.WxCreatedAt = unix
	entity.WxUpdatedAt = unix
	entity.UpdateMaterial(dto.Web)
}

//删除图片素材
func (dto *MaterialDto) deleteToWx() {
	entity := dto.Data

	if entity.ThumbMediaId != "" {
		material.Delete(dto.clt, entity.ThumbMediaId)
	}

	if entity.MediaId != "" {
		material.Delete(dto.clt, entity.MediaId)
	}

	entity.Status = availableStatus
}

//上传新闻素材
func (dto *MaterialDto) uploadNewsToWx() {

}

//上传图片素材
func (dto *MaterialDto) uploadImageToWx() {

	entity := dto.Data

	filename := dto.getPath(entity.Pic)
	mediaId, url, err := material.UploadImage(dto.clt, filename)
	if err != nil {
		logrus.Error(err.Error())
	}

	entity.MediaId = mediaId
	entity.Url = url
	entity.Status = availableStatus

	dto.Data = entity
}

//获取服务器绝对路径
func (dto *MaterialDto) getPath(filename string) (fileFullPath string) {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Error(err.Error())
	}

	return appPath + "/" + filename
}
