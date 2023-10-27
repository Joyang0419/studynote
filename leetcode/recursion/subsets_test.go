package recursion

import (
	"fmt"
	"testing"
)

func subsets(nums []int) [][]int {
	var result [][]int
	backtrack(&result, []int{}, nums, 0)
	return result
}

/*
result: 放ptr, 因為為slice 新增新元素才會改到
current: 現在要放入result的subset
nums: 透過改nums 達到全部的元素都考慮到
start: 是在調整nums idx的起始點
*/
func backtrack(result *[][]int, current []int, nums []int, startIdx int) {
	// 將當前子集加入結果中
	// temp 就是subSet要放到答案中
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	for i := startIdx; i < len(nums); i++ {
		// 將 nums[i] 加入到當前子集中
		current = append(current, nums[i])
		// 遞迴呼叫
		backtrack(result, current, nums, i+1)
		// 回溯，移除最後一個元素
		current = current[:len(current)-1]
	}
}

/*
	           []
	       /        \
	      /          \
	    [1]          []
	   /   \        /   \
	 [1,2] [1]    [2]   []
	/   \  / \    / \   / \

[1,2,3][1,2][1,3][1][2,3][2][3][]
*/
func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	ss := subsets(nums)
	for _, subset := range ss {
		fmt.Println(subset)
	}
}
