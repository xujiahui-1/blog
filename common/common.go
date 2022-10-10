package common

import (
	"blog-system/config"
	"blog-system/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/**、
Template就是个分界线
获取Template之前，调用InitTemplate函数，获取到HtmlTemplate结构体对象，
获取到Template后,在index方法中执行
*/
var Template models.HtmlTemplate //全局调用，副的值为 HtmlTemplate结构体

func LoadTemplate() {
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		//耗时
		Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		group.Done()
	}()
	group.Wait()
}

//传入 writer和数据,帮你转化成json格式并写入resp中
func JsonSuccessResp(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	marshal, err := json.Marshal(result)
	if err != nil {
		fmt.Println(" Login()中json转化出错", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

func JsonErrorResp(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	marshal, err := json.Marshal(result)
	if err != nil {
		fmt.Println(" Login()中json转化出错", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

//获取页面传来的json参数解析，返回成
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
