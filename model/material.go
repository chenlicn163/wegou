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

//从数据库中获取素材
func (material *Material) GetMaterial(web string, page int, MaterialType string, sourceType string, status int) []Material {

	pageSize := database.MaterialPageSize
	offset := pageSize * (page - 1)

	conn := database.Open(web)
	defer conn.Close()
	if conn == nil {
		return nil
	}

	var materials []Material
	materialConn := conn.Model(&Material{}).
		Offset(offset).Limit(pageSize)

	//素材类型 临时素材（1）、永久素材（2）
	if MaterialType != "" {
		materialConn = materialConn.Where("material_type = ?", MaterialType)
	}

	//资源类型
	//图片（image）	2M，支持bmp/png/jpeg/jpg/gif格式
	//语音（voice）	2M，播放长度不超过60s，mp3/wma/wav/amr格式
	//视频（video）	10MB，支持MP4格式
	//缩略图（thumb） 64KB，支持JPG格式
	//图文（news）	当资源类型为图文时，素材类型只能是永久素材
	if sourceType != "" {
		materialConn = materialConn.Where("source_type = ?", sourceType)
	}

	//素材状态 正常状态（1）、添加状态（2）、删除状态（3）
	if status != 0 {
		materialConn = materialConn.Where("status = ?", status)
	}

	materialConn.Find(&materials)

	return materials

}

//添加素材
func (material *Material) AddMaterial(web string) bool {
	conn := database.Open(web)
	defer conn.Close()
	conn.Model(&Material{}).Create(material)
	return true
}

//更新素材
func (material *Material) UpdateMaterial(web string) bool {
	conn := database.Open(web)
	defer conn.Close()
	conn.Model(&Material{}).Where("id=?", material.Id).Updates(material)
	return true
}

//删除素材
func (material *Material) DelMaterial(web string) bool {
	conn := database.Open(web)
	defer conn.Close()
	conn.Model(&Material{}).Where("id=?", material.Id).Delete(Material{})
	return true
}
