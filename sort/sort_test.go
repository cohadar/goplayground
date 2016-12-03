package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
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

func randomSlice(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Int()
	}
	return ret
}

func BenchmarkSort1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := randomSlice(1e6)
		b.StartTimer()
		sort.Ints(data)
		b.StopTimer()
		if !sort.IsSorted(sort.IntSlice(data)) {
			b.Errorf("unsorted")
		}
	}
}

// rawSort time is 40% of interface sort time. <------------<<
// go interfaces are VERY efficient, therefore templates are a waste of brain power.
func BenchmarkRawSort1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := randomSlice(1e6)
		b.StartTimer()
		rawSort(data)
		b.StopTimer()
		if !sort.IsSorted(sort.IntSlice(data)) {
			b.Errorf("unsorted")
		}
	}
}

// MODIFIED CODE BELOW THIS LINE

func rawSort(data []int) {
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(data)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSort(data, 0, n, maxDepth)
}

func swap(data []int, l, r int) {
	tmp := data[l]
	data[l] = data[r]
	data[r] = tmp
}

func less(data []int, l, r int) bool {
	return data[l] < data[r]
}

func quickSort(data []int, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if less(data, i, i-6) {
				swap(data, i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

func heapSort(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		swap(data, first, first+i)
		siftDown(data, lo, i, first)
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(data, first+child, first+child+1) {
			child++
		}
		if !less(data, first+root, first+child) {
			return
		}
		swap(data, first+root, first+child)
		root = child
	}
}

func doPivot(data []int, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && less(data, a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !less(data, pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && less(data, pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		swap(data, b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !less(data, pivot, hi-1) { // data[hi-1] = pivot
			swap(data, c, hi-1)
			c++
			dups++
		}
		if !less(data, b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !less(data, m, pivot) { // data[m] = pivot
			swap(data, m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !less(data, b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && less(data, a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			swap(data, a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	swap(data, pivot, b-1)
	return b - 1, c
}

// Insertion sort
func insertionSort(data []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(data, j, j-1); j-- {
			swap(data, j, j-1)
		}
	}
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data []int, m1, m0, m2 int) {
	// sort 3 elements
	if less(data, m1, m0) {
		swap(data, m1, m0)
	}
	// data[m0] <= data[m1]
	if less(data, m2, m1) {
		swap(data, m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if less(data, m1, m0) {
			swap(data, m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}
