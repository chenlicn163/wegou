package server

import (
	"wegou/model"
)

const (
	db = "wegou"
)

func GetWechatByCode(web string) model.Wechat {

	wechat := model.Wechat{}
	wechat.Code = web
	wechat.GetWechatByCode(db)

	return wechat
}
