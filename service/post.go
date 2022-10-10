package service

import (
	"blog-system/config"
	"blog-system/dao"
	"blog-system/models"
	"fmt"
	"html/template"
)

//文章详情 models.Post
func GetDetailPost(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)
	categoryName := dao.SelectCategoryNameById(post.CategoryId)
	userName := dao.SelectUserNameById(post.UserId)
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		Markdown:     post.Markdown,
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	if err != nil {
		return nil, err
	}
	return postRes, nil
}

//写文章页面渲染
func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.SelectAllCategory()
	if err != nil {
		fmt.Println(err)
	}
	wr.Categorys = category
	return
}

//写文章
func SavePost() {

}
