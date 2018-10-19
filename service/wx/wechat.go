package wx

import (
	"wegou/utils"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

func GetMessage(ctx *core.Context) *Message {
	return &Message{Ctx: ctx}
}

func GetCustomer(web string, openId string) *Customer {
	wechatConfig := utils.GetWechatConfig(web)
	srv := core.NewDefaultAccessTokenServer(wechatConfig.Appid, wechatConfig.Appsecret, nil)
	clt := core.NewClient(srv, nil)

	return &Customer{Clt: clt, OpenId: openId}
}
