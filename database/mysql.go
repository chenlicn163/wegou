package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func Open() (db *gorm.DB) {
	return conn
}

func init() {
	conf := GetDbConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)
	conn = db
}
