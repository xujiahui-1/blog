package models

import (
	"blog-system/config"
	"html/template"
	"time"
)

/**
文章相关
*/

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"` //自定义页面path
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryId int       `json:"categoryId"`
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       string    `json:"type"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

//接受数据库查询出来的数据,便于页面展示
type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"` //自定义页面path
	Content      template.HTML `json:"content"`
	Markdown     string        `json:"markdown"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int
	UserName     string `json:"userName"`
	ViewCount    int    `json:"viewCount"`
	Type         string `json:"type"`
	CreateAt     string `json:"createAt"`
	UpdateAt     string `json:"updateAt"`
}

//请求数据库的查询条件的封装
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"` //自定义页面path
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int
	Type       int `json:"type"`
}

//搜索相关返回
type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

//文章相关返回 //文章详情
type PostRes struct {
	config.Viewer
	config.SystemCongfig
	Article PostMore
}

//写文章
type WritingRes struct {
	Title     string
	CdnURL    string
	Categorys []Category
}
