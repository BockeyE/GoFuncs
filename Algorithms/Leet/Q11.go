package main

import "math"

//执行用时 :16 ms, 在所有 Go 提交中击败了95.81%的用户
//内存消耗 :5.7 MB, 在所有 Go 提交中击败了64.91%的用户
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	tem := float64(0)
	max := float64(0)
	for start != end {
		tem = (math.Min(float64(height[start]), float64(height[end]))) * float64(end-start)
		max = math.Max(tem, max)
		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}
	return int(max)
}
