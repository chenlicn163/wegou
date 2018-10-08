package admin

import (
	"net/http"
	"wegou/service/server"

	"github.com/gin-gonic/gin"
)

//查询粉丝
func ListFansServe(c *gin.Context) {
	fans := server.GetFans(c)
	count, pageSize, pageNum := server.GetFansCount(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "success",
		"data": map[string]interface{}{
			"fans": fans,
			"page": map[string]int{
				"count":     count,
				"page_size": pageSize,
				"page_num":  pageNum,
			},
		},
	})
}

//删除粉丝
