package verify

import "fmt"

func Verify(result int, target int) {
	if result != -1 {
		fmt.Printf("Element %d found in the array", target)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
