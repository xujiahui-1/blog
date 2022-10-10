package dao

import (
	"blog-system/models"
	"errors"
	"log"
)

//查询所有文章
func SelectAllPost(page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	query, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("SelectAllPost() is wrong", err)
		return nil, err
	}
	var posts []models.Post
	for query.Next() {
		var post models.Post
		err := query.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("query.Next() is wrong", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//查询文章总数
func GetPostCount() int {
	row := DB.QueryRow("SELECT count(*) FROM blog_post")
	var postCountNum int
	err := row.Scan(&postCountNum)
	if err != nil {
		log.Println("GetPostCount() is wrong", err)
	}
	return postCountNum
}

//查询所有文章分页
func SelectAllPostByCategoryId(cid int, page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	query, err := DB.Query("select * from blog_post where category_id =? limit ?,?", cid, page, pageSize)
	if err != nil {
		log.Println("SelectAllPost() is wrong", err)
		return nil, err
	}
	var posts []models.Post
	for query.Next() {
		var post models.Post
		err := query.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("query.Next() is wrong", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

//文章详情
func GetPostById(pid int) (*models.Post, error) {
	row := DB.QueryRow("SELECT * FROM blog_post where pid=?", pid)

	var post models.Post
	if row.Err() != nil {
		log.Println("GetPostById(pid int)错误", row.Err())
		return nil, errors.New("查询文章详情出错")
	}
	err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		log.Println("GetPostById(pid int)错误", err)
		return nil, errors.New("查询文章详情出错")
	}
	return &post, nil
}
