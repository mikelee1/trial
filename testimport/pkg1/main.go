package pkg1

import "fmt"

func init() {
	fmt.Println("pkg1")
	go func() {
		fmt.Println("pkg1 go")
	}()
}
