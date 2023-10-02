package string

/*
tc:
sc:

desc:
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 在當前字符串中查找最長公共前綴
		for j := 0; j < len(strs[i]) && j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				// 如果字符不匹配，縮短最長公共前綴
				prefix = prefix[:j]
				break
			}
		}
		// 如果最長公共前綴已經縮短為空，或者當前字符串比最長公共前綴還要短，進一步縮短最長公共前綴
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}
