package server

import (
	"fmt"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"
)

func GetMaterialCount(clt *core.Client) {

	info, err := material.GetMaterialCount(clt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

}
