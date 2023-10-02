package string

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	srunes := []rune(s)
	trunes := []rune(t)

	sort.Slice(srunes, func(i, j int) bool {
		return srunes[i] < srunes[j]
	})
	sort.Slice(trunes, func(i, j int) bool {
		return trunes[i] < trunes[j]
	})

	return string(srunes) == string(trunes)
}
