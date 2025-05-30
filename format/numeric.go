package format

import (
	"fmt"
	"strconv"
)

func ParseNumeric(s string) (n int, nRunes int, err error) {
	// allow leading spaces
	for nRunes < len(s) && s[nRunes] == ' ' {
		nRunes++
	}
	start := nRunes
	for nRunes < len(s) && s[nRunes] >= '0' && s[nRunes] <= '9' {
		nRunes++
	}
	if start == nRunes {
		err = fmt.Errorf("expected numeric value, got %q", s)
		return
	}
	numStr := s[start:nRunes]
	n, err = strconv.Atoi(numStr)
	return
}

func Iota(n int) string {
	return fmt.Sprintf("%d", n)
}

func Pad2(n int) string {
	return fmt.Sprintf("%02d", n)
}

func Pad9(n int) string {
	return fmt.Sprintf("%09d", n)
}

func PadSpace(n int) string {
	return fmt.Sprintf("%2d", n)
}

func Ordinal(n int) string {
	if n%100 >= 11 && n%100 <= 13 {
		return strconv.Itoa(n) + "th"
	}

	switch n % 10 {
	case 1:
		return strconv.Itoa(n) + "st"
	case 2:
		return strconv.Itoa(n) + "nd"
	case 3:
		return strconv.Itoa(n) + "rd"
	default:
		return strconv.Itoa(n) + "th"
	}
}

func RemoveZeroesFromDecimalPortionOfNumber(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) == 0 {
		return "0"
	}
	return s
}
