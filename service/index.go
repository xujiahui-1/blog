package service

import (
	"blog-system/config"
	"blog-system/dao"
	"blog-system/models"
	"html/template"
)

//获取首页所有信息
func GetIndexInfo(page int, pageSize int) (*models.HomeResponse, error) {
	//分类数据
	categorys, err := dao.SelectAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.SelectAllPost(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
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
		postMores = append(postMores, postMore)
	}
	total := dao.GetPostCount()
	//计算分多少页
	pageConunt := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pageConunt; i++ {
		pages = append(pages, i+1)
	}
	hr := &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Page:      page,
		Posts:     postMores,
		Total:     total,
		Pages:     pages,
		PageEnd:   page != pageConunt,
	}
	return hr, nil
}
