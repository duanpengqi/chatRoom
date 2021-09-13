package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

// 消息传输的结构体 包括消息类型 和 消息体
type Message struct {
	Type string `json: "type"` //消息类型
	Data string `json: "data"` //消息内容
}

// 登录信息的结构体
type LoginMes struct {
	UserId   int    `json: "userId"`   // 用户Id
	UserPwd  string `json: "userPwd"`  // 用户密码
	UserName string `json: "userName"` // 用户名
}

// 服务器对登录信息返回的结构体
type LoginResMes struct {
	Code  int    `json: "code"`  // 返回登录的状态码 200 表示登录成功
	Error string `json: "error"` // 返回的错误信息
}

// 注册信息的结构体
type RegisterMes struct {
	User User `json: "user"`
}

// 注册之后，服务器返回的结构体
type RegisterResMes struct {
	Code  int    `json: "code"`  // 返回登录的状态码 200 表示注册成功
	Error string `json: "error"` // 返回的错误信息
}
