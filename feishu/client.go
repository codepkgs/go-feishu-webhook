package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	WebhookAddress string
	Secret         string
	Timestamp      string
	Sign           string
}

var ErrWebhookAddress = fmt.Errorf(`feishu webhook address must begin with "http://" or "https://"`)

func NewClient(webhookAddress, secret string) (*Client, error) {
	if !strings.HasPrefix(webhookAddress, "http://") && !strings.HasPrefix(webhookAddress, "https://") {
		return nil, ErrWebhookAddress
	}

	return &Client{
		WebhookAddress: webhookAddress,
		Secret:         secret,
	}, nil
}

// sign 计算签名 -- 不管Secret有没有都计算签名
func (c *Client) sign() {
	timestamp := time.Now().UnixMilli() / 1000
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + c.Secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	h.Write(data)
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	c.Timestamp = strconv.FormatInt(timestamp, 10)
	c.Sign = sign
}
