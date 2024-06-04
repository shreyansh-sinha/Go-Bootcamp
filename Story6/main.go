package main

import "fmt"

func isOddAndMultipleOfThreeGreaterThanTen(value int) bool {

	return value%3 == 0 && value%2 != 0 && value > 10
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

	for idx := 0; idx < size; idx++ {
		if isOddAndMultipleOfThreeGreaterThanTen(arr[idx]) {
			fmt.Println(arr[idx])
		}
	}
}
