package model

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

// 1. dao = data access object
// 2. 编写对User对象（实例）操作的各种方法 。 主要就是增删改查

// 定义一个UserDao的结构体， 完成对User的增删改查操作
type UserDao struct {
	pool *redis.pool
}

// 
