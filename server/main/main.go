package main

import (
	"fmt"
	"net"
)

/* // 读取并解析消息
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
func writePkg(conn net.Conn, resMes *message.Message) (err error) {
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err = ", err)
		return
	}

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
	fmt.Printf("服务器发送的消息长度 = %d， 消息内容 = %s\n", len(data), string(data))

	return
}

// 处理登录消息
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 对消息体反序列化后进行判断
	// (1) 反序列化获取loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) err = ", err)
		return
	}

	// (2) 判断登录信息是否正确,并将需要返回的消息整理好(判断前先声明要用的返回消息的结构体)
	var resMes message.Message
	var loginResMes message.LoginResMes
	resMes.Type = message.LoginMesType
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 100
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "登录用户不存在或密码错误，请注册或重新登录！"

	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err = ", err)
		return
	}
	resMes.Data = string(data)
	// (3) 将写好的消息序列化后发出去
	err = writePkg(conn, &resMes)
	if err != nil {
		fmt.Println("writePkg(conn, &resMes) err = ", err)
	}
	return
}

// 处理注册消息
func serverProcessRegister(conn net.Conn, mes *message.Message) (err error) {
	return
}

// 获取的消息类型， 根据消息类型做出相应的处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 做登录处理
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		err = serverProcessRegister(conn, mes)
		// 做注册处理
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}
*/
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
