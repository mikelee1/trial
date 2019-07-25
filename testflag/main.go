package main

import (
	"flag"
	"fmt"
)

var (
	adder int
	beadder int
)

func init()  {
	flag.IntVar(&adder,"adder",1,"input the adder")
	flag.IntVar(&beadder,"beadder",2,"input the beadder")
}

func main() {
	flag.Parse()
	fmt.Println(adder,beadder)

}
