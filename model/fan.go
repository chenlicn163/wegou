package model

import (
	"wegou/types"
	"wegou/utils"
)

//粉丝实体
type Fan struct {
	Id             int    `json:"id"`
	Wx             string `json:"wx"`
	Nickname       string `json:"nickname"`
	CreatedAt      int64  `json:"created_time"`
	UpdatedAt      int64  `json:"updated_time"`
	AccountId      int    `json:"account_id"`
	Status         int    `json:"status"`
	Remark         string `json:"remark"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImageURL   string `json:"headimageurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionId        string `json:"unionid"`
	GroupId        int64  `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
}

//获取粉丝
func (fan *Fan) GetFan(web string, page int) []Fan {
	pageSize := types.FanPageSize
	offset := pageSize * (page - 1)

	conn := utils.GetDb(web).Open()
	defer conn.Close()
	if conn == nil {
		return nil
	}

	var fans []Fan
	conn.Model(&Fan{}).
		Offset(offset).Limit(pageSize).
		Where("account_id=?", web).
		Find(&fans)

	return fans
}

//粉丝总数量
func (fan *Fan) GetFanCount(web string) int {

	conn := utils.GetDb(web).Open()
	defer conn.Close()
	if conn == nil {
		return 0
	}

	var count int
	conn.Model(&Fan{}).
		Where("account_id=?", web).
		Count(&count)

	return count
}

//添加粉丝
func (fan *Fan) AddFan(web string) bool {
	conn := utils.GetDb(web).Open()
	defer conn.Close()
	if conn == nil {
		return false
	}
	conn.Model(&Fan{}).Create(fan)
	return true
}
