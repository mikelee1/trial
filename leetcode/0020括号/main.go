package main

import (
	"fmt"
	"strings"
	"sync/atomic"
	"reflect"
)

func main() {
	t := isValid("(")
	fmt.Println(t)
}

var allBracts = []string{"(",")","[","]","{","}"}
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice: {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	}
	return
}

func isValid(s string) bool {

	bs := strings.Split(s,"")
	storelist := []string{}
	bmap := make(map[string]int32,3)
	for _, value := range bs {
		i := Contains(allBracts,value)
		if i == -1{
			return false
		}
		if i % 2 == 1{
			b,ok := bmap[allBracts[i-1]]
			if !ok{
				return false
			}
			i1 := Contains(allBracts,storelist[len(storelist)-1])
			if i1+1 != Contains(allBracts,value){
				return false
			}
			bmap[allBracts[i-1]] = atomic.AddInt32(&b,-1)
			storelist = storelist[:len(storelist)-1]
		}else{
			b := bmap[value]
			bmap[value] = atomic.AddInt32(&b,1)
			storelist = append(storelist,value)
		}
	}
	for _, value := range bmap {
		if value != 0{
			return false
		}
	}
	return true
}