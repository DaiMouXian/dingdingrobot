# 简单钉钉报警

- 支持以下消息
  - markdown
  - text

请先阅读[钉钉文档](https://open.dingtalk.com/document/orgapp/receive-message)

## 简单使用

```go
ding, err := NewDingRobotWithSecret(dingToken, secret)
```
- 传入设置机器人的Token,secret

```go
dingTalkCli.SendTextMessage("Markdown title", "*")
```
- 设置text文本信息，还有需要at的人员,如果at全体人员的话使用*
- 如果想要at单独人员的话，允许的格式如下
  - +1234567890
  - +987654321098765
>需要钉钉绑定好手机号

```go
err := dingTalkCli.SendMarkDownMessage("Markdown title", "### Link test\n --- \n- <font color=#ff0000 size=6>红色文字</font> \n - content2.", "+18111111115")
```
- markdown语法需要自己在text里面进行封装