package main

import (
	"time"
	"fmt"
	"context"
	"reflect"
)

var key = "key1"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx1,cancel1 := context.WithTimeout(ctx,3*time.Second)
	defer cancel1()
	go dostuff(ctx1)
	select {
	case <-ctx1.Done():
		fmt.Println("ctx1 done")
		time.Sleep(1*time.Second)
	case <-time.After(4*time.Second):
		fmt.Println("4 second over")
	}

	valueCtx := context.WithValue(ctx, key, "add value")
	//go inspect(valueCtx)
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}
func dostuff(ctx context.Context) {
	tick := time.NewTicker(1*time.Second)
	defer tick.Stop()
	for{
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		case <-tick.C:
			fmt.Println("tick")
		}
	}
}

func inspect(ctx context.Context)  {
	value := reflect.ValueOf(ctx)
	rtype := reflect.TypeOf(ctx)
	//fmt.Println(value,rtype)
	var i = 0
	//fmt.Println("value:",ctx.Value(key))
	for ;i<value.NumMethod();i++{
		if rtype.Method(i).Name=="Value"{
			b := value.Method(i).Call([]reflect.Value{reflect.ValueOf(key)})
			fmt.Println("b:",b[0],len(b))
		}
		fmt.Println(rtype.Method(i).Name)
	}

}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//get value
			//fmt.Println(ctx.Err())
			fmt.Println("[in watch]",ctx.Value(key), "is cancel")
			return
		default:
			//get value
			fmt.Println("[in watch]",ctx.Value(key), "int goroutine")
			time.Sleep(2 * time.Second)
		}
	}
}

