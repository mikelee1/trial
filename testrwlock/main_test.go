package main_test

import (
	"testing"
	"sync"
	"time"
	"fmt"
	"math/rand"
)

var updated = time.Time{}
var locker = sync.Mutex{}

func init() {
	updated = time.Now()
}

//每隔三秒更新一次
func Test_main(t *testing.T) {

	for range [10]int{} {
		go func() {
			for {
				onetime()
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}
	time.Sleep(1 * time.Minute)
}

func onetime() {
	if time.Now().Sub(updated).Seconds() > 3 {
		locker.Lock()
		defer locker.Unlock()
		if time.Now().Sub(updated).Seconds() > 3 {
			updated = time.Now()
			fmt.Println(updated.String())
		}

		time.Sleep(1 * time.Second)
		return
	}
	//fmt.Println(updated.String())
}

//并发检查node状态
var statusLock sync.RWMutex

type Docker struct {
	nodeStatusChanged bool
}

var globalNodes = []string{"running", "running"}

func Test_11(t *testing.T) {
	dr := Docker{}
	ticker := time.NewTicker(time.Duration(2) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		go func() {

			statusLock.RLock()
			nodes := globalNodes
			statusLock.RUnlock()

			if dr.nodeStatusChanged {
				statusLock.Lock()
				dr.nodeStatusChanged = false
				statusLock.Unlock()
			}

			for k, node := range nodes {
				curStatus := getRandom()
				func() {
					if node != curStatus {
						fmt.Println("not same")
						if !dr.nodeStatusChanged {
							fmt.Println("not nodeStatusChanged")
							statusLock.Lock()
							dr.nodeStatusChanged = true
							time.Sleep(time.Second)
							globalNodes[k] = curStatus

							statusLock.Unlock()
						}

					}
				}()

			}
			fmt.Println(globalNodes)

		}()
	}

	select {}
}

func getRandom() string {
	s := rand.NewSource(time.Now().UnixNano())
	i := rand.New(s).Intn(10)
	if i > 7 {
		return "stop"
	}
	return "running"
}
