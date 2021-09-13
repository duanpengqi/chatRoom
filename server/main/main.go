package main

import (
	"fmt"
	"net"
)

// 处理客户端发送来的消息
func process(conn net.Conn) {
	// 1. 先延时关闭连接， 以方后面出问题
	defer conn.Close()
	// 2. for循环读取客户端的消息
	processor := &Processor{
		Conn: conn,
	}
	err := processor.ProcessDetial()
	if err != nil {
		fmt.Println("客户端和服务器的协程出现错误 err = ", err)
		return
	}
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
