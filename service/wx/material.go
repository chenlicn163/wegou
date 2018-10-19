package wx

import (
	"fmt"
	"strconv"
	"wegou/config"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/material"
)

type WgMaterial struct {
	Clt *core.Client
}

//获取永久素材数量
func (wgMaterial *WgMaterial) FetchMaterialCount(clt *core.Client) *material.MaterialCountInfo {

	info, err := material.GetMaterialCount(clt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	return info
}

//批量获取永久素材
func (wgMaterial *WgMaterial) BatchFetchMaterial(clt *core.Client, materialType string, pageStr string) *material.BatchGetResult {
	size := config.MaterialPageSize
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	offset := (page - 1) * size
	rslt, err := material.BatchGet(clt, materialType, offset, size)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rslt)

	return rslt
}

//永久素材，同步微信服务器数据
func (wgMaterial *WgMaterial) SyncMaterial() {
	//1 获取总数
	//2 分页获取,入库
}

//永久素材，上传到微信服务器
func (wgMaterial *WgMaterial) SyncAddMaterial() {

}

//永久素材，从微信服务器上删除
func (wgMaterial *WgMaterial) SyncDelMaterial() {

}
