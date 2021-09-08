package main

import (
	"fmt"
	"net"
	"chatRoom/common/message"
	"encoding/binary"
	"io"
	"encoding/json"
)

// 处理客户端发送来的消息
func process(conn net.Conn) {
	// 1. 先延时关闭连接， 以方后面出问题
	defer conn.Close()
	// 2. for循环读取客户端的消息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发生了灵异事件，我也溜了！！！")
				return
			}else {
				fmt.Println("reakPkg err = ", err)
				return
			}
		}
		fmt.Println("客户端发来的消息为：", mes)
	}
}

// 反序列化消息的函数
func readPkg(conn net.Conn) (mes message.Message, err error){
	// 1. 先创建接收数据的buffer
	buf := make([]byte, 8192)
	// 2. 读取前4个字节获取消息的长度
	fmt.Println("读取客户端发送过来的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	// 3.	读取 [0:pkgLen]个字节 获取消息的内容, 并比较 pkgLen 和 真正获取的消息的长度，如果不一致先返回错误
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen]) // 这里如果在读数据时conn断开 则err 为io.EOF
	if n != int(pkgLen) || err != nil {
		return
	}
	// 4. 将获取到的消息进行反序列化
	// 技术就是一层窗户纸 这里的mes要加 &，  自己理解： 因为结构体为只拷贝所以并不会修改原来的值， 只有通过地址来修改
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
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
