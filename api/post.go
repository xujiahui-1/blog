package api

import (
	"blog-system/common"
	"blog-system/models"
	"blog-system/utils"
	"errors"
	"net/http"
	"time"
)

//写文章
func (*Api) SaveAndUpdatePost(w http.ResponseWriter, req *http.Request) {
	//获取userId,判断用户是否登录
	token := req.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		common.JsonErrorResp(w, errors.New("登录已过期"))
		return
	}
	uid := claims.Uid
	method := req.Method
	switch method {
	case http.MethodPost:
		//新增
		param := common.GetRequestJsonParam(req)
		cid := param["categoryId"].(int)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		postType := param["type"].(string)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cid,
			UserId:     uid,
			ViewCount:  -1,
			Type:       postType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		//	TODO 调用service层向数据库存储数据
		//service.SavePost(post)
		common.JsonSuccessResp(w, post)
	case http.MethodPut:
		//更新
	}
}
