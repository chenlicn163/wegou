package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//配置初始化
func init() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	viper.Set("app_path", appPath)
	viper.Set("log_path", filepath.Join(appPath, "logs"))
	viper.Set("config_path", filepath.Join(appPath, "config.conf"))

	viper.SetConfigType("toml")
	viper.SetConfigFile(viper.GetString("config_path"))
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}

	// 监控并动态加载配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: " + e.Name)
	})
}

//读取监听IP和监听端口
func GetWebConfig() Web {
	conf := Web{
		Host: viper.GetString("listen.host"),
		Port: viper.GetString("listen.port"),
	}
	return conf
}

//读取kafka配置
func GetKafkaConfig() Kafka {
	conf := Kafka{
		Blockers:       strings.Split(viper.GetString("kafka.broker"), ","),
		CustomerTopics: strings.Split(viper.GetString("kafka.customer_topic"), ","),
		MaterialTopics: strings.Split(viper.GetString("kafka.material_topic"), ","),
		CustomerGroup:  viper.GetString("kafka.customer_group"),
		MaterialGroup:  viper.GetString("kafka.material_group"),
	}
	return conf
}

//读取主数据库配置
func GetDbConfig(account string) Db {
	//log.Println(dbConfig)
	conf := Db{
		DbHost:     viper.GetString(account + ".host"),
		DbPort:     viper.GetString(account + ".port"),
		DbUser:     viper.GetString(account + ".user"),
		DbPassword: viper.GetString(account + ".password"),
		DbName:     viper.GetString(account + ".db_name"),
	}

	return conf
}

//读取redis配置
func GetRedisConfig() Redis {
	conf := Redis{
		Server: viper.GetString("redis.server"),
		Auth:   viper.GetString("redis.auth"),
		Db:     viper.GetInt("redis.db"),
	}

	return conf
}

//读取缓存、上传、数据库使用类型配置
func GetToolsConfig() Tools {
	conf := Tools{
		Cache:    viper.GetString("tools.cache"),
		Upload:   viper.GetString("tools.upload"),
		Database: viper.GetString("tools.database"),
	}

	return conf
}
