package main

import(
	"fmt"
)

func login(userId int, userPwd string) (err error) {
	fmt.Printf("用户Id=%v 密码=%v\n", userId, userPwd)
	return nil
}