package routes

import (
	"wegou/service/server"
	"wegou/service/wx"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
)

func WechatServe(web string) *core.Server {

	mux := core.NewServeMux() // 创建 core.Handler, 也可以用自己实现的 core.Handler

	// 注册消息(事件)处理 Handler, 都不是必须的!
	{
		mux.UseFunc(func(ctx *core.Context) { // 注册中间件, 处理所有的消息(事件)
			// TODO: 中间件处理逻辑

		})
		mux.UseFuncForMsg(func(ctx *core.Context) { // 注册中间件, 处理所有的消息
			// TODO: 中间件处理逻辑

		})
		mux.UseFuncForEvent(func(ctx *core.Context) { // 注册中间件, 处理所有的事件
			// TODO: 中间件处理逻辑
		})

		mux.DefaultMsgHandleFunc(func(ctx *core.Context) { // 设置默认消息处理 Handler
			// TODO: 消息处理逻辑
		})
		mux.DefaultEventHandleFunc(func(ctx *core.Context) { // 设置默认事件处理 Handler
			// TODO: 事件处理逻辑
		})

		mux.MsgHandleFunc(request.MsgTypeVoice, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			wx.GetMessage(ctx).Text(ctx.QueryParams.Get("web"), "您输入声音")
		})

		mux.MsgHandleFunc(request.MsgTypeText, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			wx.GetMessage(ctx).Text(ctx.QueryParams.Get("web"), "您输入了文本")
		})
		mux.MsgHandleFunc(request.MsgTypeVoice, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			wx.GetMessage(ctx).Image(ctx.QueryParams.Get("web"), "_c3Pe6DMtXU-zedUeoeuZgG_RuXQEgwAjIAfTyCzSd8")
		})
		mux.EventHandleFunc(request.EventTypeSubscribe, func(ctx *core.Context) { // 设置具体类型的事件处理 Handler
			// TODO: 事件处理逻辑
			server.AddFan(web, ctx.MixedMsg.MsgHeader.FromUserName)
			wx.GetMessage(ctx).Text(web, "欢迎关注")
		})
	}
	wechat, _ := server.GetWechatCache(web)

	// 创建 Server, 设置正确的参数.
	// 通常一个 Server 对应一个公众号, 当然一个 Server 也可以对应多个公众号, 这个时候 oriId 和 appId 都应该设置为空值!
	//srv := core.NewServer("{oriId}", "{appId}", " {token}", "{base64AESKey}", mux, nil)
	//logrus.Info(wechat)
	srv := core.NewServer(wechat.Oriid, wechat.Appid, wechat.Token, wechat.Aeskey, mux, nil)
	return srv
}
