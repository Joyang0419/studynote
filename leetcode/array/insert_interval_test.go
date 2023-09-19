package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Insert
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
還要記得題目是排序好的
第一個for loop
intervals end < newInterval start -> 加入所有開始時間小於新間隔的間隔

第二個for loop
intervals[i]start  <= newInterval end -> 合併重疊到的變成新的new interval

第三個for loop
加入所有剩下的interval，就是答案了
*/

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int

	i := 0
	// 加入所有時間的結束時間小於新間隔的開始時間
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合併所有與新間隔重疊的間隔
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// 加入所有開始時間大於新間隔的間隔
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestInsert(t *testing.T) {
	type args struct {
		interval    [][]int
		newInterval []int
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
				interval:    [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			expected: expected{
				result: [][]int{{1, 5}, {6, 9}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				Insert(tc.args.interval, tc.args.newInterval),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
