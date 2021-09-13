package processdata

import (
	"fmt"
	"net"
	"chatRoom/server/utils"
	"chatRoom/common/message"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
}

// 处理登录消息
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
		loginResMes.Code = 200
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(&resMes)
	if err != nil {
		fmt.Println("writePkg(&resMes) err = ", err)
	}
	return
}

// 处理注册消息
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	return
}