package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type Node struct {
	Length int
	Data   []int
}

func main() {
	var nodes []*Node
	scanner := bufio.NewScanner(os.Stdin)
	remain := 0
	tmpNode := &Node{}
	for scanner.Scan() {
		numStr := scanner.Text()

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return
		}

		if remain == 0 {
			remain = num
			tmpNode = &Node{
				Length: num,
			}
			continue
		}
		remain--
		tmpNode.Insert(num)

		if remain == 0 {
			nodes = append(nodes, tmpNode)
		}
	}

	for _, node := range nodes {
		for _, d := range node.Data {
			fmt.Println(d)
		}
	}
}

func (n *Node) Insert(num int) {
	added := false
	for k, v := range n.Data {
		if num == v {
			return
		}
		if num < v {
			added = true
			n.Data = append(n.Data[:k], append([]int{num}, n.Data[k:]...)...)
			break
		}
	}
	if !added {
		n.Data = append(n.Data, num)
	}

}
