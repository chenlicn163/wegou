package main

import (
	_ "wegou/config"
	"wegou/engine"
)

func main() {

	//启动服务

	r := engine.Routes()
	webConfig := engine.GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	r.Run(addr)

}
