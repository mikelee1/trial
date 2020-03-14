package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		f, err := strconv.ParseFloat(text, 32)
		if err != nil {
			break
		}
		value := f - math.Floor(f)
		if value >= 0.5 {
			fmt.Println(math.Ceil(f))
		} else {
			fmt.Println(math.Floor(f))
		}
	}
}
