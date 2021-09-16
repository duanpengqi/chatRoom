package processdata

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 20)

func updateUserStatus(data string) {
	// 1. 先解析数据
	var notifyUserStatusMes message.NotifyUserStatusMes
	err := json.Unmarshal([]byte(data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(data), &notifyUserStatusMes) err = ", err)
		return
	}
	// 2. 将数据添加到客户端管理的状态表中
	// var user message.User
	// user.UserId = notifyUserStatusMes.UserId
	// user.UserStatus = notifyUserStatusMes.Status
	// onlineUsers[notifyUserStatusMes.UserId] = &user
	// 优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	// 3. 重新展示在线用户
	outputOnlineuser()
}

// 展示在线用户
func outputOnlineuser() {
	fmt.Println("当前在线用户：")
	for id, _ := range onlineUsers {
		fmt.Println(id)
	}
}
