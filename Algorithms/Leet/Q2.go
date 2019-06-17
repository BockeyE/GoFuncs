package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tem := 0
	promote := 0
	tem = l1.Val + l2.Val
	if tem > 9 {
		promote = 1
		tem = tem % 10
	}
	res := &ListNode{tem, nil}
	//这里如果不加&，则拿到一个值对象，会导致链无法连接下去，必须变成指针对象，才可以向下指向

	org := res
	for {
		tem = 0
		if l1.Next == nil && l2.Next == nil {
			if promote == 1 {
				res.Next = &ListNode{1, nil}
			}
			break
		}
		res.Next = &ListNode{0, nil}
		res = res.Next
		//&是取指针，*是通过指针拿到对象本身
		if promote == 1 {
			res.Val = 1
		}
		if l2.Next == nil {
			l1 = l1.Next
			tem = l1.Val
		} else if l1.Next == nil {
			l2 = l2.Next
			tem = l2.Val
		} else {
			l1 = l1.Next
			l2 = l2.Next
			tem = l2.Val + l1.Val
		}
		sums := res.Val + tem
		if sums > 9 {
			promote = 1
			sums = sums % 10
		} else {
			promote = 0
		}
		res.Val = sums
		fmt.Println("did once")
	}
	return org

}
func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	numbers := addTwoNumbers(&l1, &l2)
	fmt.Println(numbers.Val)
	fmt.Println(numbers.Next.Val)
	fmt.Println(numbers.Next.Next.Val)
}
