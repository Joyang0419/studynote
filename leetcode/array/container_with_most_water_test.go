package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
maxArea
Time complexity: O(n)
Space complexity: O(1)

解題思路:
雙指針法
移動指針那邊理解一下
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxA := 0

	for left < right {
		h := min(height[left], height[right])
		w := right - left
		area := h * w
		maxA = max(maxA, area)

		// 移動指針, 比較兩邊的高度，讓他可以用最高的邊去算
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxA
}

func TestMaxArea(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			expected: expected{
				result: 49,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert.Equal(
				t,
				tc.expected.result,
				maxArea(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
