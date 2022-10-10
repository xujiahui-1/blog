package dao

import (
	"blog-system/models"
	"fmt"
	"log"
)

func SelectUserNameById(userId int) string {
	row := DB.QueryRow("SELECT user_name FROM blog_user where uid =?", userId)
	if row.Err() != nil {
		log.Println("SelectUserNameById() is wrong")
	}
	var userName string
	row.Scan(&userName)
	return userName
}

//登录
func GetUser(userName string, passwd string) *models.User {
	row := DB.QueryRow("SELECT * FROM blog_user where user_name =? and passwd =? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println("SelectUserNameById() is wrong")
		return nil
	}
	var userRes models.User
	err := row.Scan(&userRes.Uid,
		&userRes.UserName,
		&userRes.Passwd,
		&userRes.Avater,
		&userRes.CreateAt,
		&userRes.UpdateAt,
	)
	if err != nil {
		fmt.Println("用户出错", err)
		return nil
	}
	return &userRes
}
