package processdata

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的界面
func showMenu() {
	var key int
	var content string
	for {
		fmt.Println("--------------欢迎xxx登录成功--------------")
		fmt.Println("                1.在线用户")
		fmt.Println("                2.发送消息")
		fmt.Println("                3.消息列表")
		fmt.Println("                4.退出系统")
		fmt.Println("请选择（1-4）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			outputOnlineuser()
		case 2:
			fmt.Println("请输入要发送的内容：")
			fmt.Scanf("%s\n", &content)
			// 创建SmsProcess实例，调用短消息发送处理函数
			sp := &SmsProcess{}
			err := sp.SendGroupMes(content)
			if err != nil {
				fmt.Println("sp.SendGroupMes(content) err = ", err)
				return
			}
		case 3:
			fmt.Println("展示消息列表~")
		case 4:
			fmt.Println("退出系统了~")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入！！！")
		}
	}
}

// 开启一个偷偷保持和服务器连接的协程，用来监听服务器发送来的推送。
func serverProcessMes(conn net.Conn) (err error) {

	// 1. 创建一个transfer
	tf := &utils.Transfer{
		Conn: conn,
	}
	// 2. 循环监听与服务器连接的通道
	for {
		// fmt.Println("================客户端正在等待服务器推送消息====================")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err = ", err)
			return err
		}
		// 判断消息类型分别处理
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 用户上线提醒
			updateUserStatus(mes.Data)
		case message.SmsMesType:
			// fmt.Println("短消息 =========== ", mes)
			// 有新的群消息
			outputGroupMes(mes.Data)
		default:
			fmt.Println("服务器返回的消息类型不匹配！")
		}
	}

}
