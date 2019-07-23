package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a := map[string]interface{}{
		"appid":     "wx2f19daf67101a3a2",
		"fee_type":  "CNY",
		"sign_type": "MD5",
		"total_fee": 10,
		//"sign": "BB39F3CF35392255D56EE21AB31E2961",
		"out_trade_no": "wx1559141108300384492", //一个
		"trade_type":   "JSAPI",
		"body":         "预定早餐",
		"device_info":  "WEB",
		"mch_id":       "1537102541",
		"nonce_str":    "20190529224508�", //一个
		"notify_url":   "https://yidaizi.liyuanye.club/order/wxpay/notify",
		"openid":       "oRWOm5M02GuQVTJ_nxdtUJunbC28",
	}
	b := WxPayCalcSign(a, "aaa")
	fmt.Println(b)

	d := map[string]string{
		"appid":     "wx2f19daf67101a3a2",
		"fee_type":  "CNY",
		"sign_type": "MD5",
		"total_fee": strconv.Itoa(10),
		//"sign": "BB39F3CF35392255D56EE21AB31E2961",
		"out_trade_no": "wx1559141108300384492", //一个
		"trade_type":   "JSAPI",
		"body":         "预定早餐",
		"device_info":  "WEB",
		"mch_id":       "1537102541",
		"nonce_str":    "20190529224508�", //一个
		"notify_url":   "https://yidaizi.liyuanye.club/order/wxpay/notify",
		"openid":       "oRWOm5M02GuQVTJ_nxdtUJunbC28",
	}
	c := WxPayCalcSignString(d, "aaa")
	fmt.Println(c)
}

//微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//fmt.Println("加密前-----", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	//fmt.Println("加密后-----", upperSign)
	return upperSign
}

//微信支付计算签名的函数
func WxPayCalcSignString(mReq map[string]string, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//fmt.Println("加密前-----", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	//fmt.Println("加密后-----", upperSign)
	return upperSign
}
