package model

import (
	"wegou/types"
	"wegou/utils"
)

type Wechat struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Oriid       string `json:"oriid"`
	Appid       string `json:"appid"`
	Appsecret   string `json:"appsecret"`
	Token       string `json:"token"`
	Aeskey      string `json:"aeskey"`
	AccountType int    `json:"account_type"`
	ServiceType int    `json:"service_type"`
	Status      int    `json:"status"`
	AccountId   int    `json:"account_id"`
}

//获取公众号
func (wechat *Wechat) GetWechat(web string, page int) []Wechat {
	pageSize := types.AccountPageSize
	offset := pageSize * (page - 1)

	conn := utils.Open(web)
	defer conn.Close()
	if conn == nil {
		return nil
	}

	var wechats []Wechat
	conn.Model(&Wechat{}).
		Offset(offset).Limit(pageSize).
		Find(&wechats)

	return wechats
}

func (wechat *Wechat) GetWechatByCode(web string) {
	conn := utils.Open(web)
	defer conn.Close()
	conn.Model(&Wechat{}).
		Where("code=?", wechat.Code).First(&wechat)

}

//添加公众号
func (wechat *Wechat) AddWechat(web string) bool {
	conn := utils.Open(web)
	defer conn.Close()
	conn.Model(&Wechat{}).Create(wechat)
	return true
}
