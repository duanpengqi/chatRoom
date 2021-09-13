package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// 创建一个全局连接池pool变量
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive, // 和数据库的最大连接数量，0表示没有限制
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
