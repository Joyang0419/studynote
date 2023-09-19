package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
MoveZeroes
Time complexity: O(n)
Space complexity:  O(1)

解題思路:
使用兩個指針。一個指針用於遍歷切片，另一個指針用於指向下一個非 0 數字應該放的位置。
遍歷切片。每當遇到非 0 數字時，將其移動到第二個指針指定的位置，然後增加該指針。
遍歷完畢後，將第二個指針後的所有位置設為 0。
*/
func MoveZeroes(nums []int) {
	// 一根指針紀錄現在非0存放到哪
	pos := 0

	// 將所有非 0 數字移到前面
	for _, num := range nums {
		if num != 0 {
			nums[pos] = num
			pos++
		}
	}

	// 剩下的位置設為 0
	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
}

func TestMoveZeroes(t *testing.T) {
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
				nums: []int{0, 1, 0, 3, 12},
			},
			expected: expected{
				result: []int{1, 3, 12, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		MoveZeroes(tc.args.nums)
		assert.Equal(
			t,
			tc.expected.result,
			tc.args.nums,
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
