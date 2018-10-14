package server

import (
	"encoding/json"
	"errors"
	"wegou/model"
	"wegou/utils"

	"github.com/sirupsen/logrus"
)

//设置公众号缓存
func SetAccountCache(web string) {
	account := GetAccountByName(web)
	jsonAccount, err := json.Marshal(account)

	if err != nil {
		logrus.Error("json account error:" + err.Error())
	} else {
		utils.Redis(web).Set("account", string(jsonAccount))
	}
}

//获取公众号缓存
func GetAccountCache(web string) (account model.Account, err error) {

	jsonAccount, err := utils.Redis(web).Get("account")

	if err != nil {
		return account, errors.New("json account error:" + err.Error())
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &account)
	}

	return account, nil

}
