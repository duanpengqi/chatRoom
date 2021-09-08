package main

import (
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(userId int, userPwd string) (err error) {
	// 1. 连接到服务器， 并延时关闭
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn.Dial err = ", err)
		return
	}
	defer conn.Close()

	// 2. 处理用户输入的消息
	// 2.1 声明 发送的消息结构体、 登录的消息结构体
	var mes message.Message
	var loginMes message.LoginMes
	// 2.2 序列化结构体
	mes.Type = message.LoginMesType
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err = ", err)
		return
	}
	// fmt.Printf("检测序列化后的消息类型为切片：%T", data)
	// 2.3 将序列化的消息为切片类型（[]uint8）再转换成字符串,并赋给mes中的Data字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ", err)
		return
	}
	// 3. 将处理过后的数据发送给服务器（先发送消息长度，再发送消息体）
	// 3.1 因为conn.Write方法需要的参数是一个byte切片所以这里要把发送的信息进行转换    Write(b []byte) (n int, err error)
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

	fmt.Println("等10s钟我就溜了。。。")
	time.Sleep(time.Second * 10)
	fmt.Println("10s到了我溜了。。。")
	return
}
