package array

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 自定义排序函数，根据区间的结束时间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	removed := 0           // 初始化移除区间的数量
	end := intervals[0][1] // 设置第一个区间的结束时间

	// 从第二个区间开始遍历每个区间
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的开始时间小于前一个区间的结束时间，说明存在重叠，移除一个区间
		if intervals[i][0] < end {
			removed++
		} else {
			// 如果没有重叠，更新结束时间
			end = intervals[i][1]
		}
	}

	return removed
}

func TestEraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{1, 3},
				},
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
				eraseOverlapIntervals(tc.args.intervals),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
