package server

import (
	"strconv"
	"time"
	"wegou/model"
	"wegou/types"

	"github.com/gin-gonic/gin"
)

const (
	completedStatus   = 1 //完善的
	unCompletedStatus = 2 //未完善的
)

//粉丝列表
func GetFan(c *gin.Context) types.Dto {
	result := types.Dto{}
	web := c.Param("web")
	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Message = types.AccountParamErrorMsg
		return result
	}

	pageStr := c.Query("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	fan := model.Fan{}
	fans := fan.GetFan(web, page)

	pageCount := fan.GetFanCount(web)
	pageSize := types.FanPageSize

	var pageNum int
	if pageCount%pageSize == 0 {
		pageNum = pageCount / pageSize
	} else {
		pageNum = pageCount/pageSize + 1
	}

	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	result.Data = map[string]interface{}{
		"materials": fans,
		"page": map[string]int{
			"page_count": pageCount,
			"page_size":  pageSize,
			"page_num":   pageNum,
		},
	}
	return result
}

//添加粉丝
func AddFan(web string, wx string) types.Dto {
	result := types.Dto{}

	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Code = types.AccountParamErrorMsg
		return result
	}

	wechat, err := GetWechatCache(web)
	if err != nil {
		result.Code = types.AccountNotExistCode
		result.Code = types.AccountNotExistMsg
		return result
	}

	createdAt := time.Now().Unix()
	fan := model.Fan{
		Wx:             wx,
		Nickname:       "",
		AccountId:      wechat.Id,
		CreatedAt:      createdAt,
		UpdatedAt:      createdAt,
		Status:         unCompletedStatus,
		Remark:         "",
		Sex:            0,
		Language:       "",
		City:           "",
		Province:       "",
		Country:        "",
		Headimgurl:     "",
		SubscribeTime:  0,
		Unionid:        "",
		Groupid:        0,
		TagidList:      "",
		SubscribeScene: "",
	}

	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	result.Data = fan
	return result
}
