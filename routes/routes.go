package routes

import (
	"wegou/controller"
	"wegou/service/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthWechat())
	r.Any("/wechat_callback/:web", func(c *gin.Context) {
		web := c.Param("web")

		srv := WechatServe(web)
		srv.ServeHTTP(c.Writer, c.Request, nil)
	})
	wegou := r.Group("/wegou")

	wegou.Use(middlewares.AuthWechat())
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
