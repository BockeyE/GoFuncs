package main

import (
	"fmt"
	"sort"
	s "strings"
	"time"
)

var p = fmt.Println

func main() {

	// 下面是strings包里面提供的一些函数实例。注意这里的函数并不是
	// string对象所拥有的方法，这就是说使用这些字符串操作函数的时候
	// 你必须将字符串对象作为第一个参数传递进去。
	p(s.Contains("test", "es"))
	p(s.Count("adadadadass", "s"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()

	//这两个方法不是string包函数
	//获取字符串长度
	p("Len: ", len("hello"))
	//获取指定索引的字符
	p("Char:", "hello"[1])

	p("=========================================")
	p("123")
	str := []string{"c", "a", "d", "b"}
	sort.Strings(str)
	fmt.Println("strings:", str)

	number := []int{1, 23, 4, 9, 5}
	sort.Ints(number)
	fmt.Println("int: ", number)

	done := sort.IntsAreSorted(number)
	fmt.Println("sorted : ", done)
	p("=========================================")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go work(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}

func work(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("work:", id, "|process job:", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
