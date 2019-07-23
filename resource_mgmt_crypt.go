package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(pwd string) string {
	pwd = "superadmin123"
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		//logger.Panicf("Error generating bcrypt password : %s", err)
	}
	return string(bytes)
}
func main() {
	res := CreatePassword("1")
	fmt.Println(res)
}
