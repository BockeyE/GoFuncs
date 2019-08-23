package main

import (
	"fmt"
	"strings"
)

//执行用时 :4  ms, 在所有 Go 提交中击败了96.70%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了38.57%的用户

func romanToInt(s string) int {
	values := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	start := 0
	for i := 0; i < len(values); i++ {
		if strings.HasPrefix(s, strs[i]) {
			start = i
			break
		}
	}
	aim := 0
	for start < len(values) {
		if strings.HasPrefix(s, strs[start]) {
			aim = aim + values[start]
			s = s[len(strs[start]):]
			continue
		} else {
			start++
		}
	}
	return aim
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
