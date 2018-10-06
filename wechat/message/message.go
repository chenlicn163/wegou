package message

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
)

func Text(ctx *core.Context, content string) response.Text {
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

func Image(ctx *core.Context, mediaId string) response.Image {
	image := response.Image{
		MsgHeader: core.MsgHeader{
			FromUserName: ctx.MixedMsg.MsgHeader.ToUserName,
			ToUserName:   ctx.MixedMsg.MsgHeader.FromUserName,
			CreateTime:   ctx.MixedMsg.MsgHeader.CreateTime,
			MsgType:      response.MsgTypeImage,
		},
		//Image: struct{ MediaId string }{MediaId: mediaId},
	}
	image.Image.MediaId = mediaId

	return image
}
