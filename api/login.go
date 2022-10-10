package api

import (
	"blog-system/common"
	"blog-system/service"
	"fmt"
	"log"
	"net/http"
)

func (this *Api) Login(w http.ResponseWriter, req *http.Request) {
	//		接受用户名密码,返回对应的json数据
	param := common.GetRequestJsonParam(req)
	username := param["username"].(string)
	pass := param["passwd"].(string)
	log.Println("用户请求登录:账号密码为:", username, pass)
	loginResp, err := service.Login(username, pass)
	if err != nil {

		common.JsonErrorResp(w, err)
		return
	}
	common.JsonSuccessResp(w, loginResp)
	fmt.Println(param)
}
