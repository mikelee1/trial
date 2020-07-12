package main

import (
	"strings"
	"fmt"
)

//varchar(.*\)) 替换为string

func main() {
	case1 := `
	contract_no           string  //	信托合同编号
	prod_code             string  //	产品代码
	contract_name         string  //	信托合同名称
	consignor             string  //	委托人
	trustee               string  //	受托人
	beneficiaries         string  //	受益人
	principal_beneficiary float64 //	信托受益权本金（元）
	build_date            string  //	信托设立日
	distribution_date     string  //	信托利益分配日
	expected_rate         float64 //	信托预期收益率（%）
	expected_end_date     string  //	信托预期到期日
	expiry_date           string  //	信托法定到期日
	social_credit_code    string  //	统一社会信用代码
	trustee_contact       string  //	受托机构联系人
	trustee_tel           string  //	受托机构电话
	sync_type             byte    //	同步状态 0-未同步 1-同步失败 2-同步成功 9-其他
	create_time           string
	`
	fmt.Println(Trans(case1))

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
