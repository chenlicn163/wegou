package middlewares

import (
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

func AuthAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		web := c.Param("web")
		account, _ := server.GetAccountCache(web)
		if account.Id == 0 {
			server.SetAccountCache(web)
		}

		c.Next()
	}
}
