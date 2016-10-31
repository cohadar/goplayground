package unsafe_test

import (
	"testing"
	"unsafe"
)

// Note: I tried this test only on my 64bit macbook
func TestSizeof(t *testing.T) {
	if unsafe.Sizeof(int32(12345)) != 4 {
		t.Errorf("int32 is not 4 bytes?")
	}
	if unsafe.Sizeof(int64(0x12393254234)) != 8 {
		t.Errorf("int64 is not 8 bytes?")
	}
	if unsafe.Sizeof(rune('x')) != 4 {
		t.Errorf("rune is not 4 bytes?")
	}
	if unsafe.Sizeof(false) != 1 {
		t.Errorf("bool is not a byte?")
	}
	int_size := unsafe.Sizeof(int(12345))
	str := "hello"
	ptr_size := unsafe.Sizeof(&str)
	if ptr_size != int_size {
		t.Errorf("pointer size != int size") // this can actually happen on some architectures
	}
	if unsafe.Sizeof("hello") != int_size+ptr_size {
		t.Errorf("string is not int size + ptr size?")
	}
	slice := []int{}
	slice_size := unsafe.Sizeof(slice)
	if slice_size != ptr_size+int_size+int_size {
		t.Errorf("slice is not a pointer + len + cap?")
	}
	mapp := make(map[string]int)
	if unsafe.Sizeof(mapp) != ptr_size {
		t.Errorf("map is not a pointer?")
	}
	ch := make(chan int)
	if unsafe.Sizeof(ch) != ptr_size {
		t.Errorf("channel is not a pointer?")
	}
	add := func(a, b int) int {
		return a + b
	}
	if unsafe.Sizeof(add) != ptr_size {
		t.Errorf("function pointer != data pointer?") // this can also happen on some architectures
	}
	if unsafe.Sizeof(uintptr(0)) != ptr_size {
		t.Errorf("universal pointer != data pointer?") // this can also happen on some architectures
	}
	s1 := struct {
		b bool
		c int
	}{false, 3}
	if unsafe.Sizeof(s1) != 2 * int_size {
		t.Errorf("struct did not use alignment?")
	}
	s2 := struct {
		a int16
		b bool
		c int32
		d int
	}{0, false, 0, 0}
	if unsafe.Sizeof(s2) != 2 * int_size {
		t.Errorf("struct did not use field packing?")
	}	
}

func TestAlighAndOffset(t *testing.T) {
	s := struct {
		a bool  // 0: 1
		// pad 3
		b int32 // 4: 4
		c int16 // 8: 2
		// pad 6
		d string  // 16: 16
	}{}
	ptr_size := unsafe.Sizeof(&s)
	if unsafe.Alignof(s) != ptr_size {
		t.Errorf("struct not aligned to biggest field (string data pointer)")
	}
	if unsafe.Alignof(s.b) != 4 {
		t.Errorf("struct has no padding")
	}	
	if unsafe.Alignof(s.c) != 2 {
		t.Errorf("struct not packed")
	}	
	if unsafe.Offsetof(s.b) != 4 {
		t.Errorf("no padding by int32")
	}
	if unsafe.Offsetof(s.c) != 8 {
		t.Errorf("no padding by int36")
	}	
	if unsafe.Offsetof(s.d) != 16 {
		t.Errorf("no padding by string")
	}		
}
