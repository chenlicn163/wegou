package wx

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

type Message struct {
	Ctx *core.Context
}

//发送消息
func (msg *Message) sendMsg(content interface{}) {
	aesKey := string(msg.Ctx.AESKey)
	if aesKey != "" {
		msg.Ctx.AESResponse(msg, 0, "", nil)
	} else {
		msg.Ctx.RawResponse(msg)
	}
}

//文本消息
func (msg Message) Text(web string, content string) {
	text := response.NewText(msg.Ctx.MixedMsg.MsgHeader.FromUserName, msg.Ctx.MixedMsg.MsgHeader.ToUserName,
		msg.Ctx.MixedMsg.MsgHeader.CreateTime, content)
	msg.sendMsg(text)
}

//图片消息
func (msg *Message) Image(web string, mediaId string) {
	image := response.NewImage(msg.Ctx.MixedMsg.MsgHeader.FromUserName, msg.Ctx.MixedMsg.MsgHeader.ToUserName,
		msg.Ctx.MixedMsg.MsgHeader.CreateTime, mediaId)
	image.Image.MediaId = mediaId
	msg.sendMsg(image)
}
