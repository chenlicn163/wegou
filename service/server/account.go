package server

import (
	"wegou/model"
)

func GetAccountByName(web string) model.Account {

	account := model.Account{}
	account.Name = web
	account.GetAccountByName("wegou")

	return account
}
