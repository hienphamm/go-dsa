package verify

import "fmt"

func Verify(result int, target int) {
	if result != -1 {
		fmt.Printf("Element %d found at index %d", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
