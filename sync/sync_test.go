package sync_test

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	x := int32(0)
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(px *int32) {
			for j := 0; j < 1e6; j++ {
				m.Lock()
				x++
				m.Unlock()
			}
			wg.Done()
		}(&x)
	}
	wg.Wait()
	if x != 1e7 {
		t.Errorf("%d", x)
	}
}
