package weixinpay

import (
	"crypto/md5"
	"fmt"
	"sort"
)

// 根据反馈回来的信息，生成签名结果
func getVerifySign(preStr string, key string) string {
	preSignStr := preStr + "&key=" + key
	return fmt.Sprintf("%X", md5.Sum([]byte(preSignStr)))
}

// 把数组所有元素排序，并按照“参数=参数值”的模式用“&”字符拼接成字符串
func createLinkString(params map[string]string) string {
	//过滤sign参数
	keys := paraFilter(params)
	sort.Strings(keys)

	preStr := ""
	for _, key := range keys {
		value := params[key]
		//过滤空值
		if len(value) == 0 {
			continue
		}

		preStr = preStr + key + "=" + value + "&"
	}
	if len(preStr) > 0 {
		preStr = preStr[0:(len(preStr) - 1)]
	}
	return preStr
}

// 除去数组中的签名参数
func paraFilter(sArray map[string]string) []string {
	keys := make([]string, 0, len(sArray))

	if sArray == nil || len(sArray) == 0 {
		return keys
	}

	for k := range sArray {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}

	return keys
}
