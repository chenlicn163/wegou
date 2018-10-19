package utils

import (
	"encoding/json"
	"errors"
	"reflect"
	"wegou/config"

	"github.com/sirupsen/logrus"
)

func InArray(need interface{}, needArr interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(needArr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(needArr)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(need, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

//获取公众号缓存
func GetWechatConfig(web string) (wechat config.Db) {
	jsonAccount, err := GetCache(web).Get("wechat")
	if err != nil {
		logrus.Error(errors.New("json account error:" + err.Error()))
		return wechat
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &wechat)
	}
	return wechat
}
