package main

import "fmt"

func main() {
	fmt.Println("abcde"[3:4])
	fmt.Println('\u0000')
	fmt.Println('\u0020')
}
