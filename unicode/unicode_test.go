package unicode_test

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestHex(t *testing.T) {
	for _, r := range "0123456789abcdefABCDEF" {
		if !unicode.Is(unicode.ASCII_Hex_Digit, r) {
			t.Errorf("%c is not a hex digit", r)
		}
	}
	for _, r := range "!@@#$%%^&*()_+" {
		if unicode.Is(unicode.ASCII_Hex_Digit, r) {
			t.Errorf("%c is a hex digit?", r)
		}
	}
}

func TestIsz(t *testing.T) {
	if !unicode.IsLower('x') {
		t.Errorf("x is not lowercase?")
	}
	if unicode.ToUpper('x') != 'X' {
		t.Errorf("xX")
	}
}

func TestEncoding(t *testing.T) {
	s := "€euro€xxx€"
	ar := []rune{}
	for b := []byte(s); len(b) > 0; {
		r, size := utf8.DecodeRune(b)
		b = b[size:]
		ar = append(ar, r)
	}
	if s != string(ar) {
		t.Errorf("%q != %q", s, string(ar))
	}
}
