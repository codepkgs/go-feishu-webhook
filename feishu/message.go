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
