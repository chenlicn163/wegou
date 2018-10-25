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
	deletedStatus   = 3 //删除中
)

type MaterialDto struct {
	Web     string
	Code    string
	Message string
	Data    interface{}
}

func (result *MaterialDto) Material(msg string) {
	topic := gjson.Get(msg, "kafka.topic").String()
	switch topic {
	case "material-add":
		result.AddMaterial(msg)
	}

}

func (dto *MaterialDto) AddMaterial(msg string) {

	web := gjson.Get(msg, "web").String()
	dto.Web = web
	materialId := gjson.Get(msg, "material_id").String()

	id, err := strconv.Atoi(materialId)
	if err != nil {

	}

	//读取material
	dto.getMaterialById(id)
	entity := dto.Data.(model.Material)

	switch entity.SourceType {
	case material.MaterialTypeImage:
		//上传到微信服务器
		dto.uploadImageToWx()
	case material.MaterialTypeVoice:
	case material.MaterialTypeThumb:
	case material.MaterialTypeNews:

	}

	//更新数据库
	dto.upDataImageDb()
}

func (dto *MaterialDto) getMaterialById(id int) {
	entity := model.Material{Id: id}
	entity.GetMaterialById(dto.Web)
	dto.Data = entity
}

func (dto *MaterialDto) uploadImageToWx() {

	entity := dto.Data.(model.Material)

	wechatCache := WechatCache{Web: dto.Web}
	wechat, _ := wechatCache.Get()
	srv := core.NewDefaultAccessTokenServer(wechat.Appid, wechat.Appsecret, nil)
	clt := core.NewClient(srv, nil)

	filename := dto.getPath(entity.Pic)
	mediaId, url, err := material.UploadImage(clt, filename)
	if err != nil {
		logrus.Error(err.Error())
	}

	entity.MediaId = mediaId
	entity.Url = url

	dto.Data = entity
}

func (dto *MaterialDto) upDataImageDb() {
	entity := dto.Data.(model.Material)
	unix := time.Now().Unix()
	entity.UpdatedAt = unix
	entity.WxCreatedAt = unix
	entity.WxUpdatedAt = unix
	entity.UpdateMaterial(dto.Web)
	logrus.Info(entity)
	//logStr := "web:" + web + ",upload  " + material.MaterialTypeImage + " success!" + "mediaId:" + mediaId + ",url:" + url
	//logrus.Info(logStr)
}

func (dto *MaterialDto) getPath(filename string) (fileFullPath string) {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Error(err.Error())
	}

	return appPath + "/" + filename

}
