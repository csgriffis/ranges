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
		// start first range
		if len(r) == 0 {
			r = append(r, n)
			continue
		}

		if len(r) == 1 {
			// extend current range
			if n == r[0]+1 {
				r = append(r, n)
				continue
			} else {
				// close range
				ranges = append(ranges, rangeToString(r))
				// reset range
				r = []int{n}
				continue
			}
		}

		if len(r) == 2 {
			if n == r[1]+1 {
				// extend current range
				r[1] = n
			} else {
				// close range
				ranges = append(ranges, rangeToString(r))
				// reset range
				r = []int{n}
			}
		}
	}

	// handle last range
	ranges = append(ranges, rangeToString(r))

	return ranges
}
