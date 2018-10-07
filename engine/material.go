package engine

import (
	"encoding/json"
	"net/http"
	"wegou/wechat/service/server"

	"github.com/gorilla/mux"
)

//查询永久素材
func ListMaterialServe(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	page := query.Get("page")
	MaterialType := query.Get("material")
	sourceType := query.Get("source")
	status := query.Get("status")

	materials := server.GetMaterial(page, MaterialType, sourceType, status)
	DataJson := DataJson{
		Code:    "0",
		Message: "success",
		Data:    materials,
	}

	jsonStr, _ := json.Marshal(DataJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)

	/*wechatConfig := GetWechatConfig()
	srv := core.NewDefaultAccessTokenServer(wechatConfig.AppId, wechatConfig.AppSecret, nil)
	clt := core.NewClient(srv, nil)
	r.ParseForm()
	materialType := r.Form.Get("type")
	page := r.Form.Get("page")
	if materialType != "" {
		rslt := server.BatchFetch(clt, materialType, page)
		fmt.Println(rslt)
		//materialCount := GetMaterialCount(clt)
	}*/

}

//删除永久素材
func DeleteMaterialServe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	flag := server.DelMaterial(id, false)

	StatusJson := StatusJson{}
	if !flag {
		StatusJson.Code = "20001"
		StatusJson.Message = "删除失败"
	} else {
		StatusJson.Code = "0"
		StatusJson.Message = "success"
	}

	jsonStr, _ := json.Marshal(StatusJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonStr))
}
