package models

import "time"

type User struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avater   string    `json:"avater"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avater   string `json:"avater"`
}
