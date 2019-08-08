package main

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
)

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	//arrList := [numRows]list.List{}  动态长度 的数组不能这么写
	//arrList := make([]*list.List, numRows)  这里的数组，如果是[]list.List,
	//则传入的是一个对象值，而不是指针，换句话说，取出来时，没办法拿到该对象本身的引用指针
	//因此必须指定数组类型为类型指针
	arrList := make([]*list.List, numRows)
	for i := 0; i < numRows; i++ {
		arrList[i] = list.New()
	}
	index := 0
	for {
		if index >= len(s) {
			break
		}
		for j := 0; j < numRows && index < len(s); j++ {
			//arrList[j].PushBack(strconv.Itoa(int(s[index])))
			//这个iota函数的作用，是把本用来代表字符号55的数字，转变为‘55’的显示字符串，如果重新拿到55对应字符，直接强转即可
			t := s[index]
			arrList[j].PushBack(string(t))
			index++
		}
		for j := numRows - 2; j > 0 && index < len(s); j-- {
			t := s[index]
			arrList[j].PushBack(string(t))
			index++
		}
	}
	buider := bytes.Buffer{}
	for k := 0; k < numRows; k++ {
		ar := arrList[k]
		fmt.Println(ar.Len())
		for i := ar.Front(); i != nil; i = i.Next() {
			c, _ := i.Value.(string)
			// 这个是go中的跨类型强转，go要求必须用断言，第二个变量是一个bool，用来判断是否强转成功，可以进行判断控制
			buider.WriteString(c)
		}
	}
	return buider.String()

}

func main() {
	fmt.Println(reflect.TypeOf("FFF"[2]))
	s := convert("LEETCODEISHIRING", 4)
	fmt.Println(s)

}
