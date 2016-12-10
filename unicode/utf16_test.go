package utf16_test

import (
	"testing"
	"unicode/utf16"
)

func MyEncode(s []rune) []uint16 {
	r := s[0]
	if r < 0 || r > 0x10FFFF {
		return []uint16{0xFFFD}
	}	
	if r >= 0xD800 && r <= 0xDFFF {
		return []uint16{0xFFFD}
	}		
	if r < 0x10000 {
		return []uint16{uint16(r)}
	}
	r -= 0x10000
	lo := uint16((r >> 10) & 0x3FF | 0xD800)
	hi := uint16(r & 0x3FF | 0xDC00)
	return []uint16{lo, hi}
}

func TestMyEncode(t *testing.T) {
	for i := rune(-10); i <= 0x10FFFF+10; i++ {
		in := []rune{i}
		d1 := utf16.Encode(in)
		d2 := MyEncode(in)
		if len(d1) != len(d2) {
			t.Errorf("%#U len %d != %d", i, len(d1), len(d2))
			return
		}
		for j := 0; j < len(d1); j++ {
			if d1[j] != d2[j] {
				t.Errorf("%#U data %X %X -  %X %X\n", i, d1[0], d1[1], d2[0], d2[1])
				return
			}
		}
	}
}
