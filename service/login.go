package service

import (
	"blog-system/dao"
	"blog-system/models"
	"blog-system/utils"
	"errors"
)

func Login(userName string, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "xujiahui")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		//	登录失败业务逻辑
		return nil, errors.New("账号密码不正确")
	}
	var lr = &models.LoginRes{}
	token, err := utils.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("账号密码不正确")
	}
	lr.UserInfo.UserName = user.UserName
	lr.UserInfo.Avater = user.Avater
	lr.UserInfo.Uid = user.Uid
	lr.Token = token
	return lr, nil
}
