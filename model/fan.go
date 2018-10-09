package model

import "wegou/database"

type Fan struct {
	Id        int    `json:"id"`
	Wx        string `json:"wx"`
	Nickname  string `json:"nickname"`
	CreatedAt int64  `json:"created_time"`
	UpdatedAt int64  `json:"updated_time"`
	AccountId int    `json:"account_id"`
	Status    int    `json:"status"`
}

func (fan *Fan) GetFan(web string, page int) []Fan {
	pageSize := database.FanPageSize
	offset := pageSize * (page - 1)

	conn := database.Open(web)
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

	conn := database.Open(web)
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
	conn := database.Open(web)
	defer conn.Close()
	conn.Model(&Fan{}).Create(fan)
	return true
}
