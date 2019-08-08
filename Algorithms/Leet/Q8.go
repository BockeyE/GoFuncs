package main

import (
	"fmt"
	"strconv"
	"strings"
)

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了45.45%的用户
func myAtoi(str string) int {
	//trim函数用法，第一个是目标，第二个是首位需要去掉的字符，可以自定义
	str = strings.Trim(str, " ")
	fmt.Println(str)
	if strings.HasPrefix(str, "+-") {
		return 0
	}
	if strings.HasPrefix(str, "+") {
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}
	start := 0
	end := len(str)
	if !(str[0] == 45 || (str[0] >= 48 && str[0] <= 57)) {
		return 0
	}
	for k := start + 1; k < len(str); {
		if !(str[k] >= 48 && str[k] <= 57) {
			end = k
			break
		}
		k++
	}
	sub := str[start:end]
	if len(sub) == 1 && sub == "-" {
		return 0
	}
	ints, _ := strconv.Atoi(sub)
	if ints > 2147483647 {
		return 2147483647
	}
	if ints < -2147483648 {
		return -2147483648
	}
	return ints
}

func main() {
	fmt.Println(myAtoi("   -42d "))
}
