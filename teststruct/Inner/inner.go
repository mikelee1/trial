package Inner

import "fmt"

type B struct {
	Address string
}

func (c B) Print() {
	fmt.Println("B")
}
