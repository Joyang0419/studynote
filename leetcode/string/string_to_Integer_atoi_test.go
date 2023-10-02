package string

import (
	"fmt"
	"math"
	"testing"
)

/*
desc:
在這個實現中，
myAtoi 函數首先跳過任何前置空格，
然後檢查正負符號，
然後讀取並轉換所有連續的數字字符，
同時檢查是否發生溢出。最後返回結果。這是一個簡單直接的解決方案，容易理解並實現。
*/
func myAtoi(s string) int {
	i, n, sign := 0, len(s), 1
	// 忽略前置空格
	for i < n && s[i] == ' ' {
		i++
	}
	// 考慮正負符號
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	// 轉換數字，並檢查溢出
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		// 十進位啊, 有下一個數字，代表原本的數字要 * 10
		result = result*10 + int(s[i]-'0')
		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if result*sign < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return result * sign
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("-43"))
}
