package main

import "testing"

func TestOddNumberGreaterThanTenAndMultipleOfThreeShouldReturnFalseForEvenNumber(t *testing.T) {

	got := isOddAndMultipleOfThreeGreaterThanTen(42)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}

func TestOddNumberGreaterThanTenAndMultipleOfThreeShouldReturnTrueForOddNumberGreaterThanTenAndMultipleOfThree(t *testing.T) {

	got := isOddAndMultipleOfThreeGreaterThanTen(45)

	if got != true {
		t.Fatalf("got %t, wanted %t", false, true)
	}
}

func TestOddNumberGreaterThanTenAndMultipleOfThreeShouldReturnFalseForOddNumberLessThanTen(t *testing.T) {

	got := isOddAndMultipleOfThreeGreaterThanTen(9)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}

func TestOddNumberGreaterThanTenAndMultipleOfThreeShouldReturnFalseForEvenNumberGreaterThanTen(t *testing.T) {

	got := isOddAndMultipleOfThreeGreaterThanTen(20)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}

func TestOddNumberGreaterThanTenAndMultipleOfThreeShouldReturnFalseForOddNumberGreaterThanTenAndNotMultipleOfThree(t *testing.T) {

	got := isOddAndMultipleOfThreeGreaterThanTen(17)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}
