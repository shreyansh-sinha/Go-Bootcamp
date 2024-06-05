package main

import (
	"bootcamp/Story7/utils"
	"fmt"
)

func getElementsSatisfyingAnyConditions(conditions []func(int) bool, arr []int) []int {

	size := len(arr)
	var slices []int
	for idx := 0; idx < size; idx++ {
		if utils.Any(conditions, arr[idx]) {
			slices = append(slices, arr[idx])
		}
	}
	return slices
}

func main() {

	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	for idx := 0; idx < size; idx++ {
		var value int
		fmt.Scan(&value)
		arr[idx] = value
	}

	conditions := []func(int) bool{utils.IsOdd, utils.IsPrime}
	ans := getElementsSatisfyingAnyConditions(conditions, arr)

	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
