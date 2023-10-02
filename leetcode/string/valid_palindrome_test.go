package string

import (
	"testing"
	"unicode"
)

/*
 */
func isPalindrome(s string) bool {
	s = cleanedString(s)
	return s == reverse(s)
}

// cleanedString 只保留文字跟數字
func cleanedString(s string) string {
	var cleaned []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	return string(cleaned)
}

// reverse 反轉一個字符串並返回它
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestIsPalindrome(t *testing.T) {
	isPalindrome("ssss")
}
