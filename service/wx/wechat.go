package wx

import (
	"wegou/utils"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

//获取微信回复
func GetMessage(ctx *core.Context) *WgMessage {
	return &WgMessage{Ctx: ctx}
}

//获取微信粉丝
func GetCustomer(web string, openId string) *WgUser {
	wechatConfig := utils.GetWechatConfig(web)
	srv := core.NewDefaultAccessTokenServer(wechatConfig.Appid, wechatConfig.Appsecret, nil)
	clt := core.NewClient(srv, nil)

	return &WgUser{Clt: clt, OpenId: openId}
}
