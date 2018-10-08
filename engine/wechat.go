package engine

import (
	"wegou/engine/admin"
	"wegou/service/wechat/message"

	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
)

func Run() {

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
			text := message.Text(ctx, "您输入声音")
			ctx.RawResponse(text)
		})

		mux.MsgHandleFunc(request.MsgTypeText, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			text := message.Text(ctx, "您输入了文本")
			ctx.RawResponse(text)
		})
		mux.MsgHandleFunc(request.MsgTypeVoice, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			//text := message.Text(ctx, "您输入了图片")
			//ctx.RawResponse(text)

			image := message.Image(ctx, "_c3Pe6DMtXU-zedUeoeuZgG_RuXQEgwAjIAfTyCzSd8")
			//fmt.Println(image)
			ctx.RawResponse(image)
		})
		mux.EventHandleFunc(request.EventTypeSubscribe, func(ctx *core.Context) { // 设置具体类型的事件处理 Handler
			// TODO: 事件处理逻辑
			text := message.Text(ctx, "欢迎关注")
			ctx.RawResponse(text)
		})
	}

	//读取服务器配置
	wechatConfig := GetWechatConfig()
	// 创建 Server, 设置正确的参数.
	// 通常一个 Server 对应一个公众号, 当然一个 Server 也可以对应多个公众号, 这个时候 oriId 和 appId 都应该设置为空值!
	//srv := core.NewServer("{oriId}", "{appId}", " {token}", "{base64AESKey}", mux, nil)
	srv := core.NewServer(wechatConfig.OriId, wechatConfig.AppId, wechatConfig.Token, wechatConfig.AesKey, mux, nil)
	r := gin.Default()
	r.GET("/wechat_callback", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request, nil)
	})

	management := r.Group("/admin")
	management.GET("/material/:web", admin.ListMaterialServe)
	management.DELETE("/material/:web/:id", admin.DeleteMaterialServe)
	management.PUT("/material/:web", admin.AddMaterialServe)
	management.GET("/test", admin.AddFileServe)

	webConfig := GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	r.Run(addr)
}

func Test() {

}
