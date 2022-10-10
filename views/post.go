package views

import (
	"blog-system/common"
	"blog-system/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//文章详情页,detail页面渲染
func (*HTMLApi) Detail(w http.ResponseWriter, req *http.Request) {
	detail := common.Template.Detail
	//获取路径参数
	path := req.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)

	if err != nil {
		fmt.Println("strconv.Atoi(str)出错了", err)
	}
	//使用pid调用service中方法，得到返回数据并赋值给postRes
	postRes, err := service.GetDetailPost(pid)
	if err != nil {
		detail.WriteData(w, errors.New("查询文章详情出错"))
		return
	}

	detail.WriteData(w, postRes)
}
