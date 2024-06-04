package main

import "fmt"

func printOddNumbers(arr []int) []int {

	len := len(arr)
	var slices []int
	for idx := 0; idx < len; idx++ {
		if arr[idx]%2 != 0 {

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

	ans := printOddNumbers(arr)

	for idx := 0; idx < len(ans); idx++ {
		fmt.Println(ans[idx])
	}
}
