package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	OriId     string
	AppId     string
	Token     string
	AppSecret string
	AesKey    string
}

func init() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	viper.Set("app_path", appPath)
	//viper.Set("log_path", filepath.Join(appPath, "logs"))
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

func GetDbConfig() Config {
	conf := Config{
		OriId:     viper.GetString("wechat.oriid"),
		AppId:     viper.GetString("wechat.appId"),
		Token:     viper.GetString("wechat.token"),
		AppSecret: viper.GetString("wechat.appcecret"),
		AesKey:    viper.GetString("wechat.aeskey"),
	}

	return conf
}
