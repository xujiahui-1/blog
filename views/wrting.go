package views

import (
	"blog-system/common"
	"blog-system/service"
	"net/http"
)

//写文章
func (this *HTMLApi) Writing(w http.ResponseWriter, req *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
