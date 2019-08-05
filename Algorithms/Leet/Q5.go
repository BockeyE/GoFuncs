package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	} else if len(s) == 2 {
		if s[1] == s[0] {
			return s
		} else {
			return s[0:1]
		}
	}
	res := s[0:1]

	for i := 1; i < (len(s) - 1); i++ {
		if s[i-1] == s[i+1] {
			tem := getSub(s, i-1, i+1)
			if len(tem) > len(res) {
				res = tem
			}
		}
		if s[i-1] == s[i] {
			tem := getSub(s, i-1, i)
			if len(tem) > len(res) {
				res = tem
			}
		}
		if s[i] == s[i+1] {
			tem := getSub(s, i, i+1)
			if len(tem) > len(res) {
				res = tem
			}
		}
	}
	return res
}

func getSub(s string, i, j int) string {
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}
	i++
	//go的取子字符串功能，与py类似，i:j是指起始位置，而不是java中substring的起点与长度
	return s[i:j]
}

func main() {
	s := longestPalindrome("abbbbbbadd")
	fmt.Println(s)
}
