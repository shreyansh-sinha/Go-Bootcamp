package main

import "fmt"

func isEvenAndMultipleOfFive(value int) bool {

	return value%10 == 0
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
		if isEvenAndMultipleOfFive(arr[idx]) {
			fmt.Println(arr[idx])
		}
	}
}
