package main

import (
	"fmt"
	"myproj.lee/try/common/execshell"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	count := 2
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			//str, err := Run(`peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'`)
			str, err := execshell.Run(`peer chaincode query -C mychannel -n mycc1 -c '{"Args":["checkMemberByFullName","{\"custFullName\":\"kjsdf\"}"]}'`)
			if err != nil {
				fmt.Println(str, err)
			}

		}()
	}
	wg.Wait()
}
