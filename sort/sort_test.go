package sort

import (
	"fmt"
	"sort"
)

func ExampleInts() {
	x := []int{2, 6, 3, 7, 9, 2, 5, 0, 9, 1, 2}
	sort.Ints(x)
	fmt.Println(x)
	// Output: [0 1 2 2 2 3 5 6 7 9 9]
}

func ExampleIntSlice() {
	x := []int{2, 6, 3, 7, 9, 2, 5, 0, 9, 1, 2}
	sort.IntSlice(x).Sort()
	fmt.Println(x)
	// Output: [0 1 2 2 2 3 5 6 7 9 9]
}

func ExampleIntSlice2() {
	x := []int{2, 6, 3, 7, 9, 2, 5, 0, 9, 1, 2}
	sort.Sort(sort.IntSlice(x))
	fmt.Println(x)
	// Output: [0 1 2 2 2 3 5 6 7 9 9]
}

func ExampleSearch() {
	x := []int{0, 1, 2, 2, 2, 3, 5, 6, 7, 9, 9}
	i := sort.Search(len(x), func(i int) bool {
		return x[i] > 2
	})
	fmt.Println(i)
	// Output: 5
}

type njak struct {
	x rune
	y int
}

func (o njak) String() string {
	return fmt.Sprintf("%c%d", o.x, o.y)
}

type njakByNatural []njak

func (o njakByNatural) Len() int {
	return len(o)
}

func (o njakByNatural) Less(i, j int) bool {
	if o[i].x == o[j].x {
		return o[i].y < o[j].y
	}
	return o[i].x < o[j].x
}

func (o njakByNatural) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func ExampleCustom() {
	n := []njak{{'c', 2}, {'a', 3}, {'b', 2}, {'a', 1}, {'c', 3}, {'b', 5}}
	sort.Sort(sort.Reverse(njakByNatural(n)))
	fmt.Println(n)
	// Output: [c3 c2 b5 b2 a3 a1]
}
