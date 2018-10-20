package utils

import (
	"wegou/config"
	"wegou/utils/database"

	"github.com/jinzhu/gorm"
)

type Database interface {
	Open() *gorm.DB
}

//获取数据库
func GetDb(web string) (db *database.Mysql) {
	toolsConfig := config.GetToolsConfig()

	conf := config.Db{}
	switch web {
	case "wegou":
		conf = config.GetDbConfig(web)
	default:
		conf = GetWechatConfig(web)
	}
	switch toolsConfig.Database {
	case "mysql":
		db = &database.Mysql{}
		db.DbUser = conf.DbUser
		db.DbPassword = conf.DbPassword
		db.DbHost = conf.DbHost
		db.DbPort = conf.DbPort
		db.DbName = conf.DbName
	}

	return db
}
