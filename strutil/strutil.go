package strutil

import (
	"math"
	"strings"
)

// IsPalindrome ...
func IsPalindrome(s string) bool {

	runes := []rune(s)
	j := len(runes) - 1

	for i := 0; i < len(runes)/2; i++ {

		if runes[i] != runes[j] {
			return false
		}
		j--
	}

	return true
}

// IsAnagram ...
func IsAnagram(subject string, s string) bool {

	if len(subject) != len(s) || subject == s {
		return false
	}

	subject = strings.ToLower(subject)
	s = strings.ToLower(s)

	m := make(map[rune]uint64, int(math.Min(float64(len(subject)), 26))) // 26 letters in alphabet

	for _, r := range subject {

		val, found := m[r]
		if found {
			val++
			m[r] = val
			continue
		}

		m[r] = 1
	}

	for _, r := range s {

		val, found := m[r]
		if !found {
			return false
		}

		if val == 1 {
			delete(m, r)
			continue
		}

		val--
		m[r] = val
	}

	return len(m) == 0
}
