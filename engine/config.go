package engine

import (
	"strings"
	"wegou/engine/task"

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

func GetKafkaConfig() task.Kafka {
	conf := task.Kafka{
		Blockers:       strings.Split(viper.GetString("kafka.broker"), ","),
		CustomerTopics: strings.Split(viper.GetString("kafak.customer_topic"), ","),
		MaterialTopics: strings.Split(viper.GetString("kafak.material_topic"), ","),
		CustomerGroup:  viper.GetString("kafka.customer_group"),
		MaterialGroup:  viper.GetString("kafka.material_group"),
	}
	return conf
}
