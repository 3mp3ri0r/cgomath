package cgomath

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)

	if result != 15 {
		t.Error("sum calculation incorrect")
	}
}

func TestSumZeroLength(t *testing.T) {
	numbers := []int{}
	result := Sum(numbers)

	if result != 0 {
		t.Error("sum should return 0")
	}
}
