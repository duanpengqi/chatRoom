package processdata

import (
	"chatRoom/common/message"
	"chatRoom/server/model"
	"chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
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
	this.UserId = loginMes.UserId
	// (2) 判断登录信息是否正确,并将需要返回的消息整理好(判断前先声明要用的返回消息的结构体)
	var resMes message.Message
	var loginResMes message.LoginResMes
	resMes.Type = message.LoginResMesType
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 100
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "登录用户不存在或密码错误，请注册或重新登录！"
	// }
	// 这里需要用MyUserDao全局实例的方法进行与redis的交互
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error() // Error() 将err中的信息取出来
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误！"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功~")
		// a. 登录成功后将自己添加到在线用户列表中
		userMgr.AddOnlineUser(this)
		// b. 将自己登录的状态告诉其他用户
		this.NotifyOtheronlineUsers(loginMes.UserId)
		// c. 然后将userMgr.onlineUsers中的用户id添加到返回信息中
		for k, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, k)
		}
		fmt.Println()
		fmt.Println("userMgr.onlineUsers = ", userMgr.onlineUsers)
		fmt.Println()
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
	// 对消息体反序列化后进行判断
	// (1) 反序列化获取registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes) // 这里反序列化的是总的mes中的Data
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerMes) err = ", err)
		return
	}

	// (2) 判断登录信息是否正确,并将需要返回的消息整理好(判断前先声明要用的返回消息的结构体)
	var resMes message.Message
	var registerResMes message.RegisterResMes
	resMes.Type = message.RegisterResMesType

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 502
			registerResMes.Error = err.Error() // Error() 将err中的信息取出来
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服务器内部错误！"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功~")
	}

	// (3) 注册成功后将需要返回的消息序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) err = ", err)
		return
	}
	resMes.Data = string(data)
	// (4) 将序列化好的消息序列化后发出去
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(&resMes)
	if err != nil {
		fmt.Println("writePkg(&resMes) err = ", err)
	}
	return
}

// 当有用户状态发生变化通知其他好友
func (this *UserProcess) NotifyOtheronlineUsers(userId int) {
	// 循环告诉userMgr.onlineUsers中的用户
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId) // !!! 这里应该传登录用户的userId, 而不是要被通知的用户的id
	}
}

// 真正打包消息通知好友
func (this *UserProcess) NotifyMeOnline(userId int) {
	// 1. 整理要返回的消息
	var mes message.Message
	var notifyUserStatusMes message.NotifyUserStatusMes

	mes.Type = message.NotifyUserStatusMesType
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 2. 将notifyUserStatusMes序列化赋值给mes
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err = ", err)
		return
	}
	mes.Data = string(data)

	// 3. 将整理好的消息推出去
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("tf.WritePkg(&mes) err = ", err)
		return
	}
}
