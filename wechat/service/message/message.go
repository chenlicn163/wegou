package message

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

//文本消息
func Text(ctx *core.Context, content string) *response.Text {
	text := response.NewText(ctx.MixedMsg.MsgHeader.FromUserName, ctx.MixedMsg.MsgHeader.ToUserName,
		ctx.MixedMsg.MsgHeader.CreateTime, content)
	return text
}

//图片消息
func Image(ctx *core.Context, mediaId string) *response.Image {
	image := response.NewImage(ctx.MixedMsg.MsgHeader.FromUserName, ctx.MixedMsg.MsgHeader.ToUserName,
		ctx.MixedMsg.MsgHeader.CreateTime, mediaId)
	image.Image.MediaId = mediaId
	return image
}
