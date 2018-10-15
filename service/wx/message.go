package wx

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

//文本消息
func Text(web string, ctx *core.Context, content string) {
	text := response.NewText(ctx.MixedMsg.MsgHeader.FromUserName, ctx.MixedMsg.MsgHeader.ToUserName,
		ctx.MixedMsg.MsgHeader.CreateTime, content)
	sendMsg(text, ctx)
}

//图片消息
func Image(web string, ctx *core.Context, mediaId string) {
	image := response.NewImage(ctx.MixedMsg.MsgHeader.FromUserName, ctx.MixedMsg.MsgHeader.ToUserName,
		ctx.MixedMsg.MsgHeader.CreateTime, mediaId)
	image.Image.MediaId = mediaId
	sendMsg(image, ctx)
}

//发送消息
func sendMsg(msg interface{}, ctx *core.Context) {
	aesKey := string(ctx.AESKey)
	if aesKey != "" {
		ctx.AESResponse(msg, 0, "", nil)
	} else {
		ctx.RawResponse(msg)
	}

}
