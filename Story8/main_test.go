package main

import (
	"bootcamp/Story7/utils"
	"testing"
)

func TestGetElementsSatisfyingAnyConditionsShouldReturnElementsSatisfyingAnyOneCondition(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	want := [...]int{1, 2, 3, 5, 7, 9}

	conditions := []func(int) bool{utils.IsOdd, utils.IsPrime}

	got := getElementsSatisfyingAnyConditions(conditions, arr)

	if len(got) != len(want) {
		t.Fatalf("Want %q, got %q", want, got)
	}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Fatalf("Want %q, got %q", want[i], got[i])
		}
	}
}
