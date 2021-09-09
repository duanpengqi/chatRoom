package main

import (
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 读取并解析消息
func readPkg(conn net.Conn) (mes message.Message, err error) {
	// 1. 先创建接收数据的buffer
	buf := make([]byte, 8192)
	// 2. 读取前4个字节获取消息的长度
	fmt.Println("读取客户端发送过来的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	// 3.	读取 [0:pkgLen]个字节 获取消息的内容, 并比较 pkgLen 和 真正获取的消息的长度，如果不一致先返回错误
	pkgLen := binary.BigEndian.Uint32(buf[0:4])
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

// 序列化（打包）并发送消息
func writePkg(conn net.Conn, sendMes *message.Message) (err error) {
	data, err := json.Marshal(sendMes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ", err)
		return
	}
	// 3. 将处理过后的数据发送给服务器（先发送消息长度，再发送消息体）

	// 3.1 发送消息长度
	//		由于conn.Write方法需要的参数是一个byte切片所以这里要把发送的信息进行转换    Write(b []byte) (n int, err error)
	// 如何把消息的长度转换成切片？？？
	// binary -> ByteOrder接口中的方法 实现了数字 与 字节的转化
	var buf [4]byte
	pkgLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[:4], pkgLen) // binary.BigEndian.PutUint32([]byte, uint32) 实现了将uint32数字转换成字节序列

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write len err = ", err)
		return
	}
	// 3.2 发送消息体
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write data err = ", err)
		return
	}
	fmt.Printf("客户端发送的消息长度 = %d， 消息内容 = %s\n", len(data), string(data))

	return
}
