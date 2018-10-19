package utils

import (
	"wegou/config"
	"wegou/types"
	"wegou/utils/database"

	"github.com/jinzhu/gorm"
)

type Database interface {
	Open() *gorm.DB
}

func GetDb(web string) (db *database.Mysql) {
	toolsConfig := config.GetToolsConfig()

	conf := types.Db{}
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
