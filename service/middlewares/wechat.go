package middlewares

import (
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

type WechatAuth struct {
}

func (wechatAuth *WechatAuth) Do() gin.HandlerFunc {
	return func(c *gin.Context) {
		web := c.Param("web")
		wechatCache := server.WechatCache{Web: web}
		account, _ := wechatCache.Get()
		//logrus.Info(account)
		if account.Id == 0 {
			wechatCache.Set()
		}

		c.Next()
	}
}
