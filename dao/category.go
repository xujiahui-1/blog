package dao

import (
	"blog-system/models"
	"log"
)

//查找全部分类方法
func SelectAllCategory() ([]models.Category, error) {
	query, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("SelectAllCategory is wrong", err)
		return nil, err
	}
	var categorys []models.Category
	for query.Next() {
		var category models.Category
		err = query.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("query.Next() is wrong", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

//id查找分类名
func SelectCategoryNameById(categoryId int) string {
	row := DB.QueryRow("SELECT name FROM blog_category where cid =?", categoryId)

	var categoryName string
	row.Scan(&categoryName)
	return categoryName
}
