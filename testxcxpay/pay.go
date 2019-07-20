package main

import (
	"github.com/medivhzhan/weapp/payment"
	"fmt"
	"net/http"
)

func main() {
	var err error
	defer func() {
		fmt.Println(err)
	}()

	// 新建支付订单
	form := payment.Order{
		// 必填
		AppID:      "APPID",
		MchID:      "aa",
		Body:       "商品描述",
		NotifyURL:  "通知地址",
		OpenID:     "通知用户的 openid",
		OutTradeNo: "商户订单号",
		TotalFee:   10,

		// 选填 ...
		NoCredit:  false,
		Detail:    "商品详情",
	}

	res, err := form.Unify("支付密钥")
	if err != nil {
		// handle error
		return
	}

	fmt.Printf("返回结果: %#v", res)

	// 获取小程序前点调用支付接口所需参数
	params, err := payment.GetParams(res.AppID, "微信支付密钥", res.NonceStr, res.PrePayID)
	fmt.Println(params)
	if err != nil {
		// handle error
		return
	}
}

func nofity(w http.ResponseWriter, req *http.Request) error {
	err := payment.HandlePaidNotify(w , req ,  func(ntf payment.PaidNotify) (bool, string) {
		// 处理通知
		fmt.Printf("%#v", ntf)

		// 处理成功 return true, ""
		// or
		// 处理失败 return false, "失败原因..."
		return false, ""
	})
	return err
}