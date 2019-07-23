package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//func main() {
//	x := 1
//	println(x)        // 1
//	{
//		println(x)    // 1
//		x := 2
//		println(x)    // 2    // 新的 x 变量的作用域只在代码块内部
//	}
//	println(x)        // 1
//}

// map 错误示例
//func main() {
//	//var m map[string]int
//	m := make(map[string]int)// map 的正确声明，分配了实际的内存
//	m["one"] = 1        // error: panic: assignment to entry in nil map
//	var s []int
//	s = append(s, 1)
//	print(cap(s))
//
//	x := [3]int{1,2,3}
//	func(arr [3]int) {
//		arr[0] = 7
//		fmt.Println(arr)    // [7 2 3]
//	}(x)
//	fmt.Println(x)            // [1 2 3]    // 并不是你以为的 [7 2 3]
//}

//// slice 正确示例
//func main() {
//	var s []int
//	s = append(s, 1)
//}

//
//func main() {
//	var wg sync.WaitGroup
//	done := make(chan struct{})
//	ch := make(chan interface{})
//
//	workerCount := 2
//	for i := 0; i < workerCount; i++ {
//		wg.Add(1)
//		go doIt(i, ch, done, &wg)    // wg 传指针，doIt() 内部会改变 wg 的值
//	}
//
//	for i := 0; i < workerCount+1; i++ {    // 向 ch 中发送数据，关闭 goroutine
//		ch <- i
//	}
//	//关闭done的channel，导致协程退出
//	close(done)
//	wg.Wait()
//	close(ch)
//	fmt.Println("all done!")
//}
//
//func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
//	fmt.Printf("[%v] is running\n", workerID)
//	defer wg.Done()
//	for {
//		select {
//		case m := <-ch:
//			fmt.Printf("[%v] m => %v\n", workerID, m)
//		case <-done:
//			fmt.Printf("[%v] is done\n", workerID)
//			return
//		}
//	}
//}

//mike?为什么没有2 send result
//func main() {
//	ch := make(chan int)
//	done := make(chan struct{})
//
//	for i := 0; i < 3; i++ {
//		fmt.Println(i)
//		go func(idx int) {
//			select {
//			case ch <- (idx + 1) * 2:
//				fmt.Println(idx, "Send result")
//			case <-done:
//				fmt.Println(idx, "Exiting")
//
//			}
//		}(i)
//	}
//
//	fmt.Println("Result: ", <-ch)
//	close(done)
//	time.Sleep(3 * time.Second)
//}
//
//func main() {
//	defer func() {
//		fmt.Println("recovered: ", recover())
//	}()
//	panic("not good")
//}

//func main()  {
//	var v int
//	if v = 2; v > 1 {
//		fmt.Println(v)
//	}
//	fmt.Println(v)
//}

func main() {
	var err error
	js_code := "023Xsn8J063fXg2OQV8J0T1J8J0Xsn8j"
	appid := "wx259f7a8178103266"
	secret := "616c98b85b398d8ab4704f2b962c781d"
	requestString := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, js_code)
	fmt.Println(requestString)
	resp, err := http.Get(requestString)
	if err != nil {

	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}
