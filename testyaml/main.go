package main

import (
	"fmt"
	"log"

	"github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	A string
	//B struct {
	//	RenamedC int   `yaml:"c"`
	//	D        []int `yaml:",flow"`
	//}
	B map[string]Info
}

type Info struct {
	Name string
}

func main() {
	t := T{}
	yamlFile, err := ioutil.ReadFile("testyaml/test.yaml")
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	a, err := client.ParseHostURL(t.A)
	fmt.Println(a, err)

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
