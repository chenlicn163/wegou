package types

//站点配置
type Web struct {
	Port string
	Host string
}

//微信配置
type Wechat struct {
	Oriid      string `json:"oriid"`
	Appid      string `json:"appid"`
	Appsecret  string `json:"appsecret"`
	Token      string `json:"token"`
	Aeskey     string `json:"aeskey"`
	DbHost     string `json:"db_host"`
	DbName     string `json:"db_name"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
}

//数据库配置
type Db struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

//kafk配置
type Kafka struct {
	Blockers       []string
	CustomerTopics []string
	MaterialTopics []string
	CustomerGroup  string
	MaterialGroup  string
}

//redis配置
type Redis struct {
	Server string
	Auth   string
	Db     int
}

type Dto struct {
	Data    interface{}
	Code    string
	Message string
}
