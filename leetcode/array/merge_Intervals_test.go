package array

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
merge
Time complexity: O(n log n)
Space complexity: O(n)

解題思路:
一開始要排序, 用start排，後續操作比較方便
然後宣告一個current := intervals[0]
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 根據 Start 進行排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	// 開始的第一個區間
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// 如果當前區間與下一個區間重疊
		// if current.End >= intervals[i].Start
		if current[1] >= intervals[i][0] {
			if current[1] < intervals[i][1] {
				current[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, current)
			current = intervals[i]
		}
	}
	// 添加最後一個區間
	merged = append(merged, current)

	return merged
}

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			expected: expected{
				result: [][]int{
					{1, 6},
					{8, 10},
					{15, 18},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				merge(tc.args.intervals),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
