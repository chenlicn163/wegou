package server

import (
	"wegou/model"

	"github.com/gin-gonic/gin"
)

func GetFans(c *gin.Context) []model.Fans {

	return []model.Fans{}
}

func GetFansCount(c *gin.Context) (int, int, int) {
	return 0, 0, 0
}
