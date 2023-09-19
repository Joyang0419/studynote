package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
sortedSquares
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
雙指針法, left: 頭, right: 尾巴
res是結果答案
for left <= right, <=的原因，那個=是為了把最後一個元素也放進去,
leftSquare > rightSquare相比，比較大依照position放入res list,
相對應的left++ or right ++,
position要--
*/
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 左右指針初始化
	left, right := 0, n-1

	// 填充結果陣列的指針, 從尾巴開始填入數字
	position := n - 1

	// 當左指針小於等於右指針時, 交會了，代表所有元素都搞定了
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		// 將較大的平方值填充到結果陣列，並移動相應的指針
		if leftSquare > rightSquare {
			res[position] = leftSquare
			left++
		} else {
			res[position] = rightSquare
			right--
		}

		// 從結果list尾巴開始塞答案
		position--
	}

	return res
}

func TestSortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result []int
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
				nums: []int{-4, -1, 0, 3, 10},
			},
			expected: expected{
				result: []int{0, 1, 9, 16, 100},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				sortedSquares(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
