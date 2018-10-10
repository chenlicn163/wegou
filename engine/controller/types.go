package controller

const (
	WechatSuccessCode     = "0"
	MaterialDelFailedCode = "20001"
	MaterialDelFailedMsg  = "删除永久素材失败"
	MaterialAddFailedCode = "20002"
	MaterialAddFailedMsg  = "新增加永久素材失败"
	FanAddFailedCode      = "20101"
	FanAddFailedMsg       = "粉丝添加失败"
)

type StatusJson struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
