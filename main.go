package main

import (
	"blog-system/common"
	"blog-system/router"
	"fmt"
	"net/http"
)

//初始化模板加载
func init() {
	common.LoadTemplate()
	router.Router() //路由
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
