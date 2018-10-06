package server

import (
	"fmt"
	"net/http"
	"wegou/config"

	"gopkg.in/chanxuehong/wechat.v2/mp/material"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

func MaterialServe(w http.ResponseWriter, r *http.Request) {
	wechatConfig := config.GetWechatConfig()
	srv := core.NewDefaultAccessTokenServer(wechatConfig.AppId, wechatConfig.AppSecret, nil)
	clt := core.NewClient(srv, nil)
	r.ParseForm()
	materialType, foundType := r.Form["type"]
	if foundType {
		rslt, err := material.BatchGet(clt, materialType[0], 0, 20)
		if err != nil {
			fmt.Println("type is error")
		}
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
