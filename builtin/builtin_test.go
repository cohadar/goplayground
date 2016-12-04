package builtin_test

import (
	"fmt"
	"testing"
)

func ExampleAppend() {
	a := []int{2, 3, 5, 7, 11}
	a = append(a, 13, 17, 19)
	fmt.Println(a)
	// Output: [2 3 5 7 11 13 17 19]
}

func ExampleAppend_slice() {
	a := []int{2, 3, 5, 7, 11}
	b := []int{13, 17, 19}
	a = append(a, b...)
	fmt.Println(a)
	// Output: [2 3 5 7 11 13 17 19]
}

func ExampleAppend_string() {
	a := []byte("hello ")
	a = append(a, "world"...)
	fmt.Println(string(a))
	// Output: hello world
}

const (
	m0 = 1<<iota - 1
	m1
	m2
	m3
	m4
)

func ExampleMasks() {
	fmt.Println(m0, m1, m2, m3, m4)
	// Output: 0 1 3 7 15
}

func ExampleCap() {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[0:3]
	var c []int = nil
	d := make(chan int, 7)
	fmt.Print(cap(a))
	fmt.Print(cap(&a))
	fmt.Print(cap(b))
	fmt.Print(cap(c))
	fmt.Print(cap(d))
	// Output: 55507
}

func TestClose(t *testing.T) {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	for i := 0; i < 5; i++ {
		v, ok := <-c
		if v != i || !ok {
			t.Error(i)
		}
	}
	_, ok := <-c
	if ok {
		t.Error("should be closed")
	}
	
}

func ExampleCopy() {
	a := make([]byte, 3, 10)
	n := copy(a, "njak")
	fmt.Println(n, string(a))
	// Output: 3 nja
}

func TestComplex(t *testing.T) {
	a := complex(3, -4)
	if a != 3-4i {
		t.Error(a)
	}
	if real(a) != 3 {
		t.Error("unreal")
	}
	if imag(a) != -4 {
		t.Error("unimag")
	}
}

func TestDelete(t *testing.T) {
	m := make(map[string]int)
	m["aaa"] = 1
	m["bbb"] = 22
	m["ccc"] = 333
	v, ok := m["bbb"]
	if !ok {
		t.Error("not ok")
	}
	if v != 22 {
		t.Error("bad bad thing")
	}	
	delete(m, "bbb")
	_, ok = m["bbb"]
	if ok {
		t.Error("ok")
	}
}

func TestLen(t *testing.T) {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c <- i * 23
		if len(c) != i + 1 {
			t.Error(i, len(c))
		}
	}
	for i := 0; i < 5; i++ {
		<-c
		if len(c) != 4 - i {
			t.Error(i, len(c))
		}
	}
	s := "hello €"
	if len(s) != len([]byte(s)) {
		t.Error("hello €")
	}
}

func TestMake(t *testing.T) {
	a := make([]int, 5, 11)
	if len(a) != 5 || cap(a) != 11 {
		t.Error(a)
	}
	b := make([]int, 7)
	if len(b) != 7 || cap(b) != 7 {
		t.Error(b)
	}
	m := make(map[string]int, 10)
	// cap does not work on maps
	if len(m) != 0 {
		t.Error(len(m))
	}
	c := make(chan int, 10)
	if cap(c) != 10 {
		t.Error(cap(c))
	}
	cl := make(chan<- int, 10)
	cr := make(<-chan int, 10)
	fmt.Println(len(cl), len(cr))
}

func TestNew(t *testing.T) {
	pa := new([7]int)
	if len(pa) != 7 {
		t.Error("not 7")
	}
}

func ExampleDefer() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }	
	// Output: 3210
}

func ExamplePanic() {
    v := f()
    fmt.Printf("t%v", v)
    // Output: #g0g1g2g3!d3d2d1d0r4t0
}

func f() int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("r%v", r)
        }
    }()
    fmt.Print("#")
    g(0)
    fmt.Print("never reach")
    return 9
}

func g(i int) {
    if i > 3 {
        fmt.Print("!")
        panic(i)
    }
    defer fmt.Printf("d%v", i)
    fmt.Printf("g%v", i)
    g(i + 1)
}
