package main

import (
	"wegou/config"
	"wegou/routes"
	"wegou/task"
)

func init() {
	config.InitLog()
}

func main() {

	//启动服务
	kafkaConfig := config.GetKafkaConfig()

	go (&task.CustomerConCumer{}).Consumer(kafkaConfig)
	go (&task.MaterialConsumer{}).Consumer(kafkaConfig)

	r := routes.Routes()
	webConfig := config.GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	r.Run(addr)

}
