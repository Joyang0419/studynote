package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
containsDuplicate
Time complexity: O(n)
Space complexity: O(n)

解題思路:
用一個map去紀錄 for loop slice 放進去, 當key值存在代表重複了
*/
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result bool
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
				nums: []int{1, 2, 3, 1},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			containsDuplicate(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
