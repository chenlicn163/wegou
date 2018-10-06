package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

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

func GetWebConfig() Web {
	conf := Web{
		Host: viper.GetString("listen.host"),
		Port: viper.GetString("listen.port"),
	}
	return conf
}
