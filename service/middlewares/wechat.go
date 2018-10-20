package middlewares

import (
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

type WechatAuth struct {
}

//公众号中间件，获取公众号信息存入缓存
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
