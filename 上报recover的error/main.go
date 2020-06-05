package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(Test())
}

func Test() (res string, err1 error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
			e, ok := err.(error)
			fmt.Println("e,ok : ", e, ok)
			if ok {
				fmt.Println(err)
				res = "aaaa"
				err1 = e
			}
		}
	}()
	panic(errors.New("sldkfj"))
	return
}
