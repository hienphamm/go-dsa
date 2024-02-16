package main

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			target:   2,
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := linearSearch(tc.arr, tc.target)
			if result != tc.expected {
				t.Errorf("linearSearch(%v, %d) = %d; expected %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}
