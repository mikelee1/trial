package main

import (
	"strings"
	"fmt"
	"regexp"
)

//varchar(.*\)) 替换为string

func main() {
	case1 := `
	tx_code                     string  //	交易代码
	prod_short_name             string  //	产品简称
	prod_code                   string  //	产品代码
	transferor_name             string  //	出让方名称
	transferor_acct_name        string  //	出让方资产账号
	transferee_name             string  //	受让方名称
	transferee_acct_name        string  //	受让方资产账号
	tx_scale                    float64 //	交易规模
	turnover                    float64 //	成交金额
	loan_mngm_fee_rate          float64 //	贷款管理费率
	loan_mngm_fee_pay_freq      string  //	贷款管理费支付频率
	accrued_interest            float64 //	应计利息
	accrued_interest_pay_method string  //	应计利息支付方式
	argmt_rate                  float64 //	安排费率
	fixture_date                string  //	成交日期
	settlement_method           string  //	结算方式
	settlement_date             string  //	结算日
	description                 string  //	备注
	sync_type                   byte    //	同步状态 0-未同步 1-同步失败 2-同步成功 9-其他
	create_time                 string  //  创建时间
	`
	//fmt.Println(Trans(case1))

	fmt.Println(ToSnakeCase(case1))
}

func Trans(snack string) string {
	aaa := strings.Split(snack, " ")
	snackWords := []string{}
	for _, sw := range aaa {
		snackWords = append(snackWords, strings.Split(sw, "	")...)
	}

	var res []string
	for _, word := range snackWords {
		if strings.Contains(word, "_") {
			wordInSnack := strings.Split(word, "_")
			word := ""
			for _, v1 := range wordInSnack {
				b := ""
				for k2, v2 := range strings.Split(v1, "") {
					if k2 == 0 {
						b += strings.ToUpper(v2)
					} else {
						b += v2
					}
				}
				word += b
			}
			res = append(res, word)
		} else {
			res = append(res, word)
		}
	}
	return strings.Join(res, " ")
}

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	fmt.Println(snake)
	return strings.ToLower(snake)
}
