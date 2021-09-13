package message

// 定义一个用户的结构体，用来接收redis中的用户数据信息
// 为了序列化和反序列化成功， 结构体字段对应的tag名字 应和 redis的user的key中对应的名字保持一致
type User struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
