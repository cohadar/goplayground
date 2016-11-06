package list_test

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	l := new(list.List)
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	for i := 0; i < 10; i++ {
		l.PushFront(-1-i)
	}
	i := -10
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value != i {
			t.Errorf("%d != %d", e.Value, i)
		}
		i++
	}
}
