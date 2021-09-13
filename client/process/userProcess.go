package processdata

import (
	"chatRoom/common/message"
	"chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	// 暂时不需要字段， 我感觉可以把uerId和userPwd放进来
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
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
	var registerMes message.RegisterMes
	// 2.2 序列化结构体
	mes.Type = message.RegisterMesType
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) err = ", err)
		return
	}
	// 2.3 将序列化的消息为切片类型（[]uint8）再转换成字符串,并赋给mes中的Data字段
	mes.Data = string(data)

	// 3. 序列化（打包）并发送消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mes)

	// 4. 读取服务器返回来的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	// 5. 把读取到的数据中的Data反序列化并输出
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerResMes) err = ", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功~")
		return
	} else {
		fmt.Println(registerResMes.Error)
		return
	}

	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mes)

	// 4. 读取服务器返回来的消息
	mes, err = tf.ReadPkg()
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
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功~")
		// 1. 登录成功开启一个协程保持与服务器的连接
		go serverProcessMes(conn)
		// 2.后进入欢迎xxx登录成功界面
		showMenu()

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
