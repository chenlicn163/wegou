package database

import (
	"fmt"
	"wegou/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
	config.Db
}

func (dbMysql *Mysql) Open() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbMysql.DbUser, dbMysql.DbPassword, dbMysql.DbHost, dbMysql.DbPort, dbMysql.DbName)
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	conn.DB().SetMaxIdleConns(10)
	conn.SingularTable(true)
	return conn
}
