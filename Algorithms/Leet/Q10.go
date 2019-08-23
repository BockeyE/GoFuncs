package main

//执行用时 :20 ms, 在所有 Go 提交中击败了52.54%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了83.33%的用户

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	var firstm bool = (!(s == "")) &&
		(p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return (isMatch(s, p[2:])) || (firstm && isMatch(s[1:], p))
	} else {
		return firstm && isMatch(s[1:], p[1:])
	}

}
