package dingtalk

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	client := NewClient("token", "secret")
	text := NewTextMessage()
	text.SetContent("测试机器人使用")
	req, res, err := client.Send(text)
	if err != nil {
		panic(err)
	}
	println(req)
	println(fmt.Sprintf("%v:%s", res.ErrCode, res.ErrMsg))
}
