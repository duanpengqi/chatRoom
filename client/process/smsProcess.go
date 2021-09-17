package processdata

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct{}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 将消息整理好发送出去
	// 1. 创建发送消息的结构体实例
	var mes message.Message
	var smsMes message.SmsMes

	mes.Type = message.SmsMesType
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	smsMes.Content = content

	// 2. 序列化smsMes,并赋值给mes.Data
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes -> json.Marshal(smsMes) err = ", err)
		return
	}
	mes.Data = string(data)
	// 3. 将整理好的消息发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("SendGroupMes -> tf.WritePkg(&mes) err = ", err)
	}

	return
}
