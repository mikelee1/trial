package main

import (
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string `yaml:"aa"`
	//B struct {
	//	RenamedC int   `yaml:"c"`
	//	D        []int `yaml:",flow"`
	//}
	B map[string]Info
	C []string
	CrossDomain bool `yaml:"CrossDomain"`
	Routers []FromTo
}

type FromTo struct {
	From string
	To string
}


type Info struct {
	Name string
}

func main() {
	t := T{}
	yamlFile, err := ioutil.ReadFile("testyaml/app.conf")
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(t.CrossDomain)
	fmt.Println(t.Routers)


	//a, err := client.ParseHostURL(t.A)
	//fmt.Println(a, err)

	//
	//m := make(map[interface{}]interface{})
	//
	//err = yaml.Unmarshal([]byte(data), &m)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m:\n%v\n\n", m)
	//
	//d, err = yaml.Marshal(&m)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
