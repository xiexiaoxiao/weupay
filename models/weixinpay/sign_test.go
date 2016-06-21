package weixinpay

import (
	"testing"
)

func TestSign(t *testing.T) {
	params := make(map[string]string)
	params["appid"] = "111"
	params["mch_id"] = "222"
	preStr := createLinkString(params)
	t.Log(preStr)
	sign := getVerifySign(preStr, "123")
	t.Log(sign)
}
