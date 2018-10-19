package server

import (
	"encoding/json"
	"errors"
	"wegou/model"
	"wegou/utils"

	"github.com/sirupsen/logrus"
)

const (
	db = "wegou"
)

type WechatCache struct {
	Web string
}

func (wechatCache *WechatCache) GetWechatByCode() model.Wechat {

	wechat := model.Wechat{}
	wechat.Code = wechatCache.Web
	wechat.GetWechatByCode(db)

	return wechat
}

//设置公众号缓存
func (wechatCache *WechatCache) Set() {
	wechat := wechatCache.GetWechatByCode()
	jsonAccount, err := json.Marshal(wechat)

	if err != nil {
		logrus.Error("json wechat error:" + err.Error())
	} else {
		utils.GetCache(wechatCache.Web).Set("wechat", string(jsonAccount))
	}
}

//获取公众号缓存
func (wechatCache *WechatCache) Get() (wechat model.Wechat, err error) {

	jsonAccount, err := utils.GetCache(wechatCache.Web).Get("wechat")

	if err != nil {
		return wechat, errors.New("json account error:" + err.Error())
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &wechat)
	}

	return wechat, nil

}
