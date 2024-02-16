package main

import "github.com/hienphamm/go-dsa/verify"

func linearSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}

func main() {
	arr := []int{1, 5, 7, 40, 100}
	target := 7
	result := linearSearch(arr, target)
	verify.Verify(result, target)
}
