package processdata

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

// 收到群消息， 并展示
func outputGroupMes(data string) {
	// 1. 将消息反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes -> json.Unmarshal([]byte(data), &smsMes) err = ", err)
		return
	}
	// 2. 展示
	info := fmt.Sprintf("%d说：%s\n", smsMes.UserId, smsMes.Content)
	fmt.Println()
	fmt.Println(info)
	fmt.Println()
}
