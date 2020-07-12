package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	file, err := ioutil.ReadFile("./unmarshalnginx/nginx.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
	n := Nginx{

	}
	err = json.Unmarshal(file, &n)
	if err != nil {
		panic(err)
	}
	fmt.Println(n.Events)

}

type Nginx struct {
	Events string
}
