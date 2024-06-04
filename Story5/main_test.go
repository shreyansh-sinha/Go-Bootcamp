package main

import "testing"

func TestEvenAndMultipleOfFiveShouldReturnTrueForMultipleOfTen(t *testing.T) {

	got := isEvenAndMultipleOfFive(20)

	if got != true {
		t.Fatalf("got %t, wanted %t", false, true)
	}
}

func TestEvenAndMultipleOfFiveShouldReturnFalseForOddMultipleOfFive(t *testing.T) {

	got := isEvenAndMultipleOfFive(15)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}
