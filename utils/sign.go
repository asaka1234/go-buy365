package utils

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
)

// 计算请求签名
func SignDeposit(paramMap map[string]interface{}, accessKey string) string {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串
	rawString := ""
	lo.ForEach(keys, func(x string, index int) {
		value := cast.ToString(paramMap[x])
		rawString += fmt.Sprintf("%s=%s", x, URLDecodeSafe(value))

		if index != len(keys)-1 {
			//不是最后一个,则拼接
			rawString += "&"
		}
	})

	// 3. 将secretKey拼接到最后
	rawString += accessKey

	//4. 计算md5签名
	signResult := GetMD5([]byte(rawString))
	return signResult
}

// 验证签名
func VerifySignDeposit(params map[string]interface{}, signKey string) bool {
	// Check if sign exists in params
	signValue, exists := params["sign"]
	if !exists {
		return false
	}

	// Get the sign value and remove it from params
	key := fmt.Sprintf("%v", signValue)
	delete(params, "sign")

	// Generate current signature
	currentKey := SignDeposit(params, signKey)

	// Compare the signatures
	return key == currentKey
}

// 计算withdraw请求签名
func SignWithdraw(paramMap map[string]interface{}, accessKey string) string {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	//拼接签名原始字符串（value直接拼）
	rawString := ""
	lo.ForEach(keys, func(x string, index int) {
		value := cast.ToString(paramMap[x])
		rawString += value
	})

	// 3. 将secretKey拼接到最后
	rawString += accessKey

	//4. 计算md5签名
	signResult := GetMD5([]byte(rawString))
	return signResult
}
