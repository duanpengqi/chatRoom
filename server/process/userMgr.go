package processdata

import (
	"fmt"
)

// 因为每当有用户登录时都会用到UserMgr实例中的方法
// 所以创建一个全局的UserMgr实例比较方便
var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess // map中数据布局：100: userProcess
}

// 使用map前需要先make开辟空间, 初始化userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 往onLineUsers中增加userProcess (在map中如果key相同，可以用来修改值)
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除onLineUsers中的某一条数据
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 查询所有数据
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id返回
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 从map中取出一个值带检测方式
	up, ok := this.onlineUsers[userId]
	fmt.Println("看看ok枪管里卖的什么药 = ", ok) // ok = ok 时取值成功， 否则不成功
	if !ok {
		err = fmt.Errorf("用户%d不存在！", userId)
		return
	}
	return
}
