package model

import (
	"wegou/database"
)

//素材
type Material struct {
	Id               int    `json:"id"`
	Pid              int    `json:"pid"`
	Title            string `json:"title"`
	Pic              string `json:"pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
	ThumbMediaId     string `json:"thumb_media_id"`
	MediaId          string `json:"media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	MaterialType     int    `json:"material_type"`
	AccountId        int    `json:"account_id"`
	Status           int    `json:"status"`
	SourceType       string `json:"source_type"`
}

//从数据库中获取永久素材，数据库中不存在，提示稍后查看
func (material *Material) GetMaterial(page int) []Material {

	pageSize := database.MaterialPageSize
	offset := pageSize * (page - 1)

	conn := database.Open()
	if conn == nil {
		return nil
	}

	var materials []Material
	conn.Model(&Material{}).Offset(offset).Limit(pageSize).Find(&materials)

	return materials

}

//添加永久素材
func (material *Material) AddMaterial() bool {
	//1判断限制
	//2上传素材到服务器，更新数据库
	//3触发任务，定时处理，把数据更新到微信服务器

	return true
}

func (material *Material) DelMaterial() bool {
	//1标记数据库需要删除
	//2触发任务，定时处理，删除微信服务器数据

	return true
}
