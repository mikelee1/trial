package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000))
}
func main() {
	fmt.Println(CreateCaptcha())
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Float64()-0.5)
}
