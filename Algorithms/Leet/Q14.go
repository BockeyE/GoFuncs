package main

import "fmt"

//执行用时 :0  ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了60.21%的用户

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	c := strs[0]
	for i := 0; i < len(strs); i++ {
		c = getTwo(c, strs[i])
	}
	return c
}
func getTwo(s1, s2 string) string {
	if s1 == s2 {
		return s1
	}
	end := 0
	for end < len(s1) && end < len(s2) {
		if s1[end] == s2[end] {
			end++
		} else {
			break
		}
	}
	return s1[0:end]
}
func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}
