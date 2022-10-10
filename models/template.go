package models

import (
	"fmt"
	"html/template"
	"io"
	"time"
)

/**
模板相关操作
*/
type TemplateBlog struct {
	*template.Template
}
type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

//执行的封装
func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

//初始化方法 里面调用模板加载
func InitTemplate(templateDir string) HtmlTemplate {
	tp := LoadTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	htmlTemplate := HtmlTemplate{
		Index:      tp[0],
		Category:   tp[1],
		Custom:     tp[2],
		Detail:     tp[3],
		Login:      tp[4],
		Pigeonhole: tp[5],
		Writing:    tp[6],
	}

	return htmlTemplate
}

//加载模板的封装
func LoadTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		//获得模板对象
		t := template.New(viewName)
		//这里映射后端的方法到页面中
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		//使用模板解析文件,因为各模板中的数据涉及到嵌套，所以要都解析了才行
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		post := templateDir + "layout/post-list.html"
		personal := templateDir + "layout/personal.html"
		pagination := templateDir + "layout/pagination.html"
		files, err := t.ParseFiles(templateDir+viewName, home, header, footer, post, personal, pagination)
		if err != nil {
			fmt.Println("解析模板出错", err)
		}
		var tb TemplateBlog
		tb.Template = files
		tbs = append(tbs, tb)

	}
	return tbs
}

//判断偶数方法
func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(str []string, index int) string {
	return str[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
