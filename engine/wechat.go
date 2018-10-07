package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wegou/wechat/model"
	"wegou/wechat/service/message"
	"wegou/wechat/service/server"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
)

func Run() {

	//读取服务器配置
	wechatConfig := GetWechatConfig()
	fmt.Println(wechatConfig)

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
		mux.MsgHandleFunc(request.MsgTypeText, func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
			// TODO: 消息处理逻辑
			//text := message.Text(ctx, "您输入了图片")
			//ctx.RawResponse(text)

			image := message.Image(ctx, "_c3Pe6DMtXU-zedUeoeuZgG_RuXQEgwAjIAfTyCzSd8")
			fmt.Println(image)
			ctx.RawResponse(image)
		})
		mux.EventHandleFunc(request.EventTypeSubscribe, func(ctx *core.Context) { // 设置具体类型的事件处理 Handler
			// TODO: 事件处理逻辑
			text := message.Text(ctx, "欢迎关注")
			ctx.RawResponse(text)
		})
	}

	// 创建 Server, 设置正确的参数.
	// 通常一个 Server 对应一个公众号, 当然一个 Server 也可以对应多个公众号, 这个时候 oriId 和 appId 都应该设置为空值!
	//srv := core.NewServer("{oriId}", "{appId}", " {token}", "{base64AESKey}", mux, nil)
	srv := core.NewServer(wechatConfig.OriId, wechatConfig.AppId, wechatConfig.Token, wechatConfig.AesKey, mux, nil)

	// 在回调 URL 的 Handler 里处理消息(事件)
	http.HandleFunc("/wechat_callback", func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r, nil)
	})
	http.HandleFunc("/material/info", MaterialServe)

	webConfig := GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	http.ListenAndServe(addr, nil)
}

//处理永久素材
func MaterialServe(w http.ResponseWriter, r *http.Request) {
	wechatConfig := GetWechatConfig()
	srv := core.NewDefaultAccessTokenServer(wechatConfig.AppId, wechatConfig.AppSecret, nil)
	clt := core.NewClient(srv, nil)
	r.ParseForm()
	materialType := r.Form.Get("type")
	page := r.Form.Get("page")
	if materialType != "" {
		rslt := server.BatchFetch(clt, materialType, page)
		fmt.Println(rslt)
		//materialCount := GetMaterialCount(clt)
	}

	/*
		signature, foundSignature := r.Form["signature"]
		timestamp, foundTimestamp := r.Form["timestamp"]
		nonce, foundNonce := r.Form["nonce"]

		if foundSignature && foundTimestamp && foundNonce {
			fmt.Println(signature, timestamp, nonce)
			fmt.Println("校验")
		} else {
			fmt.Println(foundSignature, signature, "响应")
		}*/

}

func Test() {
	mat := model.Material{}
	materials := mat.GetMaterial(1)

	/*var jsonMaterials []string
	for _, material := range materials {
		jsonStr, _ := json.Marshal(material)
		jsonMaterials = append(jsonMaterials, string(jsonStr))
	}*/
	//fmt.Println(materials)
	result, _ := json.Marshal(materials)
	fmt.Printf("%s", result)

}
