package builtin_test

import "fmt"

func ExampleAppend() {
	a := []int{2, 3, 5, 7, 11}
	a = append(a, 13, 17, 19)
	fmt.Println(a)
	// Output: [2 3 5 7 11 13 17 19]
}
