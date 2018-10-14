package routes

import (
	"wegou/controller"
	"wegou/service/middlewares"
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthAccount())
	r.Any("/wechat_callback/:web", func(c *gin.Context) {
		web := c.Param("web")
		query := c.Request.URL.Query()
		query.Add("web", web)

		server.GetAccountCache(web)
		account, _ := server.GetAccountCache(web)

		srv := WechatServe(account)
		srv.ServeHTTP(c.Writer, c.Request, query)
	})
	wegou := r.Group("/wegou")

	wegou.Use(middlewares.AuthAccount())
	//素材管理
	wegou.GET("/material/:web", controller.ListMaterialServe)
	wegou.DELETE("/material/:web/:id", controller.DeleteMaterialServe)
	wegou.PUT("/material/:web", controller.AddMaterialServe)
	//粉丝管理
	wegou.GET("/fan/:web", controller.ListFansServe)

	//前台

	r.POST("/test", controller.AddFileServe)

	return r
}
