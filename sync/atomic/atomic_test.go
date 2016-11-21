package atomic_test

import (
	"sync/atomic"
	"testing"
)

func TestValue(t *testing.T) {
	var v atomic.Value
	v.Store(32)
	x := v.Load().(int)
	if x != 32 {
		t.Errorf("!32")
	}
}

const MIGHT = 0x7afebabedeadbeef
const MAGIC = 0x0444222277775555
const CONAN = 0x00c0fafac0c0fafa

func TestAdd(t *testing.T) {
	var g int64 = MAGIC
	atomic.AddInt64(&g, MIGHT)
	if g != MIGHT + MAGIC {
		t.Errorf("You have been killed by Troglodyte")
	}
}

func TestCAS(t *testing.T) {
	var g int64 = MIGHT
	if atomic.CompareAndSwapInt64(&g, MAGIC, CONAN) {
		t.Errorf("Conan needs no magic fool")
	}
	if !atomic.CompareAndSwapInt64(&g, MIGHT, CONAN) {
		t.Errorf("Conan is mighty")
	}
	if g != CONAN {
		t.Errorf("Where is my sword")
	}
}

