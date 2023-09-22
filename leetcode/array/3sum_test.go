package array

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
threeSum
Time complexity:
首先對陣列進行排序，這需要 O(n log n) 的時間複雜度
雙指針的技巧找到和為零的三元組，這需要 O(n^2) 的時間複雜度
整體的時間複雜度為 O(n^2)

Space complexity: O(n)
原地排序 nums 不需要額外的空間，所以排序的空間複雜度為 O(1)。
result 答案集合的大小取決於有多少不重複的三元組可以組成和為0。在最壞的情況下（例如，所有的數字都是0），三元組的數量是 O(n^2)。但這種情況很罕見，大多數時候答案集合的大小遠小於 O(n^2)。但從理論分析的角度，我們可以說答案集合的空間複雜度是 O(n^2)。
除了上述的存儲空間，解法中其他的變數（例如 i, l, r 等）使用的空間都是常數，所以他們的空間複雜度是 O(1)。
因此，整體的空間複雜度是 O(n^2 + 1)，
但是當我們談到大 O 符號表示的複雜度時，我們關心的是最高次的項，所以我們可以簡化它為 O(n^2)。

解題思路:
開場先排序, 方便後續操作
每一個避免重複解是因為下面的notice
Notice that the solution set must not contain duplicate triplets.
Notice that the order of the output and the order of the triplets does not matter.

nums[i] > 0 -> break, 是因為[> 0, > 0, > 0] 已經不可能加總 = 0

	if i > 0 && nums[i] == nums[i-1] {
				continue
			}

-> 同樣的idx指出來的數字，會有重複的問題，所以直接continue掉

// 雙指針法
*/
func threeSum(nums []int) [][]int {
	// 開始先排序, 方便後續使用，光第一條當前數字大於0, 就不可能找到0的3元組, 下面我會講
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 如果當前數字大於0，則不可能再找到和為0的三元組
		// 這就是排序的魅力 數字 > 0, 就三個元素[> 0, > 0, > 0], 已經不可能加總 = 0
		// 終止條件
		if nums[i] > 0 {
			break
		}

		// 避免重複解, 當i > 0 了，
		// 他就開始確認，目前idx指到的數字跟上一個數字一樣嗎?, 這一切都是建立在他「排序」過了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 雙指針法
		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(
					result,
					[]int{nums[i], nums[left], nums[right]},
				)
				left++
				right--

				// 避免重複解, left-1, 是原本left上面加進去的數字
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 避免重複解, right+1, 是原本right加進去的數字
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if total < 0 {
				left++
			} else { // total > 0
				right--
			}
		}
	}

	return result
}

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result [][]int
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
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			expected: expected{
				result: [][]int{
					{-1, -1, 2},
					{-1, 0, 1},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				threeSum(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
