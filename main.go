package main

import (
	"net/http"
	"wegou/config"
	"wegou/wechat/message"
	"wegou/wechat/server"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

//1 定义配置 appid、token

//2 启动端口
const (
	port = ":8090"
)

func main() {
	config := config.GetDbConfig()

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

		mux.MsgHandleFunc("text", func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			text := message.Text(ctx, "您输入了文本")
			ctx.RawResponse(text)
		})
		mux.MsgHandleFunc("image", func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			text := message.Text(ctx, "您输入了图片")
			ctx.RawResponse(text)
		})
		mux.EventHandleFunc("subscribe", func(ctx *core.Context) { // 设置具体类型的事件处理 Handler
			// TODO: 事件处理逻辑
			text := message.Text(ctx, "欢迎关注")
			ctx.RawResponse(text)
		})
	}

	// 创建 Server, 设置正确的参数.
	// 通常一个 Server 对应一个公众号, 当然一个 Server 也可以对应多个公众号, 这个时候 oriId 和 appId 都应该设置为空值!
	//srv := core.NewServer("{oriId}", "{appId}", " {token}", "{base64AESKey}", mux, nil)
	srv := core.NewServer(config.OriId, config.AppId, config.Token, config.AesKey, mux, nil)

	// 在回调 URL 的 Handler 里处理消息(事件)
	http.HandleFunc("/wechat_callback", func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r, nil)
	})
	http.HandleFunc("/material/list", server.Serve)
	http.ListenAndServe(":80", nil)
}
