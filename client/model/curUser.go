package model

import (
	"chatRoom/common/message"
	"net"
)

// 在初始化后可以在很多地方都用到，因此可以声明一个全局的变量
type CurUser struct {
	Conn net.Conn
	message.User
}
