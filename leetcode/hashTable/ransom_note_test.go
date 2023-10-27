package hashTable

/*
tc:
初始化 letters 陣列的時間複雜度是 O(1)，因為其大小是固定的 26。
第一個 for 循環遍歷 magazine 字符串，其時間複雜度是 O(m)，其中 m 是 magazine 的長度。
第二個 for 循環遍歷 ransomNote 字符串，其時間複雜度是 O(n)，其中 n 是 ransomNote 的長度。

sc:
letters 陣列的大小是固定的，所以空間複雜度是 O(1)。
除了 letters 陣列外，沒有使用任何其他額外的數據結構。
總空間複雜度: O(1)

desc:
使用了一個大小為 26 的整數陣列來跟踪每個字母的計數。
ch-'a', 這裡就是細節
*/
func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26) // 創建一個陣列來存放26個字母的計數

	// 遍歷magazine，計算每個字符的出現次數
	for _, ch := range magazine {
		letters[ch-'a']++
	}

	// 遍歷ransomNote，對每個字符的計數進行減法
	for _, ch := range ransomNote {
		letters[ch-'a']--
		if letters[ch-'a'] < 0 {
			return false // 如果計數為負，表示ransomNote中的字符比magazine中的多
		}
	}
	return true
}
