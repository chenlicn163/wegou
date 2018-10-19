package wx

import (
	"wegou/model"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/user"
)

type WgUser struct {
	Clt    *core.Client
	OpenId string
}

func (wgUser *WgUser) Get() (fan model.Fan) {
	info, err := user.Get(wgUser.Clt, wgUser.OpenId, user.LanguageZhCN)
	if err != nil {
		return fan
	}
	fan.Nickname = info.Nickname
	fan.Sex = info.Sex
	fan.Language = info.Language
	fan.City = info.City
	fan.Province = info.Province
	fan.HeadImageURL = info.HeadImageURL
	fan.SubscribeTime = info.SubscribeTime
	fan.UnionId = info.UnionId
	fan.GroupId = info.GroupId
	fan.TagidList = info.TagIdList
	fan.SubscribeScene = info.SubscribeScene

	return fan
}
