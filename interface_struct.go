package main

import (
	"fmt"
	"strconv"
	"reflect"
	"encoding/json"
	"bytes"
	//"myproj/try/goexample"
)


type Man struct {
	Name string
}
type Street struct {
	Name string
}

func (m *Man)Eat()  {
	fmt.Println("i am eating")
}
func (m *Man)Walk(st Street) (string,error) {
	fmt.Printf("%s is walking on %s\n",m.Name,st.Name)
	return "2017",fmt.Errorf("no\n")
}


type Maner interface {
	Eat()
	Walk(st Street) (string,error)
}

func checkinterface()  {
	var man Maner
	//lee := Man{"lee"}
	//man = &lee
	var lee = new(Man)
	man = lee

	man.Eat()

	street := Street{"cangqian"}

	time,error := man.Walk(street)
	fmt.Printf(time)
	fmt.Printf("%v\n",error)
}
//type
func checktype(x interface{})  {

	defer func() {
		if err:=recover();err!=nil{
			fmt.Printf("caught a panic: %v\n",err)
		}
	}()

	switch x.(type) {
	case string:
		fmt.Printf("is string\n")
	case []byte:
		fmt.Printf("is []byte\n")
	default:
		fmt.Printf("not found\n")
		panic("no type found")
	}
}
//递归
func fibonacci(x int) int {
	if x <= 1{
		return 1
	}
	return fibonacci(x-1)+fibonacci(x-2)
}

func closurefunc(a string) func() string {
	return func() string {
		fmt.Printf("a:%s\n",a)
		return a
	}
}
//闭包函数
func MakeAddSuffix(suffix string) func(cont string) string {
	return func(cont string) string {
		return cont+suffix
	}
}

func deferclosure() (ret int) {

	defer func() {
		ret++
	}()
	return 1
}

func sliceadd(n *[3]int) (res int) {

	for _,v := range n {
		res += v
	}
	return
}

//[]byte转string
func writebuffer(b *bytes.Buffer,bs []byte) string {
	b.WriteByte(99)
	for _,v := range bs{
		b.WriteByte(v)
	}
	return b.String()
}

func printall()  {
	a :=[]byte("111")
	fmt.Printf("%v\n",a)

	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n ", i, c)
	}

	var b []byte
	var s1 string


	s1 = "444"
	b = append(b,s1...)
	fmt.Printf("%v\n",b)
	//make(map[string]int)
	maplist := map[string]int{"a":1}
	fmt.Println("%v\n",maplist)
	var mapfunc = make(map[int]func()int)
	mapfunc = map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
	}
	mapfunc[3]= func() int {
		return 30
	}
	fmt.Printf("%v\n",mapfunc)
}

type F1 struct {
	car string
}

func (f1 *F1)run()  {
	fmt.Printf("\n%v\n",f1.car)
}

type Dazhong struct {
	car string
}

func (daz *Dazhong)run()  {
	fmt.Printf("\n%s\n",daz.car)
}


type Sporter interface {
	run()
}

func stringtobyteslice(s string) []byte {
	a,_ := json.Marshal(s)
	return a
}

func stringslicetobyteslice(ss []string) []byte {
	a,_ := json.Marshal(ss)
	return a
}

func structtobyteslice(s *Man) ([]byte) {
	res,_ :=json.Marshal(s)
	return res
}

func byteslicetostruct(s []byte) *Man {
	var a *Man
	json.Unmarshal(s,&a)
	return a
}

func byteslicetostringslice(s []byte) []string {
	var a []string
	json.Unmarshal(s,&a)
	return a
}

func byteslicetostring(s []byte) string {
	var a string
	json.Unmarshal(s,&a)
	return a
}

func zhi_and_yinyong_yuyi()  {//值语义和引用语义
	var a = [3]int{1,2,3}
	var c = []int{1,2,3}
	var b = a

	b[1] = 3
	fmt.Printf("a is %v,b is %v\n",a,b)
	fmt.Printf("a type:%v,c type:%v\n",reflect.TypeOf(a),reflect.TypeOf(c))
	fmt.Printf("a value:%v,c value:%v\n",reflect.ValueOf(a),reflect.ValueOf(c))

}

func interfacetype(f3 interface{})  {

	switch t:=f3.(type) {
	case *F1:
		fmt.Printf("lalalala%s is f1 type\n",t.car)
	case *Dazhong:
		fmt.Printf("lalalala%s is dazhong type\n",t.car)
	}
}


func main()  {
	checkinterface()
	checktype([][]int{})
	checktype([]byte(strconv.Itoa(1)))
	fmt.Printf("fibonacci result:%v\n",fibonacci(5))

	clofuc := closurefunc("A")
	clofuc()
	addBmp := MakeAddSuffix(".bmp")
	fmt.Printf("%v\n",addBmp("pic1"))

	fmt.Printf("%v\n",deferclosure())

	n := [3]int{}
	fmt.Printf("%v\n",sliceadd(&n))

	bufer := new(bytes.Buffer)
	byteslice := []byte("what")
 	fmt.Printf("buffer:%v\n",writebuffer(bufer,byteslice))

	printall()
	//goexample.Invokeme()

	//f1 := F1{"ferrari"}
	f1 := new(F1)
	f1.car = "ferrari"
	f1.run()

	var f2 F1
	f2.car = "benci"
	f2.run()


	var f3 Dazhong
	f3.car = "jiakechong"

	var sp Sporter
	sp = &f3
	sp.run()
	interfacetype(sp)
	switch t:=sp.(type) {
	case *F1:
		fmt.Printf("%s is f1 type\n",t.car)
	case *Dazhong:
		fmt.Printf("%s is dazhong type\n",t.car)
	}

	fmt.Printf("type of a:%v\n",reflect.TypeOf(sp))
	fmt.Printf("value of a:%s\n\n",reflect.ValueOf(sp))

	fmt.Printf("string to byte slice %v\n",stringtobyteslice("waht"))
	fmt.Printf("byte slice to string %v\n\n",byteslicetostring( stringtobyteslice("waht") ))

	fmt.Printf("string slice to byte slice %v \n",stringslicetobyteslice([]string{"1","2"}))
	fmt.Printf("byte slice to string slice %v \n\n",byteslicetostringslice(  stringslicetobyteslice([]string{"1","2"})  ))

	fmt.Printf("struct to byte %v \n",structtobyteslice(&Man{"lee"}))
	fmt.Printf("byte to struct %v \n\n",byteslicetostruct(  structtobyteslice(&Man{"lee"})  ))

	zhi_and_yinyong_yuyi()




}