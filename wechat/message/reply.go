package message

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

func Text(ctx *core.Context,content string) response.Text{
	text := response.Text{
		MsgHeader: core.MsgHeader{
			FromUserName: ctx.MixedMsg.MsgHeader.ToUserName,
			ToUserName:   ctx.MixedMsg.MsgHeader.FromUserName,
			CreateTime:   ctx.MixedMsg.MsgHeader.CreateTime,
			MsgType:      response.MsgTypeText,
		},
		Content: content,
	}
	return text
}

