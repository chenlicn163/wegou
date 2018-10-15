package middlewares

import (
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

func AuthWechat() gin.HandlerFunc {
	return func(c *gin.Context) {
		web := c.Param("web")
		account, _ := server.GetWechatCache(web)
		if account.Id == 0 {
			server.SetWechatCache(web)
		}

		c.Next()
	}
}
