package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tem := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		vi := nums[i]

		value, ok := tem[vi]
		if ok {
			if target == 2*vi {
				//var arr []int
				arr := []int{i, value}
				return arr
			}
		}
		tem[vi] = i
	}
	fmt.Println("map ", tem)
lop2:
	for i := 0; i < len(nums); i++ {
		vi := nums[i]

		value, ok := tem[target-vi]
		if ok {
			fmt.Println("values ", vi, ok)
			fmt.Println("targets ", target, ok)
			if target == 2*vi {
				fmt.Println("enter continue")
				continue lop2
			}
			fmt.Println("not out")
			arr := []int{i, value}
			return arr
		}
	}
	return nil

}
func main() {
	arr := []int{3, 3}
	c := twoSum(arr, 6)
	fmt.Println(c)

}
