package model

//用户实体
type Account struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Password    string   `json:"password"`
	CreatedTime int64    `json:"created_time"`
	UpdatedTime int64    `json:"updated_time"`
	LoginTime   int64    `json:"login_time"`
	Status      int      `json:"status"`
	Wechats     []Wechat `gorm:"ForeignKey:AccountId;AssociationForeignKey:Refer"`
}
