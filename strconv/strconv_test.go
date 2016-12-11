package strconv_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	for i := int16(1); i != 0; i++ {
		s := strconv.Itoa(int(i))
		if j, err := strconv.Atoi(s); err != nil || int(i) != j {
			t.Error(i)
			return
		}
	}
}

func ExampleQuote() {
	fmt.Println(strconv.Quote("	č€🐸"))
	fmt.Println(strconv.Quote("\x09" + "\xc4\x8d" + "\xe2\x82\xac" + "\xf0\x9f\x90\xb8"))
	fmt.Println(strconv.Quote("\u0009" + "\u010D" + "\u20AC" + "🐸"))
	fmt.Println(strconv.Quote(`	"'`))
	// Output: 
	// "\tč€🐸"
	// "\tč€🐸"
	// "\tč€🐸"
	// "\t\"'"
}

func ExampleQuoteRuneToASCII() {
	fmt.Println(strconv.QuoteRuneToASCII('A'))
	fmt.Println(strconv.QuoteRuneToASCII(0x11))
	fmt.Println(strconv.QuoteRuneToASCII('€'))
	fmt.Println(strconv.QuoteRuneToASCII('🐸'))
	// Output:
	// 'A'
	// '\x11'
	// '\u20ac'
	// '\U0001f438'
}