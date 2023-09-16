package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
MaxProfit
Time complexity: O(n)
Space complexity: O(1)

解題思路:
開場先用一個常數minPrice = prices[0]
maxProf := 0

他題目是一個int slice: [7, 1, 5, 3, 6, 4]
所以當遇到更低價格成立時，那就先不用算profit,
因為此時， minPrice 就是 price -> 這就是continue的原因
當迭代到的price不是minPrice, 就去算profit, 如果profit > maxProf -> maxProfit = profit
*/
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
			continue
		}
		if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			expected: expected{
				result: 5,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			MaxProfit(tc.args.prices),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
