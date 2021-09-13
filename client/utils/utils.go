package utils

import (
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct{
	Conn net.Conn
	Buf [8192]byte
}

// 读取并解析消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// 1. 先创建接收数据的buffer
	// buf := make([]byte, 8192)
	// 2. 读取前4个字节获取消息的长度
	fmt.Println("读取客户端发送过来的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	// 3.	读取 [0:pkgLen]个字节 获取消息的内容, 并比较 pkgLen 和 真正获取的消息的长度，如果不一致先返回错误
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen]) // 这里如果在读数据时conn断开 则err 为io.EOF
	if n != int(pkgLen) || err != nil {
		return
	}
	// 4. 将获取到的消息进行反序列化
	// 技术就是一层窗户纸 这里的mes要加 &，  自己理解： 因为结构体为只拷贝所以并不会修改原来的值， 只有通过地址来修改
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

// 序列化（打包）并发送消息
func (this *Transfer) WritePkg(resMes *message.Message) (err error) {
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err = ", err)
		return
	}

	// var buf [4]byte
	pkgLen := uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen) // binary.BigEndian.PutUint32([]byte, uint32) 实现了将uint32数字转换成字节序列

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write len err = ", err)
		return
	}
	// 3.2 发送消息体
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write data err = ", err)
		return
	}
	fmt.Printf("服务器发送的消息长度 = %d， 消息内容 = %s\n", len(data), string(data))

	return
}