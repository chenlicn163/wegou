package engine

import (
	"github.com/spf13/viper"
)

//站点配置
type Web struct {
	Port string
	Host string
}

//微信配置
type Wechat struct {
	OriId     string
	AppId     string
	Token     string
	AppSecret string
	AesKey    string
}

func GetWebConfig() Web {
	conf := Web{
		Host: viper.GetString("listen.host"),
		Port: viper.GetString("listen.port"),
	}
	return conf
}

func GetWechatConfig() Wechat {
	conf := Wechat{
		OriId:     viper.GetString("wechat.oriid"),
		AppId:     viper.GetString("wechat.appId"),
		Token:     viper.GetString("wechat.token"),
		AppSecret: viper.GetString("wechat.appsecret"),
		AesKey:    viper.GetString("wechat.aeskey"),
	}
	return conf
}
