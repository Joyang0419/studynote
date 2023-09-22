package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
productExceptSelf
Time complexity: O(n)
第一個循環計算左側的乘積，其時間複雜度為 O(n)。
第二個循環計算右側的乘積，也是 O(n)。
第三個循環將左側和右側的乘積組合起來，時間複雜度為 O(n)。
O(3n) -> 通常會忽略大O表示法中的常數，所以空間複雜度為 O(n)。
Space complexity: O(n)
我們創建了三個輔助的陣列：left, right, 和 output。每個陣列的大小都是 n。
O(3n) -> 通常會忽略大O表示法中的常數，所以空間複雜度為 O(n)。
解題思路:
這一題沒有空間複雜度限制
// notice
You must write an algorithm that runs in O(n) time and without using the division operation.

左右乘積這塊事情要想一下
只排除自己的情況下，所以左右指針的第1個位置都直接補1
*/
func productExceptSelf(nums []int) []int {
	length := len(nums)

	// 初始化左側和右側的乘積陣列
	left := make([]int, length)
	right := make([]int, length)

	// output 用於存儲最終結果
	output := make([]int, length)

	// 計算左側的乘積
	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	// 計算右側的乘積
	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	// 組合左側和右側的乘積
	for i := 0; i < length; i++ {
		output[i] = left[i] * right[i]
	}

	return output
}

func TestProductExceptSelf(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
			},
			expected: expected{
				result: []int{24, 12, 8, 6},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				productExceptSelf(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
