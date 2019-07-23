package main

import (
	"fmt"
	"regexp"
)

func main() {
	var Password string
	regp := `^[a-zA-Z0-9]{3,6}$`
	rgxp := regexp.MustCompile(regp)
	for {
		fmt.Println("input:")
		fmt.Scanln(&Password)
		fmt.Println(Password)
		passwordverified := rgxp.MatchString(Password)
		if !passwordverified {
			fmt.Println("no pass")
			return
		}
		fmt.Println("pass")
	}
}
