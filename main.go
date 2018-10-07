package main

import (
	_ "wegou/config"
	"wegou/engine"
)

func main() {

	//启动服务
	engine.Run()

	//engine.Test()
}
