package processdata

import (
	"chatRoom/common/message"
	"chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	// 暂时不需要字段。。。
}

// 转发群消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 1. 反序列化消息获取到发消息人的userId
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err = ", err)
		return
	}
	// 2. 循环服务器在线用户列表， 并过滤掉自己
	for k, v := range userMgr.onlineUsers {
		if smsMes.UserId == k {
			continue
		}
		// 3. 将消息发送出去
		this.SendMesToEachOnlineUser(mes, v.Conn)
	}
	return
}

// 真正转发群消息
func (this *SmsProcess) SendMesToEachOnlineUser(mes *message.Message, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(mes)
	fmt.Println("转发消息失败：tf.WritePkg(mes) err = ", err)
}
