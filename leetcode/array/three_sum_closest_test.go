package array

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func threeSumClosest(nums []int, target int) int {
	// 首先对数组进行排序
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		// 使用双指针技巧来寻找最接近的和
		left, right := i+1, len(nums)-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			// 更新最接近的和
			if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = currentSum
			}

			// 根据当前和与目标的关系，调整指针
			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				threeSumClosest(tc.args.nums, tc.args.target),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
