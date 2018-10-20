package model

import (
	"wegou/types"
	"wegou/utils"
)

//公众号实体
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
	DbHost      string `json:"db_host"`
	DbName      string `json:"db_name"`
	DbPort      string `json:"db_port"`
	DbUser      string `json:"db_user"`
	DbPassword  string `json:"db_password"`
	AuthStatus  int    `json:"auth_status"`
	AccountId   int    `json:"account_id"`
}

//获取公众号
func (wechat *Wechat) GetWechat(web string, page int) []Wechat {
	pageSize := types.AccountPageSize
	offset := pageSize * (page - 1)

	conn := utils.GetDb(web).Open()
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

//根据公众号code获取公众号信息
func (wechat *Wechat) GetWechatByCode(web string) {
	conn := utils.GetDb(web).Open()
	defer conn.Close()
	if conn == nil {
		return
	}
	conn.Model(&Wechat{}).
		Where("code=?", wechat.Code).First(&wechat)

}

//添加公众号
func (wechat *Wechat) AddWechat(web string) bool {
	conn := utils.GetDb(web).Open()
	defer conn.Close()
	if conn == nil {
		return false
	}
	conn.Model(&Wechat{}).Create(wechat)
	return true
}
