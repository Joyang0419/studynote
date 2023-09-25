package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
canCompleteCircuit
Time complexity: O(n)

Space complexity:O(1), 因為全部都是常數

解題思路:
最後是確認totalGas > totalCost
每一站確認tank可以到下一站，若不行，就把start i + 1, 就代表換起始點, 跟把tank歸0
*/
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	tank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		/*
			確認可以到下一站嗎
		*/
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

func TestCanCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
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
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert.Equal(
				t,
				tc.expected.result,
				canCompleteCircuit(tc.args.gas, tc.args.cost),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
