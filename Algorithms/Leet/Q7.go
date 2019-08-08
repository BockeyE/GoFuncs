package main

import "fmt"

func reverse(x int) int {
	if x >= 2147483647 || x <= -2147483648 {
		return 0
	}
	if x < 0 {
		return -act(-x)
	} else {
		return act(x)
	}
}
func act(x int) int {
	ret := 0
	pop := 0
	for x != 0 {
		pop = x % 10
		x = x / 10
		if (ret > (2147483647 / 10)) || ((ret == 2147483647/10) && (pop > 7)) {
			return 0
		}
		ret = ret * 10
		ret = ret + pop
	}
	return ret
}

func main() {
	fmt.Println(reverse(-214748364))
}
