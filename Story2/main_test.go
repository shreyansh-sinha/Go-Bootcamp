package main

import "testing"

func TestPrintOddNumbersReturnOddNumbers(t *testing.T) {

	arr := make([]int, 5)
	for idx := 0; idx < 5; idx++ {
		arr[idx] = idx
	}

	want := [...]int{1, 3}

	got := printOddNumbers(arr)

	if len(got) != len(want) {
		t.Fatalf("got %q, wanted %q", got, want)
	}

	for idx := 0; idx < len(want); idx++ {
		if want[idx] != got[idx] {
			t.Fatalf("got %q, wanted %q", got, want)
		}
	}
}

func TestPrintOddNumbersReturnsEmptyForEvenOnlyArray(t *testing.T) {

	arr := make([]int, 5)
	for idx := 0; idx < 5; idx++ {
		arr[idx] = 0
	}

	want := [...]int{}

	got := printOddNumbers(arr)

	if len(got) != len(want) {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}
