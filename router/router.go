package router

import (
	"blog-system/api"
	"blog-system/views"
	"net/http"
)

//1页面 2数据 3静态资源
func Router() {
	//路由和执行的方法进行匹配 页面
	http.HandleFunc("/", views.HTML.Index)

	//静态文件的加载配置
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))

	//取url的参数进行分类查询
	http.HandleFunc("/c/", views.HTML.Category)

	//登录路由
	http.HandleFunc("/login", views.HTML.Login)

	//登录请求处理路由
	http.HandleFunc("/api/v1/login", api.API.Login)

	//文章详情页路由
	http.HandleFunc("/p/", views.HTML.Detail)

	//写作页面加载路由
	http.HandleFunc("/writing", views.HTML.Writing)

	//发布文章路由
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
}
