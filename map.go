package main

import (
	"fmt"
	"sync"
	"encoding/json"
)

var sended sync.Map
func main()  {
	for _,v := range []string{"a","b"}{
		sended.Store(v,true)
	}

	for _,v := range []string{"a","b"}{
		fmt.Println(sended.Load(v))
	}

	sended.Range(func(key, value interface{}) bool {
		for _,v := range []string{"a","b"}{
			sended.Store(v,false)
		}
		return true
	})

	for _,v := range []string{"a","b"}{

		a,_ := sended.Load(v)
		fmt.Println(a.(bool))
	}

	test11()
}


func test11()  {
	var testmap = make(map[string]bool)
	for _,v := range []string{"a","b"}{
		sended.Store(v,true)
	}

	for _,v := range []string{"a","b"}{
		fmt.Println(sended.Load(v))
	}

	sended.Range(func(key, value interface{}) bool {
		for _,v := range []string{"a","b"}{
			sended.Store(v,false)
		}
		return true
	})

	for _,v := range []string{"a","b"}{

		a,_ := sended.Load(v)
		fmt.Println(a.(bool))
	}
	sended.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		key1,_ := key.(string)
		fmt.Println(value)
		value1,_ := value.(bool)
		testmap[key1] = value1
		return true
	})
	fmt.Println(testmap)
	a,err:=json.Marshal(testmap)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%v",string(a))
}