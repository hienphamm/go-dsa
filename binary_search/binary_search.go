package main

import "fmt"

func binarySearch(arr []int, target int) int {
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
	if result != -1 {
		fmt.Printf("Element %d found at index %d", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
