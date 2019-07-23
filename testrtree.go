package main

import (
	"fmt"

	"github.com/charignon/rtree/rtree"
)

func main() {
	r := rtree.NewRegularRTree().WithCapacity(3)
	r.Insert(rtree.Rect{1, 2, 1, 1}, "A")
	r.Insert(rtree.Rect{1, 5, 1, 1}, "B")
	r.Insert(rtree.Rect{5, 10, 1, 1}, "C")
	r.Insert(rtree.Rect{500, 560, 1, 1}, "D")
	r.Insert(rtree.Rect{-200, -100, 1, 1}, "E")
	fmt.Println(r.Search(rtree.Rect{0, 2, 0, 20}))
}
