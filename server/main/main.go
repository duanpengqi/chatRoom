package main

import (
	"chatRoom/server/model"
	"fmt"
	"net"
	"time"
)

// 处理客户端发送来的消息
func process(conn net.Conn) {
	// 1. 先延时关闭连接， 以方后面出问题
	defer conn.Close()
	// 2. for循环读取客户端的消息
	err := processDetial(conn)
	if err != nil {
		fmt.Println("processDetial(conn) err = ", err)
		return
	}
}

// 创建一个全局的UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	fmt.Println("服务器在8889端口监听...")

	// 1. 在本地监听端口8889端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889") // 0.0.0.0:8889表示在本地监听
	if err != nil {
		fmt.Println("net Listen err = ", err)
		return
	}

	// 2. 监听成功后 for循环等待客户端来连接服务器
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept() // 等待，接受客户连接
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}

		// 3. 开启一个协程，来为该客户服务
		go process(conn)
	}
}
