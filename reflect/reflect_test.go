package reflect_test

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4}
	if reflect.DeepEqual(a, b) {
		t.Error("Not really eq")
	}
}
