package main

import "testing"

func TestPrimeNumberShouldReturnTrueForPrimeNumber(t *testing.T) {

	got := isPrime(23)

	if got != true {
		t.Fatalf("got %t, wanted %t", false, true)
	}
}

func TestPrimeNumberShouldReturnFalseForCompositenumber(t *testing.T) {

	got := isPrime(1)

	if got != false {
		t.Fatalf("got %t, wanted %t", true, false)
	}
}
