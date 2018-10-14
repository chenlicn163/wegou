package model

import (
	"wegou/types"
	"wegou/utils"
)

type Account struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Oriid       string `json:"oriid"`
	Appid       string `json:"appid"`
	Appsecret   string `json:"appsecret"`
	Token       string `json:"token"`
	Aeskey      string `json:"aeskey"`
	CreatedTime int64  `json:"created_time"`
	UpdatedTime int64  `json:"updated_time"`
	AccountType int    `json:"account_type"`
	ServiceType int    `json:"service_type"`
	Status      int    `json:"status"`
}

//获取公众号
func (account *Account) GetAccount(web string, page int) []Account {
	pageSize := types.AccountPageSize
	offset := pageSize * (page - 1)

	conn := utils.Open(web)
	defer conn.Close()
	if conn == nil {
		return nil
	}

	var accounts []Account
	conn.Model(&Account{}).
		Offset(offset).Limit(pageSize).
		Find(&accounts)

	return accounts
}

func (account *Account) GetAccountByName(web string) {
	conn := utils.Open(web)
	defer conn.Close()
	conn.Model(&Account{}).
		Where("name=?", account.Name).First(&account)

}

//添加公众号
func (account *Account) AddAccount(web string) bool {
	conn := utils.Open(web)
	defer conn.Close()
	conn.Model(&Fan{}).Create(account)
	return true
}
