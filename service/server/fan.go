package server

import (
	"strconv"
	"time"
	"wegou/database"
	"wegou/model"

	"github.com/gin-gonic/gin"
)

const (
	completedStatus   = 1 //完善的
	unCompletedStatus = 2 //未完善的
)

//粉丝列表
func GetFan(c *gin.Context) []model.Fan {
	pageStr := c.Query("page")
	web := c.Param("web")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	fan := model.Fan{}
	fans := fan.GetFan(web, page)

	return fans
}

//粉丝数量
func GetFanCount(c *gin.Context) (int, int, int) {

	web := c.Param("web")

	fan := model.Fan{}
	count := fan.GetFanCount(web)
	pageSize := database.FanPageSize

	var pageNum int
	if count%pageSize == 0 {
		pageNum = count / pageSize
	} else {
		pageNum = count/pageSize + 1
	}

	return count, pageSize, pageNum
}

//添加粉丝
func AddFan(web string, accountId int, wx string) (bool, error) {

	createdAt := time.Now().Unix()
	fan := model.Fan{
		Wx:        wx,
		Nickname:  "",
		AccountId: accountId,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Status:    unCompletedStatus,
	}

	fan.AddFan(web)
	return true, nil
}
