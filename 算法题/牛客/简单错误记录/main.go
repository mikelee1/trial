package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

type Res struct {
	Name  string
	Line  string
	Count int
}

type Ress struct {
	Ress  []*Res
	Store map[string]int
}

func main() {
	ress := Ress{
		Store: map[string]int{},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		list := strings.Split(text, " ")
		if len(list)%2 != 0 {
			continue
		}

		for i := 0; i < len(list)/2; i++ {

			path := strings.Split(list[i*2], "\\")
			filename := path[len(path)-1]
			key := filename + list[i*2+1]
			//fmt.Println(key, ress.Store[key])
			//fmt.Println(list)
			if k, exist := ress.Store[key]; exist && k > -1 {
				//fmt.Println(len(ress.Ress), k)
				ress.Ress[k].Count++
				continue
			}

			bound := len(filename)
			if bound > 16 {
				bound = 16
			}
			//fmt.Println(list[1])
			res := &Res{
				Name:  filename[len(filename)-bound:],
				Count: 1,
				Line:  list[i*2+1],
			}
			//fmt.Println(res)
			if len(ress.Ress) == 8 {
				for k, _ := range ress.Store {
					ress.Store[k]--
				}
			}
			ress.Store[key] = len(ress.Ress)
			if len(ress.Ress) == 8 {
				ress.Ress = append(ress.Ress[1:], res)
				ress.Store[key] = 7
			} else {
				ress.Ress = append(ress.Ress, res)
			}
		}

	}

	//for k, v := range ress.Ress {
	//	fmt.Println(k, v)
	//}
	for _, v := range ress.Ress {
		fmt.Println(v.Name, v.Line, v.Count)
	}
}
