package controller

import (
	"net/http"
	"wegou/service/server"
	"wegou/types"

	"github.com/gin-gonic/gin"
)

type FanController struct{}

//查询粉丝
func (fanController FanController) ListFans(c *gin.Context) {
	result := server.FanDto{}
	result.GetFan(c)
	var data map[string]interface{}
	if result.Code == types.WechatSuccessCode {
		rslt := result.Data.(map[string]interface{})
		page := rslt["page"].(map[string]int)

		data = map[string]interface{}{
			"fans": rslt["fans"],
			"page": map[string]int{
				"page_count": page["page_count"],
				"page_size":  page["page_size"],
				"page_num":   page["page_num"],
			},
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    result.Code,
			"message": result.Message,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    result.Code,
			"message": result.Message,
		})
	}

}

//删除粉丝
