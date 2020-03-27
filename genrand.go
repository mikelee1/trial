package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10))
}
func main() {
	fmt.Println(CreateCaptcha())
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Float64()-0.5)
}
