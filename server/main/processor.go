package main

import (
	"fmt"
	"net"
	processdata"chatRoom/server/process"
	"chatRoom/server/utils"
	"chatRoom/common/message"
	"io"
)

//创建一个processor的结构体
type Processor struct {
	Conn net.Conn
}

// 根据消息类型做出相应的处理（这里反序列化将消息类型获取出来）
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	up := &processdata.UserProcess{
		Conn: this.Conn,
	}
	switch mes.Type {
	case message.LoginMesType:
		// 做登录处理
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		err = up.ServerProcessRegister(mes)
		// 做注册处理
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

// 具体处理客户端发送来消息（这里的消息完全还没有被处理）
func (this *Processor) ProcessDetial() (err error) {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		for {
			mes, err := tf.ReadPkg()
			if err != nil {
				if err == io.EOF {
					fmt.Println("发生了灵异事件，我也溜了！！！")
					return err
				} else {
					fmt.Println("reakPkg err = ", err)
					return err
				}
			}
			fmt.Println("客户端发来的消息为：", mes)
			// fmt.Println("客户端发来的消息类型为：", mes.Type)
			// fmt.Printf("mes.Data==============类型%T， 值为%v\n", mes.Data, mes.Data)

			// 3. 处理用户发送来的消息
			err = this.ServerProcessMes(&mes)
			if err != nil {
				fmt.Println("serverProcessMes() 处理消息失败：", err)
				return err
			}
	}
}