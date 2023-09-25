package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
findMaxLength
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
1. 初始化 maxLength 为 0, cxount 为 0，并创建一个 countIndexMap 哈希表来保存每个 count 值的最早索引。
2. 遍历数组 nums。如果遇到 0，则将 count 减 1；如果遇到 1，则将 count 加 1。
3. 查看 countIndexMap 中是否已经存在当前的 count 值。如果存在，则使用当前索引和之前保存的索引来更新 maxLength。如果不存在，则将当前的 count 值和索引添加到 countIndexMap 中。
4. 返回 maxLength 作为结果。
*/
func findMaxLength(nums []int) int {
	maxLength := 0
	count := 0
	countIndexMap := make(map[int]int)

	// 處理整个数组都是解的情况，用-1, 是你要算長度時，用[0, 1]去推導答案，你會有解答
	countIndexMap[0] = -1

	for i, num := range nums {
		if num == 0 {
			count -= 1
		} else {
			count += 1
		}

		// 如果map存在，就代表已經可以成對了，就會去算成對的長度，是否大於最大長度，就會取代
		if prevIndex, exists := countIndexMap[count]; exists {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			countIndexMap[count] = i
		}
	}

	return maxLength
}

func TestFindMaxLength(t *testing.T) {
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
				nums: []int{0, 1, 1, 0},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 0, 1, 1, 1, 0},
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 1, 1},
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
				findMaxLength(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
