package server

import (
	"azoya/nova/json"
	"strconv"
	"time"
	"wegou/model"
	"wegou/service/consumer"
	"wegou/types"

	"github.com/gin-gonic/gin"
)

const (
	completedStatus   = 1 //完善的
	unCompletedStatus = 2 //未完善的
)

//粉丝Dto
type FanDto struct {
	Code    string
	Message string
	Data    interface{}
}

//粉丝列表
func (result *FanDto) GetFan(c *gin.Context) {

	web := c.Param("web")
	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Message = types.AccountParamErrorMsg
		return
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
	return
}

//添加粉丝
func (result *FanDto) AddFan(web string, wx string) {

	if web == "" {
		result.Code = types.AccountParamErrorCode
		result.Code = types.AccountParamErrorMsg
		return
	}

	wechat, err := (&WechatCache{Web: web}).Get()
	if err != nil {
		result.Code = types.AccountNotExistCode
		result.Code = types.AccountNotExistMsg
		return
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
		HeadImageURL:   "",
		SubscribeTime:  0,
		UnionId:        "",
		GroupId:        0,
		TagidList:      []int{},
		SubscribeScene: "",
	}
	fan.AddFan(web)

	kafka := map[string]interface{}{
		"kafka":      map[string]string{"topic": "customer-add"},
		"open_id":    wx,
		"account_id": wechat.Id,
	}
	kafkaBytes, err := json.Marshal(kafka)
	if err != nil {
		result.Code = types.FanAddKafkaFailedCode
		result.Message = types.FanAddKafkaFailedMsg

		return
	} else {
		(&consumer.Task{Topics: "customer-add"}).AsyncProducer(string(kafkaBytes))
	}

	result.Code = types.WechatSuccessCode
	result.Message = types.WechatSuccessMsg
	return
}
