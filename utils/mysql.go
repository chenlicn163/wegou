package utils

import (
	"fmt"
	"wegou/config"

	"wegou/service/server"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var conn *gorm.DB

func Open(account string) (db *gorm.DB) {
	conf := config.GetDbConfig(account)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)

	return db
}

func OpenWechat(web string) (db *gorm.DB) {
	wechat, _ := server.GetWechatCache(web)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		wechat.DbUser, wechat.DbPassword, wechat.DbHost, wechat.DbPort, wechat.DbName)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)

	return db
}

/*func init() {
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
*/
