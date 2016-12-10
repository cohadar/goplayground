package utf8_test

import (
	"testing"
	"unicode/utf8"
)

func assertEncodingLen(t *testing.T, run rune, len int) {
	p := make([]byte, utf8.UTFMax)
	d := utf8.EncodeRune(p, run)
	if d != len {
		t.Errorf("'%c', expected: %d, got: %d", run, len, d)
	}
}

func TestEncode(t *testing.T) {
	assertEncodingLen(t, 'X', 1)
	assertEncodingLen(t, 'č', 2)
	assertEncodingLen(t, '€', 3)
	assertEncodingLen(t, '🐸', 4)
}
