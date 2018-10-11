package types

//站点配置
type Web struct {
	Port string
	Host string
}

//微信配置
type Wechat struct {
	OriId     string
	AppId     string
	Token     string
	AppSecret string
	AesKey    string
}

//数据库配置
type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type StatusJson struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Dto struct {
	Data    interface{}
	Code    string
	Message string
}
