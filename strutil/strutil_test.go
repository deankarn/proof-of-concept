package strutil

import "testing"

func TestPalindrome(t *testing.T) {

	tests := []struct {
		val      string
		expected bool
	}{
		{
			val:      "",
			expected: true,
		},
		{
			val:      "aba",
			expected: true,
		},
		{
			val:      "racecar",
			expected: true,
		},
		{
			val:      "abc",
			expected: false,
		},
		{
			val:      "☺aa☺",
			expected: true,
		},
		{
			val:      "♥bb♥",
			expected: true,
		},
	}

	for i, tt := range tests {

		got := IsPalindrome(tt.val)
		if got != tt.expected {
			t.Fatalf("idx: %d Expected '%t' Got '%t'", i, tt.expected, got)
		}
	}
}

func TestAnagram(t *testing.T) {

	tests := []struct {
		val1     string
		val2     string
		expected bool
	}{
		{
			val1:     "",
			val2:     "",
			expected: false,
		},
		{
			val1:     "abc",
			val2:     "abcd",
			expected: false,
		},
		{
			val1:     "Race car",
			val2:     "Car race",
			expected: true,
		},
		{
			val1:     "Angel",
			val2:     "glean",
			expected: true,
		},
		{
			val1:     "Madam Curie",
			val2:     "Radium came",
			expected: true,
		},
		{
			val1:     "Eleven plus two",
			val2:     "Twelve plus one",
			expected: true,
		},
		{
			val1:     "A gentleman",
			val2:     "Elegant man",
			expected: true,
		},
		{
			val1:     "♥☺♥☺♥☺♥",
			val2:     "♥♥♥♥☺☺☺",
			expected: true,
		},
	}

	for i, tt := range tests {

		got := IsAnagram(tt.val1, tt.val2)
		if got != tt.expected {
			t.Fatalf("idx: %d Expected '%t' Got '%t'", i, tt.expected, got)
		}
	}
}
