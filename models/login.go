package models

//登录成功后的返回数据
type LoginRes struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
