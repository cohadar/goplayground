package utf8_test

import (
	"encoding/hex"
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

func MyEncodeRune(p []byte, r rune) int {
	if r < 0x80 {
		p[0] = byte(r)
		return 1
	}
	if r < 0x800 {
		p[0] = byte((r >> 6) & 0x1F | 0xC0)
		p[1] = byte(r & 0x3F | 0x80)
		return 2
	}
	if r < 0x10000 {
		if r >= 0xD800 && r <= 0xDFFF {
			r = 0xFFFD
		}
		p[0] = byte((r >> 12) & 0xF | 0xE0)
		p[1] = byte((r >> 6) & 0x3F | 0x80)
		p[2] = byte(r & 0x3F | 0x80)
		return 3
	}
	p[0] = byte((r >> 18) & 0x7 | 0xF0)
	p[1] = byte((r >> 12) & 0x3F | 0x80)
	p[2] = byte((r >> 6) & 0x3F | 0x80)
	p[3] = byte(r & 0x3F | 0x80)	
	return 4
}

func TestMyEncode(t *testing.T) {
	p1 := make([]byte, utf8.UTFMax)
	p2 := make([]byte, utf8.UTFMax)
	for i := rune(0); i <= utf8.MaxRune; i++ {
		d1 := utf8.EncodeRune(p1, i)
		d2 := MyEncodeRune(p2, i)
		if d1 != d2 {
			t.Errorf("%#U len %d != %d", i, d1, d2)
			return
		}
		for j := 0; j < d1; j++ {
			if p1[j] != p2[j] {
				t.Errorf("%#U data %v %v\n", i, hex.EncodeToString(p1[:d1]), hex.EncodeToString(p2[:d2]))
				return
			}
		}
	}
}
