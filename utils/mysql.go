package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"wegou/config"
	"wegou/types"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var conn *gorm.DB

func Open(account string) (db *gorm.DB) {
	conf := config.GetDbConfig(account)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}

	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)

	return db
}

func OpenWechat(web string) (db *gorm.DB) {
	conf, _ := GetWechatConfig(web)
	logrus.Info(conf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect to db failed,err:%+v", dsn)
	}
	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)
	return db
}

//获取公众号缓存
func GetWechatConfig(web string) (wechat types.Wechat, err error) {
	jsonAccount, err := Redis(web).Get("wechat")
	if err != nil {
		return wechat, errors.New("json account error:" + err.Error())
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &wechat)
	}
	return wechat, nil

}
