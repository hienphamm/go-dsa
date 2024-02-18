package main

import "github.com/hienphamm/go-dsa/verify"

func recursiveBinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	} else {
		mid := len(arr) / 2
		if arr[mid] == target {
			return target
		} else if arr[mid] < target {
			return recursiveBinarySearch(arr[mid+1:], target)
		} else {
			return recursiveBinarySearch(arr[:mid], target)
		}
	}
}

func main() {
	arr := []int{3, 5, 7, 12, 20, 22}
	target := 12
	result := recursiveBinarySearch(arr, target)
	verify.Verify(result, target)
}
