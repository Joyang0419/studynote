package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
longestConsecutive
Time complexity: O(n)
Space complexity:O(n)

解題思路:
用一本map去處理，他會什麼一開始就先-1, 其實滿聰明的，不然你每一個數字都只累加，時間複雜度會很高
*/
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet {
		if !numSet[num-1] { // Check if the element is the start of a sequence
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
func TestLongestConsecutive(t *testing.T) {
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			expected: expected{
				result: 4,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				longestConsecutive(tc.args.nums),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
