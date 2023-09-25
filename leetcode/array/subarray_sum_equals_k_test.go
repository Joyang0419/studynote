package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
subarraySum
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
1. 我们初始化计数为 0，和为 0，并创建一个 prefixSumCounts 哈希表来存储每个前缀和以及对应的前缀和出现的次数。
2. 我们遍历数组 nums，更新 sum 为 sum + num。
3. 我们检查 sum - k 的值是否存在于 prefixSumCounts 中。如果存在，我们更新 count 为 count + prefixSumCounts[sum - k]。
4. 我们更新 prefixSumCounts[sum] 的值为 prefixSumCounts[sum] + 1，以记录当前的前缀和。
5. 最后，我们返回 count 作为结果。

prevSum通常用于表示从数组的开始到当前索引的累积和
*/
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	// 哈希表来存储每个前缀和以及对应的前缀和出现的次数。
	prefixSumCounts := make(map[int]int)
	prefixSumCounts[0] = 1 // 初始化为1，处理整个数组都是解的情况

	for _, num := range nums {
		prefixSum += num
		key := prefixSum - k
		if prevCount, exists := prefixSumCounts[key]; exists {
			count += prevCount
		}
		prefixSumCounts[prefixSum]++
	}

	return count
}

func TestSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 4}, // [2,3]
				k:    1,
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				subarraySum(tc.args.nums, tc.args.k),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
