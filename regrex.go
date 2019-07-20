package main

import (
	"fmt"
	"regexp"
	"strings"
	"bytes"
)

func main() {  
    fmt.Println("------FindString------")  
  
    str := "测试"
	match,_:=regexp.MatchString("^[a-z0-9-\u4e00-\u9fa5]+$",str)
	fmt.Println(match)
	fmt.Printf("len is:%s\n",len(str))
	if len(str)>20{
		fmt.Printf("over 20\n")
		errMsg := fmt.Sprintf("[checkMultiPartForm] Chaincode Name should be less than 20\n")
		fmt.Printf(errMsg)
	}


    reg := regexp.MustCompile("[a-zA-Z0-9]") //六位连续的数字  
    fmt.Println(reg.MatchString(str))




	str = " welcome to baidu.com "

	//// 去除空格
	//str = strings.Replace(str, " ", "", -1)
	//// 去除换行符
	//str = strings.Replace(str, "\n", "", -1)
	a1 := strings.Split("1,2", ",")
	fmt.Println(a1)
	str = strings.TrimSpace(str)
	fmt.Println(str)




	str = "1.8"
	match,_=regexp.MatchString("^[0-9]+(\\.[0-9]+)+$",str)
	fmt.Println(match)
	fmt.Printf("count is:%v\n",strings.Count(str,"."))
	fmt.Println(match)
	fmt.Printf("len is:%s\n",len(str))


	fmt.Println("----------------")
	// This tests whether a pattern matches a string.
	match, _ = regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))

	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	fmt.Println(r.FindStringIndex("peach punch"))

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the
	// indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(r.Match([]byte("peach")))

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))


}

func Invokeme()  {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	re,_ := regexp.Compile(pat)
	searchIn1 := re.ReplaceAllString(searchIn,"##.#")
	fmt.Printf("%v",string(searchIn1))
}