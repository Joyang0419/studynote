package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
majorityElement
Time complexity: O(n)
Space complexity: O(1)

解題思路:
Boyer–Moore majority vote algorithm(摩爾投票算法)
這個算法的核心在於，
刪去一個數列中的兩個不同的數字，不會影響該數列的majority element。
假想有一群人要投票，候選人有A、B、C，假設A已知會過半數的話，
任取其中2個人取消他們的投票資格，會有以下的狀況：

被取消資格的是B跟C -> 顯然A還是過半數(而且比例更高了XD)
被取消資格的是(A, B)或(A, C) -> 一個投A的和一個不投A的同步被排除，
所以無法改變A會過半數的狀況。
同理，在不只3個候選人(數字)的時候，每次取2個人取消投票資格(移除)，亦無法改變投票結果(A還是會是majority element)。
ref: https://ithelp.ithome.com.tw/articles/10213285
*/
func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
			continue
		}
		count--
	}

	//return candidate
	// 下面這一段未必要, 可以直接 return candidate
	// Optional: check if the candidate is actually the majority element
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}

	// 這是一個好習慣 This shouldn't happen if we know there's a majority element in the list.
	return -1
}

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums     []int
		expected []int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 3},
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			majorityElement(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
