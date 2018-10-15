package model

type Account struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Password    string   `json:"password"`
	CreatedTime int64    `json:"created_time"`
	UpdatedTime int64    `json:"updated_time"`
	Wechat      []Wechat `gorm:"ForeignKey:AccountId;AssociationForeignKey:Refer"`
}
