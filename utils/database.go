package utils

import (
	"wegou/config"
	"wegou/types"
	"wegou/utils/database"
)

func GetDb(web string) (dbMysql *database.Mysql) {
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
		dbMysql = &database.Mysql{}
		dbMysql.DbUser = conf.DbUser
		dbMysql.DbPassword = conf.DbPassword
		dbMysql.DbHost = conf.DbHost
		dbMysql.DbPort = conf.DbPort
		dbMysql.DbName = conf.DbName
	}

	return dbMysql
}
