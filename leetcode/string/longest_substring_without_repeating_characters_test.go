package string

import (
	"fmt"
	"testing"
)

/*
desc:
在以上代码中，
我们使用了一个哈希表 charIndex 来记录每个字符最后出现的位置。
同时，我们使用了两个变量 left 和 right 来表示滑动窗口的左边界和右边界。
我们遍历字符串 s，并在每一步中更新哈希表和滑动窗口的边界，以及最长子串的长度 maxLength。最后，我们返回 maxLength 作为结果。
*/
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int) // 记录字符最后出现的位置
	maxLength := 0                  // 最长子串的长度
	left := 0                       // 滑动窗口的左边界

	for right, char := range s { // right 是滑动窗口的右边界
		if lastSeen, ok := charIndex[char]; ok && lastSeen >= left {
			left = lastSeen + 1 // 如果字符已经存在，移动窗口的左边界
		}
		charIndex[char] = right // 更新字符的最后出现位置
		if right-left+1 > maxLength {
			maxLength = right - left + 1 // 更新最长子串的长度
		}
	}

	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
