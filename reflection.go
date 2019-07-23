package main

import (
	"fmt"
)

type Address struct {
	City string
	Area string
}

type Student struct {
	Address
	Name string
	Age  int
}

func (this Student) Say() {
	fmt.Println("hello, i am ", this.Name, "and i am ", this.Age)
}

func (this Student) Hello(word string) {
	fmt.Println("hello", word, ". i am ", this.Name)
}

//func main(){
//	stu := &Student{Address:Address{City:"Shanghai", Area:"Pudong"}, Name:"chain", Age:23}
//	//StructInfo(stu)
//	//Annoy(stu)
//
//	val := reflect.ValueOf(stu)//&{{Shanghai Pudong} chain 23}
//	fmt.Println(val)
//	typ := reflect.TypeOf(stu)//*main.Student
//	fmt.Println(typ)
//	fmt.Println(typ.Kind())//ptr
//
//	if val.Kind() == reflect.Ptr {
//		fmt.Printf("It is a pointer. Address its value.\n")
//		val = val.Elem()//{{Shanghai Pudong} chain 23}
//		typ = typ.Elem()//main.Student
//	}
//	fmt.Println(val)
//	fmt.Println(typ)
//	fmt.Println(val.NumField())
//
//	fmt.Println("----------changevalue------------")
//	//修改struct中的值
//	for i := 0; i < val.NumField(); i = i + 1 {//3
//		fmt.Println(val.Field(i))//{Shanghai Pudong}
//		fmt.Println(typ.Field(i))//{Address  main.Address  0 [0] true}
//		fmt.Println(val.Field(i).Kind())//struct
//		if val.Field(i).Kind() == reflect.String{
//			val.Field(i).SetString("change")
//		}
//	}
//	fmt.Println(val)//{{Shanghai Pudong} change 23}
//	if newval,ok := val.Interface().(Student); ok{
//		fmt.Printf("The student is %v.\n", newval)//The student is {{Shanghai Pudong} change 23}.
//	}else{
//		fmt.Printf("wrong")
//	}
//	fmt.Println("----------methodcall------------")
//	//fmt.Println(val.NumMethod())
//	time.Sleep(2*time.Second)
//	for i := 0; i < val.NumMethod(); i = i + 1 {//2
//		//fmt.Println(val.Method(i))//0x1099f60
//		//fmt.Println(typ.Method(i))//{Say  func(main.Student) <func(main.Student) Value> 1}
//		//fmt.Println(typ.Method(i).Name)//func
//		if typ.Method(i).Name == "Hello"{
//			val.Method(i).Call([]reflect.Value{reflect.ValueOf("lee")})
//		}else{
//			val.Method(i).Call([]reflect.Value{})
//		}
//	}
//	//有参函数直接调用
//	m2 := val.MethodByName("Hello")
//	m2.Call([]reflect.Value{reflect.ValueOf("iris")})
//}

//
//type Foo struct {
//	A int `tag1:"Tag1" tag2:"Second Tag"`
//	B string
//}
//func main(){
//	// Struct
//	f := Foo{A: 10, B: "Salutations"}
//	// Struct类型的指针
//	fPtr := &f
//	// Map
//	m := map[string]int{"A": 1 , "B":2}
//	// channel
//	ch := make(chan int)
//	// slice
//	sl:= []int{1,32,34}
//	//string
//	str := "string var"
//	// string 指针
//	strPtr := &str
//
//	tMap := examiner(reflect.TypeOf(f), 0)
//	tMapPtr := examiner(reflect.TypeOf(fPtr), 0)
//	tMapM := examiner(reflect.TypeOf(m), 0)
//	tMapCh := examiner(reflect.TypeOf(ch), 0)
//	tMapSl := examiner(reflect.TypeOf(sl), 0)
//	tMapStr := examiner(reflect.TypeOf(str), 0)
//	tMapStrPtr := examiner(reflect.TypeOf(strPtr), 0)
//
//	fmt.Println("tMap :", tMap)
//	fmt.Println("tMapPtr: ",tMapPtr)
//	fmt.Println("tMapM: ",tMapM)
//	fmt.Println("tMapCh: ",tMapCh)
//	fmt.Println("tMapSl: ",tMapSl)
//	fmt.Println("tMapStr: ",tMapStr)
//	fmt.Println("tMapStrPtr: ",tMapStrPtr)
//}
//
//// 类型以及元素的类型判断
//func examiner(t reflect.Type, depth int) map[int]map[string]string{
//	outType := make(map[int]map[string]string)
//
//	// 如果是一下类型，重新验证
//	switch t.Kind() {
//	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
//		fmt.Println("这几种类型Name是空字符串：",t.Name(), ", Kind是：", t.Kind())
//		// 递归查询元素类型
//		tMap := examiner(t.Elem(), depth)
//		for k, v := range tMap{
//			outType[k] = v
//		}
//
//	case reflect.Struct:
//		for i := 0; i < t.NumField(); i++ {
//			f := t.Field(i) // reflect字段
//			outType[i] = map[string]string{
//				"Name":f.Name,
//				"Kind":f.Type.String(),
//			}
//		}
//	default:
//		// 直接验证类型
//		outType = map[int] map[string]string{depth:{"Name":t.Name(), "Kind":t.Kind().String()}}
//	}
//
//	return outType
//}
//
//

//type User struct {
//	Name string `tag:"a";index:"what"`
//	Age  int
//}
//
//func (u *User)ChangeName(s string)  {
//	u.Name = s
//	fmt.Printf("changename to %s\n",s)
//}
//
//func main()  {
//	u := &User{"小米",10}
//	val := reflect.ValueOf(u)
//	typ := reflect.TypeOf(u)
//	fmt.Println(val,typ)
//	val1 := val
//	typ1 := typ
//	if val.Kind() == reflect.Ptr{
//		val = val.Elem()
//		typ = typ.Elem()
//	}
//	for i:= 0;i<val.NumField();i++{
//		if val.Field(i).Kind()==reflect.String{
//			val.Field(i).SetString("华为")
//		}
//		if val.Field(i).Kind()==reflect.Int{
//			val.Field(i).SetInt(30)
//		}
//	}
//
//	fmt.Println(val1)
//	for i:=0;i<val1.NumMethod();i++{
//		if typ1.Method(i).Name == "ChangeName"{
//			val1.Method(i).Call([]reflect.Value{reflect.ValueOf("lee")})
//		}
//	}
//	fmt.Println(val)
//	val1.MethodByName("ChangeName").Call([]reflect.Value{reflect.ValueOf("mike")})
//	fmt.Println(val)
//	user,ok := val.Interface().(User)
//	if ok{
//		fmt.Println(user)
//	}else{
//		fmt.Println("error")
//	}
//}
