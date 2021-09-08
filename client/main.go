package main

import(
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	// 接收用户的选择
	var key int
	// 判断是否继续显示菜单
	var loop bool = false

	for {
		fmt.Println("--------------欢迎来到多人聊天系统--------------")
		fmt.Println("                 1. 用户登录")
		fmt.Println("                 2. 注册用户")
		fmt.Println("                 3. 退出系统")
		fmt.Println("请输入您的选择（1-3）：")
		fmt.Scanf("%d",&key)

		switch key {
			case 1 :
				fmt.Println("登录聊天室：")
			case 2 :
				fmt.Println("注册用户：")
			case 3 :
				fmt.Println("退出系统！")
				os.Exit(0)
			default : 
				fmt.Println("输入有误, 请重新输入！")
				loop = true
		}

		// 当选择了 1或2 后进入到具体的界面
		if key == 1 {
			fmt.Print("用户Id：")
			fmt.Scan(&userId)
			fmt.Print("密码：")
			fmt.Scan(&userPwd)
			login(userId, userPwd)
		}else if key == 2 {
			fmt.Println("欢迎来到注册用户界面~")
		}
		
		if !loop {
			return
		}
	}
}