package wx

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

type WgMessage struct {
	Ctx *core.Context
}

//发送消息
func (wgMessage *WgMessage) sendMsg(content interface{}) {
	aesKey := string(wgMessage.Ctx.AESKey)
	if aesKey != "" {
		wgMessage.Ctx.AESResponse(content, 0, "", nil)
	} else {
		wgMessage.Ctx.RawResponse(content)
	}
}

//文本消息
func (wgMessage *WgMessage) Text(web string, content string) {
	text := response.NewText(wgMessage.Ctx.MixedMsg.MsgHeader.FromUserName, wgMessage.Ctx.MixedMsg.MsgHeader.ToUserName,
		wgMessage.Ctx.MixedMsg.MsgHeader.CreateTime, content)
	wgMessage.sendMsg(text)
}

//图片消息
func (wgMessage *WgMessage) Image(web string, mediaId string) {
	image := response.NewImage(wgMessage.Ctx.MixedMsg.MsgHeader.FromUserName, wgMessage.Ctx.MixedMsg.MsgHeader.ToUserName,
		wgMessage.Ctx.MixedMsg.MsgHeader.CreateTime, mediaId)
	image.Image.MediaId = mediaId
	wgMessage.sendMsg(image)
}
