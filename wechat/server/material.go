package server

import (
	"fmt"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"
)

func GetMaterialCount(clt *core.Client) *material.MaterialCountInfo {

	info, err := material.GetMaterialCount(clt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	return info
}

func BatchGet(clt *core.Client) {
	rslt, err := material.BatchGet(clt, "image", 0, 20)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rslt)
}
