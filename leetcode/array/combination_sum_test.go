package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
combinationSum
Time complexity:
給 combinationSum 問題，最糟糕的情況可能需要探索所有可能的組合來尋找答案。
假設 candidates 的最小值是 min，那麼最大的深度（或最多的組合元素）可以是 target 除以 min。
每一層都可能需要探索 n 個候選數字，其中 n 是 candidates 的大小。所以，最壞的時間複雜度大約是 O(n^k)，其中 k 是組合的最大大小，即 target/min。

Space complexity: O(n)
主要的空間開銷是由於遞迴，它的深度可能達到 k，因此空間複雜度是 O(k)。
還有一些其他的空間使用，例如存儲結果，但主要的開銷是由於遞迴。

時間複雜度是 O(n^k)，空間複雜度是 O(k)。

解題思路:
回朔算法，要思考如何去除重複跟回朔算法的邏輯
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	backtrack(
		candidates,
		target,
		0, /* start */
		path,
		&res,
	)
	return res
}

/*
backtrack
candidates: 所有選項
target: 目標
start: 可以幫助我們確保每個元素只被考慮一次，並避免生成重複的組合。
path:
 1. 跟踪當前的選擇。
 2. 幫助我們理解應該添加哪些選擇以及何時需要回溯。
 3. 存儲當前可能的解決方案的部分信息。

res: 最終答案, 為何用ptr, 因為這樣就不用一直創建slice, 就是用同一個
*/
func backtrack(
	candidates []int,
	target int,
	start int,
	path []int,
	res *[][]int,
) {
	// 小於0, 代表沒得談了, 直接return掉，不用再往下探了
	if target < 0 {
		return
	}
	// 代表找到答案了
	if target == 0 {
		/*
			為什麼不直接path呢?
			因為直接用path的話，他是一個記憶體位置，若有地方修正path(list) element 就會影響到res,
		*/
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		// 由於數字可以重複使用，所以還是從當前的i開始
		backtrack(
			candidates,
			target-candidates[i], /* target */
			i,                    /* start */
			path,
			res,
		)
		path = path[:len(path)-1] // 回溯，移除最後一個數字
	}
}

func TestCombinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
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
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			expected: expected{
				result: [][]int{
					{2, 2, 3},
					{7},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				combinationSum(tc.args.candidates, tc.args.target),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
