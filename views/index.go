package views

import (
	"blog-system/common"
	"blog-system/service"
	"fmt"
	"net/http"
	"strconv"
)

//index模板的启动运行
func (this *HTMLApi) Index(w http.ResponseWriter, req *http.Request) {
	index := common.Template.Index //获取到index页面的模板对象
	//文章数据,分页
	err := req.ParseForm()
	pageStr := req.Form.Get("page") //页码
	pageSize := 10
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	hr, err := service.GetIndexInfo(page, pageSize) //获取首页数据

	if err != nil {
		fmt.Println("Index() is wrong", err)
	}
	index.WriteData(w, hr) //执行

}
