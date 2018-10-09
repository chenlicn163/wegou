package engine

import (
	"wegou/engine/admin"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	srv := WechatServe()
	r := gin.Default()
	r.Any("/wechat_callback/:web", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request, map[string][]string{"web": []string{c.Param("web")}})
	})
	management := r.Group("/admin")
	//素材管理
	management.GET("/material/:web", admin.ListMaterialServe)
	management.DELETE("/material/:web/:id", admin.DeleteMaterialServe)
	management.PUT("/material/:web", admin.AddMaterialServe)
	//粉丝管理
	management.GET("/fan/:web", admin.ListFansServe)

	//前台

	r.POST("/test", admin.AddFileServe)

	return r
}
