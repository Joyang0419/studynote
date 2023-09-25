package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Rotate
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
getRotateIdx:
取餘數，就可以取得rotate的位置: 原本的位置，加上要rotate的位置
*/
func rotate(nums []int, k int) {
	length := len(nums)
	copyNums := make([]int, length)
	copy(copyNums, nums)

	for idx, num := range copyNums {
		nums[getRotateIdx(idx, k, length)] = num
	}
}

func getRotateIdx(current, rotate, length int) int {
	return (current + rotate) % length
}

//func rotate(nums []int, k int) {
//	if len(nums) < k {
//		k = k % len(nums)
//		if k == 0 {
//			return
//		}
//	}
//	arr := make([]int, len(nums))
//	copy(arr, nums)
//	cutPoint := len(arr) - k
//	for i := 0; i < len(arr); i++ {
//		if cutPoint >= len(arr) {
//			cutPoint = 0
//		}
//		nums[i] = arr[cutPoint]
//		cutPoint++
//	}
//}

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expected: expected{
				result: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{-1},
				k:    2,
			},
			expected: expected{
				result: []int{-1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			expected: expected{
				result: []int{2, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.args.nums, tc.args.k)
			assert.Equal(
				t,
				tc.expected.result,
				tc.args.nums,
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
