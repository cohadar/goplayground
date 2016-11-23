package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(a, b) {
		t.Error("Not really eq")
	}
	if !reflect.DeepEqual(a, c) {
		t.Error("really eq")
	}
}

func reflexInsertionSort(data reflect.Value, a, b int, less func(va, vb interface{}) bool) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(data.Index(j).Interface(), data.Index(j-1).Interface()); j-- {
			va := data.Index(j).Interface()
			vb := data.Index(j-1).Interface()
			data.Index(j).Set(reflect.ValueOf(vb))
			data.Index(j-1).Set(reflect.ValueOf(va))
		}
	}
}

func ExampleReflexInsertionSort() {
	a := []int{7, 2, 5, 4, 9, 6, 8, 1, 0, 3}
	reflexInsertionSort(reflect.ValueOf(a), 0, len(a), func(va, vb interface{}) bool {
		return va.(int) < vb.(int)
	})
	fmt.Println(a)
	// Output: [0 1 2 3 4 5 6 7 8 9]
}
