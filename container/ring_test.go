package ring_test

import (
	"container/ring"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		if i != r.Value {
			t.Errorf("%d != %d", i, r.Value)
		}
		r = r.Next()
	}
}
