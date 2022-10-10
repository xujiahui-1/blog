package views

import (
	"blog-system/common"
	"blog-system/config"
	"net/http"
)

func (this *HTMLApi) Login(w http.ResponseWriter, req *http.Request) {
	//获取登录页模板
	login := common.Template.Login
	//不需要数据库中的数据，而是从配置文件中获取的
	login.WriteData(w, config.Cfg.Viewer)
}
