package mp

import (
	"crypto/sha1"
	"fmt"
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
	CreateTime   time.Duration
	MsgType      string
	Content      string
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
