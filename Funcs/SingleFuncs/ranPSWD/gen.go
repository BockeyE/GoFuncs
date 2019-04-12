package main

import (
	"math/rand"
	"time"
)

var pass_len int = 8

var head = []rune("+-!@#")

var tail = []rune("abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789")

func main() {
	rand.Seed(time.Now().Unix())
	pass := make([]rune, pass_len)
	pass[0] = head[rand.Intn(5)]
	for i := 1; i < pass_len; i++ {
		pass[i] = tail[rand.Intn(len(tail))]
	}
	print(string(pass))
}
