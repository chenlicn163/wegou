package server

import (
	"encoding/json"
	"errors"
	"wegou/model"
	"wegou/utils"

	"github.com/sirupsen/logrus"
)

//设置公众号缓存
func SetWechatCache(web string) {
	wechat := GetWechatByCode(web)
	jsonAccount, err := json.Marshal(wechat)

	if err != nil {
		logrus.Error("json wechat error:" + err.Error())
	} else {
		utils.Redis(web).Set("wechat", string(jsonAccount))
	}
}

//获取公众号缓存
func GetWechatCache(web string) (account model.Account, err error) {

	jsonAccount, err := utils.Redis(web).Get("wechat")

	if err != nil {
		return account, errors.New("json account error:" + err.Error())
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &account)
	}

	return account, nil

}
