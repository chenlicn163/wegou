package model

import (
	"wegou/types"
	"wegou/utils"
)

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
	Headimgurl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Groupid        int    `json:"groupid"`
	TagidList      string `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
}

func (fan *Fan) GetFan(web string, page int) []Fan {
	pageSize := types.FanPageSize
	offset := pageSize * (page - 1)

	conn := utils.OpenWechat(web)
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

func (fan *Fan) GetFanCount(web string) int {

	conn := utils.OpenWechat(web)
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

func (fan *Fan) AddFan(web string) bool {
	conn := utils.OpenWechat(web)
	defer conn.Close()
	conn.Model(&Fan{}).Create(fan)
	return true
}
