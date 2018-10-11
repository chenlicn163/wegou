package types

const (
	MaterialPageSize = 20
	FanPageSize      = 20

	WechatSuccessCode = "0"
	WechatSuccessMsg  = "success"

	//web管理
	WebFiledCode = "10001"
	WebFiledMsg  = "公众号错误"

	//永久资源管理
	MaterialDelFailedCode           = "20001"
	MaterialDelFailedMsg            = "删除永久素材失败"
	SourceTypeErrorCode             = "20002"
	SourceTypeErrorMsg              = "资源类型参数错误"
	SourceStatusErrorCode           = "20003"
	SourceStatusErrorMsg            = "素材状态参数错误"
	MaterialTypeErrorCode           = "20004"
	MaterialTypeErrorMsg            = "素材类型参数错误"
	MaterialTitleAddFailedCode      = "20005"
	MaterialTitleAddFailedMsg       = "新增加永久素材失败,请输入资源名称"
	MaterialFileAddFailedCode       = "20006"
	MaterialFileAddFailedMsg        = "新增加永久素材失败,图片上传失败"
	MaterialSourceTypeAddFailedCode = "20007"
	MaterialSourceTypeAddFailedMsg  = "新增加永久素材失败,资源类型不正确"
	MaterialIdDeleteFailedCode      = "20008"
	MaterialIdDeleteFailedMsg       = "删永久素材失败，资源ID不正确"

	//粉丝管理
	FanAddFailedCode = "20101"
	FanAddFailedMsg  = "粉丝添加失败"
)
