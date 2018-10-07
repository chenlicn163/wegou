package database

import "github.com/spf13/viper"

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

func GetDbConfig() Db {
	conf := Db{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.db_name"),
	}

	return conf
}
