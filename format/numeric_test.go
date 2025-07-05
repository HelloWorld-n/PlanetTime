package format_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/HelloWorld-n/PlanetTime/format"
)

func TestParseNumeric(t *testing.T) {
	tests := []struct {
		input          string
		expectedN      int
		expectedNRunes int
		expectedErr    error
	}{
		{"123", 123, 3, nil},
		{" 123", 123, 4, nil},
		{"123abc", 123, 3, nil},
		{"abc", 0, 0, fmt.Errorf("expected numeric value, got %q", "abc")},
		{"", 0, 0, fmt.Errorf("expected numeric value, got %q", "")},
		{" ", 0, 0, fmt.Errorf("expected numeric value, got %q", " ")},
	}

	for _, test := range tests {
		n, nRunes, err := format.ParseNumeric(test.input)
		if err == nil {
			if n != test.expectedN {
				t.Errorf("ParseNumeric(%q): n expected %d, got %d", test.input, test.expectedN, n)
			}
			if nRunes != test.expectedNRunes {
				t.Errorf("ParseNumeric(%q): nRunes expected %d, got %d", test.input, test.expectedNRunes, nRunes)
			}
		} else {
			if err.Error() != test.expectedErr.Error() {
				t.Errorf("ParseNumeric(%q): err expected %v, got %v", test.input, test.expectedErr, err)
			}
		}
	}
}

func TestIota(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{123, "123"},
	}

	for _, test := range tests {
		result := format.Iota(test.input)
		if result != test.expected {
			t.Errorf("Iota(%d): expected %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestPad2(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{10, "10"},
		{99, "99"},
		{100, "100"},
	}

	for _, test := range tests {
		result := format.Pad2(test.input)
		if result != test.expected {
			t.Errorf("Pad2(%d): expected %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestPad9(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "000000000"},
		{1, "000000001"},
		{123456789, "123456789"},
		{100, "000000100"},
	}

	for _, test := range tests {
		result := format.Pad9(test.input)
		if result != test.expected {
			t.Errorf("Pad9(%d): expected %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestPadSpace(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, " 0"},
		{1, " 1"},
		{10, "10"},
		{99, "99"},
		{100, "100"},
	}

	for _, test := range tests {
		result := format.PadSpace(test.input)
		if result != test.expected {
			t.Errorf("PadSpace(%d): expected %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0th"},
		{1, "1st"},
		{2, "2nd"},
		{3, "3rd"},
		{4, "4th"},
		{11, "11th"},
		{12, "12th"},
		{13, "13th"},
		{21, "21st"},
		{22, "22nd"},
		{23, "23rd"},
		{101, "101st"},
		{111, "111th"},
		{112, "112th"},
		{113, "113th"},
	}

	for _, test := range tests {
		result := format.Ordinal(test.input)
		if result != test.expected {
			t.Errorf("Ordinal(%d): expected %q, got %q", test.input, test.expected, result)
		}
	}
}

func TestRemoveZeroesFromDecimalPortionOfNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"12345000", "12345"},
		{"123000", "123"},
		{"1234567890", "123456789"},
		{"123", "123"},
		{"00000", "0"},
		{"0000", "0"},
	}

	for _, test := range tests {
		result := format.RemoveZeroesFromDecimalPortionOfNumber(test.input)
		if result != test.expected {
			t.Errorf("RemoveZeroesFromDecimalPortionOfNumber(%q): expected %q, got %q", test.input, test.expected, result)
		}
	}
}
func TestParseDecimal(t *testing.T) {
	tests := []struct {
		input        string
		n            int
		wantValue    int
		wantConsumed int
		wantErr      error
	}{
		{"12345", 5, 12345, 5, nil},
		{"12345", 3, 123, 3, nil},
		{"00123", 5, 123, 5, nil},
		{"12abc", 5, 12000, 2, nil},
		{"abc", 3, 0, 0, strconv.ErrSyntax},
		{"", 2, 0, 0, strconv.ErrSyntax},
		{"1", 0, 0, 0, strconv.ErrSyntax},
		{"987654321", 4, 9876, 4, nil},
		{"000", 2, 0, 2, nil},
		{"9a8", 3, 900, 1, nil},
	}

	for _, tt := range tests {
		gotValue, gotConsumed, gotErr := format.ParseDecimal(tt.input, tt.n)
		if gotErr == nil && (gotValue != tt.wantValue || gotConsumed != tt.wantConsumed) {
			t.Errorf("ParseDecimal(%q, %d) = (%d, %d, %v), want (%d, %d, %v)",
				tt.input, tt.n, gotValue, gotConsumed, gotErr, tt.wantValue, tt.wantConsumed, tt.wantErr)
		}
		if (gotErr == nil) != (tt.wantErr == nil) {
			t.Errorf("ParseDecimal(%q, %d) error = %v, want %v", tt.input, tt.n, gotErr, tt.wantErr)
		}
		if gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
			t.Errorf("ParseDecimal(%q, %d) error = %v, want %v", tt.input, tt.n, gotErr, tt.wantErr)
		}
	}
}
