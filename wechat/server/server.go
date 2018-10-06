package server

import (
	"net/http"
	"wegou/config"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	config := config.GetDbConfig()
	srv := core.NewDefaultAccessTokenServer(config.AppId, config.AppSecret, nil)
	clt := core.NewClient(srv, nil)
	GetMaterialCount(clt)
	/*r.ParseForm()
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
