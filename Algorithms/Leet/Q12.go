package main

import (
	"bytes"
)

//执行用时 :4  ms, 在所有 Go 提交中击败了98.13%的用户
//内存消耗 :3.4 MB, 在所有 Go 提交中击败了47.50%的用户

func intToRoman(num int) string {
	values := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var buffer bytes.Buffer
	for i := 0; i < len(values); i++ {
		if num/values[i] == 0 {
			continue
		}
		k := num / values[i]
		for k != 0 {
			buffer.WriteString(strs[i])
			k--
		}
		num = num % values[i]
		if num == 0 {
			break
		}
	}
	return buffer.String()
}
