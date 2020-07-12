package main_test

import (
	"testing"
	"io/ioutil"
	"strings"
	"os"
)

func Test_main(t *testing.T) {
	fileName := "./test.go"
	sep := "\""
	r, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var res []string
	ss := strings.Split(string(r), sep)
	for _, v := range ss {
		if strings.Contains(v, ",") {
			vv := strings.Split(v, ",")
			vv[0] = ToCamelCase(vv[0])
			res = append(res, strings.Join(vv, ","))
		} else {
			res = append(res, ToCamelCase(v))
		}
	}

	err = ioutil.WriteFile(fileName, []byte(strings.Join(res, sep)), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func ToCamelCase(str string) string {
	temp := strings.Split(str, "_")
	for i, r := range temp {
		if i > 0 {
			temp[i] = strings.Title(r)
		}
	}
	return strings.Join(temp, "")
}
