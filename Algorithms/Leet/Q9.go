package main

import "fmt"

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :5 MB, 在所有 Go 提交中击败了90.91%的用户
func isPalindrome(x int) bool {
	if x < 0 || x > 2147447412 {
		return false
	}
	org := x
	t := 0
	ret := 0
	for {
		if x == 0 {
			break
		}
		t = x % 10
		x = x / 10
		ret = ret * 10
		ret = ret + t
	}
	return ret == org
}

func main() {
	fmt.Println(isPalindrome(232))
}
