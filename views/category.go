package views

import (
	"blog-system/common"
	"blog-system/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//分类处理器
func (this *HTMLApi) Category(w http.ResponseWriter, req *http.Request) {
	categoryTemplate := common.Template.Category //获取到index页面的模板对象
	//取到url的请求参数
	path := req.URL.Path
	str := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("strconv.Atoi(str)出错了", err)
	}

	//文章数据,分页
	err = req.ParseForm()

	pageStr := req.Form.Get("page") //页码
	pageSize := 10
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	categoryResponse, err := service.GetPostByCategoryId(cid, page, pageSize)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	fmt.Println(categoryResponse.CategoryName)
	categoryTemplate.WriteData(w, categoryResponse)
}
