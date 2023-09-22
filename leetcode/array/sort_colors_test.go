package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
sortColors
Time complexity: O(n)
Space complexity: O(1)

解題思路:
一個指針 p0 用來追踪0的最右邊的邊界
一個指針 p2 用來追踪2的最左邊的邊界
一個指針 curr 用來遍歷數組

主要的概念是移動, curr, p2當交集時，代表所有元素都已經全部走完了
因為他只有三個數字，所以可以這樣做
*/
func sortColors(nums []int) {
	p0, curr, p2 := 0, 0, len(nums)-1

	// curr <= p2 兩邊碰到，代表所有元素都搞定了
	for curr <= p2 {
		// 當目前指到的數字是0, 跟最左邊換
		if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr] // 換位置
			p0++
			curr++
			// 當目前指到的數字是0, 跟最右邊換
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
			// 1 的就放在原位
		} else {
			curr++
		}
	}
}

func TestSortColors(t *testing.T) {
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
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			expected: expected{
				result: []int{0, 0, 1, 1, 2, 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sortColors(tc.args.nums)
			assert.Equal(
				t,
				tc.expected.result,
				tc.args.nums,
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
