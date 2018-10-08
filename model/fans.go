package model

type Fans struct {
	id          int    `json:"id"`
	wx          string `json:"wx"`
	nickname    string `json:"nickname"`
	createdTime int    `json:"created_time"`
	updatedTime int    `json:"updated_time"`
	accountId   int    `json:"account_id"`
}
