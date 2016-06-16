package weixinmp

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"time"
)

const (
	TOKEN = "lemmihtoken"
	TEXT  = "text"
)

type BaseMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

type Request struct {
	XMLName xml.Name `xml:"xml"`
	MsgId   int
	BaseMsg
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	BaseMsg
}

func Signature(timestamp, nonce string) string {
	preStr := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(preStr)
	baseStr := ""
	for _, s := range preStr {
		baseStr += s
	}
	h := sha1.New()
	h.Write([]byte(baseStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Reply(req *Request) (res *Response, err error) {
	res = &Response{}
	res.CreateTime = time.Now().Unix()
	res.ToUserName = req.FromUserName
	res.FromUserName = req.ToUserName
	res.MsgType = TEXT

	beego.Info("req MsgType:" + req.MsgType)
	beego.Info("req Content:" + req.Content)
	if req.MsgType != TEXT {
		res.Content = "暂时不支持的消息类型"
		return
	}

	res.Content = "您说的是:" + req.Content
	return
}
