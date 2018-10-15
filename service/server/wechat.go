package server

import (
	"wegou/model"
)

const (
	db = "wegou"
)

func GetWechatByCode(web string) model.Wechat {

	wechat := model.Wechat{}
	wechat.Name = web
	wechat.GetWechatByCode(db)

	return wechat
}
