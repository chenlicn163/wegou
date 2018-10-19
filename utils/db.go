package utils

import (
	"fmt"
	"wegou/config"
	"wegou/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database interface {
	Open() *gorm.DB
}

type mysql struct {
	web string
}

func GetDb(web string) (dbMysql *mysql) {
	toolsConfig := config.GetToolsConfig()
	switch toolsConfig.Database {
	case "mysql":
		dbMysql = &mysql{web: web}
	}

	return dbMysql
}

func (dbMysql *mysql) Open() *gorm.DB {
	conf := types.Db{}
	switch dbMysql.web {
	case "wegou":
		conf = config.GetDbConfig(dbMysql.web)
	default:
		conf = GetWechatConfig(dbMysql.web)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	conn.DB().SetMaxIdleConns(10)
	conn.SingularTable(true)
	return conn
}
