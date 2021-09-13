package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的UserDao创建之后， 就可以在很多地方使用
var MyUserDao *UserDao

// 定义一个UserDao结构体， 用来操作User结构体
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式， 创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{
		pool: pool,
	}
}

// 用户登录时， 根据用户的userId在redis-users中查询
// 返回一个User实例+err
func (this *UserDao) GetUserById(conn redis.Conn, userId int) (user *User, err error) {
	// 1. 查找数据
	conn.Do("AUTH", "123456")
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			// 表示在redis中没有找到数据
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	// 2. 查找到了数据反序列化数据后返回user
	user = &User{} // err =  json: Unmarshal(nil *model.User)
	err = json.Unmarshal([]byte(res), user)
	// fmt.Println("user", user)
	// fmt.Println("&user", &user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), &user) err = ", err)
		return
	}
	return
}

// // 用户登录时， 根据用户的userId在redis-users中查询
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 1. 根据userId查询redis-users
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.GetUserById(conn, userId)
	if err != nil {
		fmt.Println("this.GetUserById(conn, userId) err = ", err)
		return
	}
	// 2. 代码走到这说明有数据了，然后验证密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
