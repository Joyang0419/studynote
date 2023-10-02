package string

/*
tc:
初始化 targetChars 的哈希表需要 O(p) 的時間，其中 p 是字母異位詞字符串的長度。
主循環遍歷整個字符串 s，其時間複雜度是 O(s)，其中 s 是输入字符串的長度。在循环中，哈希表 windowChars 的更新和比較是 O(1) 的操作，因為哈希表的操作通常被認為是常數時間的，並且目標字母異位詞的長度是固定的。
因此，總時間複雜度是 O(s + p)。

sc:
targetChars 和 windowChars 哈希表的空間複雜度是 O(p)，
因為它们存储了 p 中的每個字符及其计数。
结果数组 result 的空间复杂度是 O(s)，因为在最坏的情况下，
可能每个窗口都是字母异位词，所以需要存储 s 中每个位置的索引。
因此，总空间复杂度是 O(s + p)。

desc:
尋找所有字母異位詞可以通過滑動窗口和哈希表來完成。
在這個方法中，我們使用兩個哈希表，
一個來存儲目標字母異位詞的字符計數，
另一個來存儲當前窗口的字符計數。
然後，我們移動窗口並在每個步驟中更新窗口哈希表，並檢查它是否與目標哈希表匹配。
*/

func findAnagrams(s string, p string) []int {
	var result []int
	if len(s) < len(p) {
		return result
	}

	pCharCount := make(map[rune]int)
	windowCharCount := make(map[rune]int)

	for _, char := range p {
		pCharCount[char]++
	}

	for i, char := range s {
		windowCharCount[char]++

		// Remove the leftmost character from the window
		if i >= len(p) {
			leftmostChar := rune(s[i-len(p)])
			windowCharCount[leftmostChar]--
			if windowCharCount[leftmostChar] == 0 {
				delete(windowCharCount, leftmostChar)
			}
		}

		// If the character count matches, we have found an anagram
		if i >= len(p)-1 && isAnagramMapping(pCharCount, windowCharCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func isAnagramMapping(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
