package string

import (
	"fmt"
	"testing"
)

/*
tc: O(N)
sc: O(N)

desc:
他的題目是s的字元你可以任意排列，找出最長的回文
善用golang int 相除 如果只有一個 length += count / 2 * 2 // (1 / 2) -> 這邊會變成 -> 0 * 2 = 0
回文就是至少要有2個字, 只能有一個字的當中心點
*/
func longestPalindrome(s string) int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	length := 0
	oddExists := false

	for _, count := range charCount {

		length += count / 2 * 2 // (1 / 2) -> 這邊會變成 -> 0 * 2 = 0
		if count%2 == 1 {
			oddExists = true
		}
	}

	if oddExists {
		length++
	}

	return length
}

func TestLongestPalindrome(t *testing.T) {
	s := "abccccdd"
	result := longestPalindrome(s)
	fmt.Println(result)
}
