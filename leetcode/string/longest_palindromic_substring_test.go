package string

import (
	"fmt"
	"testing"
)

/*
tc:
程式碼中有一個主循環，它遍歷了字符串中的所有字符，循環的次數是 n（其中 n 是字符串的長度）。在每次迭代中，它呼叫了 expandAroundCenter 函數兩次（一次針對奇數長度的回文，一次針對偶數長度的回文）。expandAroundCenter 函數的最壞情況下，會試圖擴展到整個字符串的長度，因此它的時間複雜度是 O(n)。
因此，主循環的總時間複雜度是 O(n) * O(n) = O(n^2)。
sc:
程式碼中沒有使用到與輸入大小成比例的額外空間，所有的變數和數據結構的大小都是固定的，不依賴於輸入的大小。例如，maxLength，start，L 和 R 等都是常數空間。
因此，空間複雜度是 O(1)。
desc:
可以使用中心擴展法。該方法的基本思想是從每個可能的中心開始，向兩邊擴展，以找到最長的回文子串。這個方法的時間復雜度為 O(n^2)，其中 n 是字符串的長度。
在這個實現中，longestPalindrome 函數遍歷了字符串 s，並嘗試從每個位置作為中心來找到最長的回文子串。expandAroundCenter 函數負責從指定的中心向兩側擴展，以找到最長的回文子串。max 函數是一個輔助函數，用於找到兩個整數中的最大值。最後，我們返回從 start 開始，長度為 maxLength 的子串作為結果。這個解決方案是相對簡單和直接的，並且容易理解。
*/
func longestSubStringPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 奇數長度的回文
		len2 := expandAroundCenter(s, i, i+1) // 偶數長度的回文
		maxLen := max(len1, len2)
		if maxLen > maxLength {
			maxLength = maxLen
			start = i - (maxLen-1)/2
		}
	}

	return s[start : start+maxLength]
}

// 核心在這邊 想像他從中心擴散出去比對
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLongestSubStringPalindrome(t *testing.T) {
	fmt.Println(longestSubStringPalindrome("babad"))
}
