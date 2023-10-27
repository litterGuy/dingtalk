# 钉钉自定义机器人Webhook

[官方文档](https://open.dingtalk.com/document/robots/custom-robot-access#title-jfe-yo9-jl2)

方便发送机器人

## 使用

1. 引入包

```go
go get github.com/litterGuy/dingtalk
```

2. 调用

```go
client := NewClient("token", "secret")
text := NewTextMessage()
text.SetContent("测试机器人使用")
req, res, err := client.Send(text)
if err != nil {
panic(err)
}
println(req)
println(fmt.Sprintf("%v:%s", res.ErrCode, res.ErrMsg))
```