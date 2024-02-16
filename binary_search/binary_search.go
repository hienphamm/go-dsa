package main

import (
	"github.com/hienphamm/go-dsa/verify"
)

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 5, 10, 20, 50, 80, 200}
	target := 10
	result := binarySearch(arr, target)
	verify.Verify(result, target)
}
