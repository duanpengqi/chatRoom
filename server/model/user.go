package model

// 定义用户的结构体
// 为了序列化和反序列化成功，要将tag中的字段和redis中的key保持一致
type User struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}