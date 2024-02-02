package feishu

import (
	"encoding/json"
)

func getReturn(bytes []byte) (*SendResult, error) {
	var r SendResult
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

type SendResult struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	Code          int    `json:"code"`
	Data          any    `json:"data"`
	Msg           string `json:"msg"`
}

// Text 发送普通文本消息
func (c *Client) Text(content string, isAtAll bool) (*SendResult, error) {
	// 更新签名
	c.sign()

	if isAtAll {
		content = content + `<at user_id="all">所有人</at>`
	}

	t := struct {
		TimeStamp string `json:"timestamp"`
		Sign      string `json:"sign"`
		MsgType   string `json:"msg_type"`
		Content   struct {
			Text string `json:"text"`
		} `json:"content"`
	}{
		TimeStamp: c.Timestamp,
		Sign:      c.Sign,
		MsgType:   "text",
		Content: struct {
			Text string `json:"text"`
		}{Text: content},
	}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(body)
	if err != nil {
		return nil, err
	}

	return getReturn(resp)
}

type RichTextContent struct {
	Tag    string `json:"tag"`
	Text   string `json:"text,omitempty"`
	Href   string `json:"href,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

func RichTextContentWithText(text string) RichTextContent {
	return RichTextContent{
		Tag:  "text",
		Text: text,
	}
}

func RichTextContentWithLink(text string, link string) RichTextContent {
	return RichTextContent{
		Tag:  "a",
		Text: text,
		Href: link,
	}
}

func RichTextContentWithAtAll() RichTextContent {
	return RichTextContent{
		Tag:    "at",
		UserId: "all",
	}
}

// RichText 富文本消息
func (c *Client) RichText(title string, contents [][]RichTextContent) (*SendResult, error) {
	// 更新签名
	c.sign()

	type lang struct {
		Title   string              `json:"title,omitempty"`
		Content [][]RichTextContent `json:"content,omitempty"`
	}

	type richText struct {
		TimeStamp string `json:"timestamp"`
		Sign      string `json:"sign"`
		MsgType   string `json:"msg_type"`
		Content   struct {
			Post struct {
				ZhCn lang `json:"zh_cn,omitempty"`
			} `json:"post"`
		} `json:"content"`
	}

	t := richText{
		MsgType:   "post",
		TimeStamp: c.Timestamp,
		Sign:      c.Sign,
	}

	t.Content.Post.ZhCn.Title = title
	for _, content := range contents {
		t.Content.Post.ZhCn.Content = append(t.Content.Post.ZhCn.Content, content)
	}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(body)
	if err != nil {
		return nil, err
	}

	return getReturn(resp)
}
