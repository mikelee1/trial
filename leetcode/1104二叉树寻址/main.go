package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pathInZigZagTree(1000))
}

func powToInt(label int) int {
	num := float64(label)
	for i := 0; true; i++ {
		if math.Pow(2, float64(i)) <= num && math.Pow(2, float64(i+1)) > num {
			return i + 1
		}
	}
	return 0
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func pathInZigZagTree(label int) []int {

	closec := make(chan bool)
	datac := make(chan int)
	go func() {
		digui(label, datac)
	}()
	res := []int{}
	go func() {
		for {
			v, ok := <-datac
			if !ok {
				closec <- true
				break
			} else {
				res = append([]int{v}, res...)
			}
		}
	}()
	<-closec
	return res
}

func digui(label int, datac chan int) {
	if label == 1 {
		datac <- 1
		close(datac)
		return
	}
	hang := powToInt(label)
	former := pow(2, hang-2) + (pow(2, hang-2) - 1 - (label-pow(2, hang-1))/2)
	datac <- label
	digui(former, datac)
}
