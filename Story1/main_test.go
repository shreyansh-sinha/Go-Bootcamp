package main

import "testing"

func TestPrintEvenNumbersReturnsEvenNumbers(t *testing.T) {

	arr := make([]int, 5)
	for idx := 0; idx < 5; idx++ {
		arr[idx] = idx
	}
	want := []int{0, 2, 4}

	got := printEvenNumbers(arr)

	if len(want) != len(got) {
		t.Fatalf("got %q, wanted %q", got, want)
	}

	len := len(want)

	for idx := 0; idx < len; idx++ {

		if got[idx] != want[idx] {
			t.Fatalf("got %q, wanted %q", got, want)
		}
	}
}

func TestPrintEvenNumbersReturnsEmptyListForOddOnlyIntegers(t *testing.T) {

	arr := make([]int, 5)
	for idx := 0; idx < 5; idx++ {
		arr[idx] = 1
	}

	want := [0]int{}

	got := printEvenNumbers(arr)

	if len(got) != len(want) {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}
