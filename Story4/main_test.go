package main

import "testing"

func TestOddPrimeNumberShouldReturnFalseForEvenNumber(t *testing.T) {

	got := isOddPrime(42)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}

func TestOddprimeNumberShouldReturnFalseForEvenprime(t *testing.T) {

	got := isOddPrime(2)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}
