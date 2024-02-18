package main

import (
	"testing"
)

func TestRecursiveBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "empty array",
			arr:      nil,
			target:   12,
			expected: -1,
		},
		{
			name:     "found element",
			arr:      []int{4, 6, 12, 29},
			target:   6,
			expected: 6,
		},
		{
			name:     "not found element",
			arr:      []int{2, 5, 6, 12},
			target:   7,
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := recursiveBinarySearch(tc.arr, tc.target)
			if result != tc.expected {
				t.Errorf("recursiveBinarySearch(%v, %d) = %d; expected %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}
