package pkg1

import "fmt"

var AAA  = "1"

func init() {
	fmt.Println("pkg1")
	go func() {
		fmt.Println("pkg1 go")
	}()
}
