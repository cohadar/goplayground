package strconv_test

import (
	"strconv"
	"testing"
)

func TestA2i(t *testing.T) {
	i64, err := strconv.ParseInt("12345", 10, 16)
	if err != nil {
		t.Errorf("12345 does not fit in 16 bits?")
	}
	if i64 != 12345 {
		t.Errorf("12345 not parsed properly")
	}
	_, err = strconv.ParseInt("12345", 10, 8)
	if err == nil {
		t.Errorf("12345 fits in 16 bits?")
	}
}

func TestAppend(t *testing.T) {
	buff := []byte{}
	buff = append(buff, "€"...)
	if len(buff) != 3 {
		t.Errorf("append string to byte slice does not do utf-8 encoding")
	}
}