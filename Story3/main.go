package main

import (
	"fmt"
	"math"
)

func isPrime(value int) bool {

	if value < 2 {
		return false
	}

	if value == 2 {
		return true
	}

	sqrtVal := int(math.Sqrt(float64(value)))

	for val := 2; val <= sqrtVal; val++ {
		if value%val == 0 {
			return false
		}
	}

	return true
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
		if isPrime(arr[idx]) {
			fmt.Println(arr[idx])
		}
	}
}
