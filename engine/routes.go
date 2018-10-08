package engine

import (
	"wegou/engine/admin"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	srv := WechatServe()
	r := gin.Default()
	r.GET("/wechat_callback", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request, nil)
	})
	management := r.Group("/admin")
	management.GET("/material/:web", admin.ListMaterialServe)
	management.DELETE("/material/:web/:id", admin.DeleteMaterialServe)
	management.PUT("/material/:web", admin.AddMaterialServe)

	r.POST("/test", admin.AddFileServe)

	return r
}
