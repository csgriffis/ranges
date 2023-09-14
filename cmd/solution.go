package cmd

import (
	"fmt"
	"sort"
)

// Solution returns all number ranges contained in the input slice.
func Solution(input []int) []string {
	var ranges []string

	if len(input) == 0 {
		return ranges
	}

	if len(input) == 1 {
		return []string{fmt.Sprintf("%d", input[0])}
	}

	// sort input to improve iterations needed to find ranges
	// Space: O(n)
	// Time (worst case): O(n (log n)(log n))
	sort.SliceStable(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	rangeToString := func(r []int) string {
		if len(r) == 1 {
			return fmt.Sprintf("%d", r[0])
		}
		return fmt.Sprintf("%d..%d", r[0], r[1])
	}

	var r []int
	for _, n := range input {
		if len(r) == 0 {
			// add range start
			r = append(r, n)
			continue
		}

		if len(r) == 1 {
			// ignore duplicate number
			if n == r[0] {
				continue
			}

			// add range end
			if n == r[0]+1 {
				r = append(r, n)
				continue
			}
		}

		if len(r) == 2 {
			// ignore duplicate number
			if n == r[1] {
				continue
			}

			// extend range end
			if n == r[1]+1 {
				r[1] = n
				continue
			}
		}

		// close range
		ranges = append(ranges, rangeToString(r))

		// reset range
		r = []int{n}
	}

	// handle last range
	ranges = append(ranges, rangeToString(r))

	return ranges
}
