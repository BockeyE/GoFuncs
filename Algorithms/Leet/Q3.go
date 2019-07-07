package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	set := make(map[uint8]int)
	left := 0
	max := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		_, ok := set[c]
		if ok {
			for {
				if s[left] != c {
					delete(set, s[left])
					left++
				} else {
					left++
					break
				}
			}
		}
		set[c] = i
		leth := (i - left) + 1
		if leth > max {
			max = leth
		}
	}
	return max
}
func main() {
	s := lengthOfLongestSubstring("tmmzuxt")
	//s := lengthOfLongestSubstring("abcabcabc")
	fmt.Println(s)
}
