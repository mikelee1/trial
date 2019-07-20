package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type A struct {
	ID string
	AffiliationProjectName string
	Content string
	CreateTime time.Time

}


func main()  {
	a := "[[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"jPXqHmGFNyFuEU9nn2IYmw==\",\"CreateTime\":\"2018-12-26T03:47:22.827705493Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"0pgC64+Bn9Fg2EGTnJt7yg==\",\"CreateTime\":\"2018-12-26T03:47:28.268634419Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"0pgC64+Bn9Fg2EGTnJt7yg==\",\"CreateTime\":\"2018-12-26T03:49:30.795468725Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"jPXqHmGFNyFuEU9nn2IYmw==\",\"CreateTime\":\"2018-12-26T03:47:17.665839171Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"0pgC64+Bn9Fg2EGTnJt7yg==\",\"CreateTime\":\"2018-12-26T03:48:21.17582641Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"0pgC64+Bn9Fg2EGTnJt7yg==\",\"CreateTime\":\"2018-12-26T03:53:12.306770168Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"/cs/y9o3OMCKB50G9KF1jQ==\",\"CreateTime\":\"2018-12-26T06:19:55.632712062Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"0pgC64+Bn9Fg2EGTnJt7yg==\",\"CreateTime\":\"2018-12-26T03:48:29.500078451Z\"}],[{\"ID\":\"101d6d1702be4aab916c437756263d44\",\"AffiliationProjectName\":\"项目1\",\"Content\":\"/QvrV/uYTRJaaJKmk/sS1w==\",\"CreateTime\":\"2018-12-26T03:53:32.530792972Z\"}]]"
	tmpa := [][]*A{}
	json.Unmarshal([]byte(a),&tmpa)
	fmt.Printf("%#v\n",tmpa)
	fmt.Println(tmpa[0][0].Content)
}
