package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A := nums1
	B := nums2
	if len(nums1) > len(nums2) {
		B := nums1
		A := nums2
	}
	m := len(A)
	n := len(B)
	iMin := 0
	iMax := 0
	halfLen := (m + n + 1) / 2
	for {
		if iMin > iMax {
			break
		}

		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && B[j-1] > A[i] {
			iMin = i + 1
		} else if i > iMin && A[i-1] > B[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = math.Max(float64(A[i-1]), float64(B[j-1]))
			}

		}

	}
}
