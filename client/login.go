package main

import (
	"chatRoom/common/message"
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
	// 2.3 将序列化的消息为切片类型（[]uint8）再转换成字符串,并赋给mes中的Data字段
	mes.Data = string(data)

	// 3. 序列化（打包）并发送消息
	err = writePkg(conn, &mes)

	// 4. 读取服务器返回来的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	// 5. 把读取到的数据中的Data反序列化并输出
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginResMes) err = ", err)
		return
	}
	if loginResMes.Code == 100 {
		fmt.Println("登录成功~")
	} else {
		fmt.Println(loginResMes.Error)
	}

	fmt.Println("等10s钟我就溜了。。。")
	time.Sleep(time.Second * 10)
	fmt.Println("10s到了我溜了。。。")
	return
}
