package wx

import (
	"wegou/utils"

	"github.com/sirupsen/logrus"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

func GetMessage(ctx *core.Context) *Message {
	return &Message{Ctx: ctx}
}

func GetCustomer(web string, openId string) *Customer {
	wechatConfig, err := utils.GetWechatConfig(web)
	if err != nil {
		logrus.Error(err.Error())
	}
	srv := core.NewDefaultAccessTokenServer(wechatConfig.Appid, wechatConfig.Appsecret, nil)
	clt := core.NewClient(srv, nil)

	return &Customer{Clt: clt, OpenId: openId}
}
