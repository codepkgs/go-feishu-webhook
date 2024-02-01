# 说明

> go版本的飞书自定义Webhook机器人sdk。

# 功能列表

## 支持的消息类型

[自定义机器人接入](https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot#756b882f)

* [X]  普通文本消息 `client.Text`
* [ ]  Markdown消息 `client.Markdown`
* [ ]  群名片消息 `client.ShareChat`
* [ ]  图片消息 `client.Image`
* [ ]  消息卡片 `client.Interactive`

# 示例

* 初始化Client

  > WebhookAddress 为创建机器人时产生的Webhook地址。
  > 如果创建的机器人的安全设置采用的是 自定义关键词 或 IP白名单，在创建client的时候，`Secret` 传入空字符串即可。
  > 如果创建的机器人的安全设置采用的是 加签，在创建client的时候，`Secret` 传入产生的密钥即可。
  >

  ```go
  // 初始化一个未采用加签的机器人
  client, err := feishu.NewClient(
      "https://open.feishu.cn/open-apis/bot/v2/hook/xxxx-xxxx-xxxx",
      "")
  if err != nil {
      fmt.Println(err)
  }
  ```
  
  ```go
  // 初始化一个采用加签的机器人
  client, err := feishu.NewClient(
      "https://open.feishu.cn/open-apis/bot/v2/hook/xxxx-xxxx-xxxx",
      "xxxxxx")
  if err != nil {
      fmt.Println(err)
  }
  ```
* 发送文本消息

  > 不支持at单个用户，只支持at所有人
  >

  ```go
  sr, err := client.Text("测试消息", false)
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", sr)
  }