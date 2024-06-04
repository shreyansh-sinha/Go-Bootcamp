package main

import (
	"bootcamp/Story7/utils"
	"fmt"
)

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
	for idx := 0; idx < size; idx++ {
		if utils.All(conditions, arr[idx]) {
			fmt.Println(arr[idx])
		}
	}
}
