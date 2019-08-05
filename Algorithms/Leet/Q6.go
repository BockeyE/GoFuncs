package main

import "container/list"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	arrList := [numRows]list.List{}
	for i := 0; i < numRows; i++ {
		arrList[i] = *list.New()
	}
	index := 0
	period := (numRows - 1) * 2
	for {
		if index >= len(s) {
			break
		}
		if index%period == 0 {
			for j := 0; j < numRows && index < len(s); j++ {
				ar := arrList.
			}
			for j := numRows - 2; j > 0 && index < len(s); j-- {

			}
		}
	}
}
