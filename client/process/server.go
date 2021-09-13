package processdata

import(
	"fmt"
	"os"
	"net"
	"chatRoom/client/utils"
)

// 登陆成功后显示该页面
func showMenu() {
	// 这里需要for循环展示用户要选择的列表
	for {
		fmt.Println("-----------------恭喜xxx登陆成功-----------------")
		fmt.Println("                  1. 显示在线用户")
		fmt.Println("                  2. 发送消息")
		fmt.Println("                  3. 信息列表")
		fmt.Println("                  4. 退出系统")
		fmt.Print("请选择（1-4）：")
		var key int
		fmt.Scan(&key)
		switch key {
			case 1:
				fmt.Println("显示用户在新列表~")
			case 2:
				fmt.Println("发送消息~")
			case 3:
				fmt.Println("获取消息列表~")
			case 4:
				fmt.Println("退出了系统~")
				os.Exit(0)
			default :
				fmt.Println("你输入的选项不正确~")
		}	
	}
}

	// 开启一个偷偷工作的携程来保持获得服务器推送过来的消息、
func serverProcessMes(conn net.Conn){
	// 创建一个Transfer实例用来读取消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err = ", err)
			return
		}
		// 读取到消息后还要做下一步的处理
		fmt.Println("读到来自服务器的推送： mes = ", mes)
	}
}
