package database

import (
	"log"

	"github.com/spf13/viper"
)

const (
	MaterialPageSize = 20
)

//数据库配置
type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func GetDbConfig(account string) Db {
	dbConfig := viper.GetString("account." + account)
	log.Println(dbConfig)
	conf := Db{
		Host:     viper.GetString(dbConfig + ".host"),
		Port:     viper.GetString(dbConfig + ".port"),
		User:     viper.GetString(dbConfig + ".user"),
		Password: viper.GetString(dbConfig + ".password"),
		DbName:   viper.GetString(dbConfig + ".db_name"),
	}

	return conf
}
