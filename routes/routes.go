package routes

import (
	"wegou/controller"
	"wegou/service/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	wechatAuth := middlewares.WechatAuth{}
	r.Use(wechatAuth.Do())
	r.Any("/wechat_callback/:web", func(c *gin.Context) {
		web := c.Param("web")

		srv := WechatServe(web)
		srv.ServeHTTP(c.Writer, c.Request, nil)
	})
	wegou := r.Group("/wegou")

	wegou.Use(wechatAuth.Do())
	//素材管理
	materialController := controller.MaterialController{}
	wegou.GET("/material/:web", materialController.ListMaterial)
	wegou.DELETE("/material/:web/:id", materialController.DeleteMaterial)
	wegou.PUT("/material/:web", materialController.AddMaterial)
	//粉丝管理
	fanController := controller.FanController{}
	wegou.GET("/fan/:web", fanController.ListFans)

	//前台
	r.POST("/test", materialController.AddFile)

	return r
}
