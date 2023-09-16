package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
TwoSum
Time complexity: O(n)
Space complexity: Hash Table（或 Map）解決方案的空間複雜度是O(n), 其中 n是數組的長度。
最壞的情況下，我們可能需要存儲整個數組的所有元素及其索引值，因此空間複雜度為O(n)。

解題思路:
- Hash Table (Map) Solution
map 存放的是 key: nums.value, value: nums的index
for idx, num range nums
把資料放入map, 並且用target - num 得到差額
當差額存在map, 代表答案找到了, return {map.value, idx}
*/

func TwoSum(nums []int, target int) []int {
	// key: nums的value, value: nums的index
	m := make(map[int]int)

	// TC: O(n)
	for i, num := range nums {
		complement := target - num
		// 看complement(target - num), 也就是差額，是否存在map
		if idx, exist := m[complement]; exist {
			return []int{idx, i}
		}
		m[num] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	type args struct {
		nums     []int
		target   int
		expected []int
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
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: expected{
				result: []int{0, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: expected{
				result: []int{1, 2},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			TwoSum(tc.args.nums, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
