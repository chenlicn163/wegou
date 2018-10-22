package task

import (
	"strconv"
	"wegou/model"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type MaterialDto struct {
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

func (result *MaterialDto) AddMaterial(msg string) {
	web := gjson.Get(msg, "account").String()
	materialId := gjson.Get(msg, "material_id").String()
	result.syncMaterial(web, materialId)
}

func (result *MaterialDto) syncMaterial(web string, materialId string) {
	wechatCache := WechatCache{Web: web}
	wechat, _ := wechatCache.Get()

	logrus.Info(wechat)
	id, err := strconv.Atoi(materialId)
	if err != nil {

	}
	material := model.Material{Id: id}
	material.GetMaterialById(web)
	logrus.Info(material)

	/*srv := core.NewDefaultAccessTokenServer(wechat.Appid, wechat.Appsecret, nil)
	clt := core.NewClient(srv, nil)*/

}
