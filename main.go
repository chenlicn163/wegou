package main

import (
	"wegou/config"
	"wegou/engine/routes"
	"wegou/engine/task"
)

func main() {

	//启动服务
	kafkaConfig := config.GetKafkaConfig()
	go task.CustomerConsumer(kafkaConfig)
	go task.MaterialConsumer(kafkaConfig)

	wechatConfig := config.GetWechatConfig()
	r := routes.Routes(wechatConfig)
	webConfig := config.GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	r.Run(addr)

}
