// 转换为驼峰
package trans

import (
	"fmt"
	"strings"
)

func ToCamelCase(str string) string {
	temp := strings.Split(str, "_")
	for i, r := range temp {
		if i > 0 {
			temp[i] = strings.Title(r)
		}
	}

	return strings.Join(temp, "")
}

func main() {
	str := "the_stealth_warrior"
	fmt.Println(ToCamelCase(str))
}
