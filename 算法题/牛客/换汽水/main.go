package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var res []int
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil || num == 0 || count > 10 {
			break
		}
		count++
		res = append(res, MaxBottle(num))
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func MaxBottle(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	remain := n % 3
	newBottle := (n - remain) / 3

	return newBottle + MaxBottle(remain+newBottle)
}
