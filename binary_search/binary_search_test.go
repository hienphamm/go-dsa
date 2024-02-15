package main

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 4, 10, 40}

	testsFound := []struct {
		name   string
		target int
		index  int
	}{
		{"ElementFoundAtStart", 2, 0},
		{"ElementFoundInMiddle", 10, 3},
		{"ElementFoundAtEnd", 40, 4},
	}

	for _, tc := range testsFound {
		t.Run(tc.name, func(t *testing.T) {
			index := binarySearch(arr, tc.target)
			if index != tc.index {
				t.Errorf("Expected index %d, got %d", tc.index, index)
			}
		})
	}

	testsNotFound := []struct {
		name   string
		target int
		index  int
	}{
		{"ElementNotFoundLessThanMin", 1, -1},
		{"ElementNotFoundGreaterThanMax", 100, -1},
		{"ElementNotFoundBetweenValues", 5, -1},
	}

	for _, tc := range testsNotFound {
		t.Run(tc.name, func(t *testing.T) {
			index := binarySearch(arr, tc.target)
			if index != tc.index {
				t.Errorf("Expected index %d, got %d", tc.index, index)
			}
		})
	}
}
