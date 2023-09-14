package cmd

import "testing"

func TestSolution(t *testing.T) {
	var tests = []struct {
		input  []int
		output []string
	}{
		{[]int{}, []string{}},     // empty input
		{[]int{1}, []string{"1"}}, // single range
		{[]int{1, 2}, []string{"1..2"}},
		{[]int{1, 2, 3}, []string{"1..3"}},
		{[]int{1, 2, 3, 5}, []string{"1..3", "5"}}, // two ranges
		{[]int{1, 2, 3, 5, 6}, []string{"1..3", "5..6"}},
		{[]int{1, 2, 3, 5, 6, 7}, []string{"1..3", "5..7"}},
		{[]int{1, 2, 3, 5, 6, 7, 10}, []string{"1..3", "5..7", "10"}},           // three ranges
		{[]int{11, 10, 7, 6, 5, 3, 2, 1}, []string{"1..3", "5..7", "10..11"}},   // unsorted input
		{[]int{1, -1, 2, 7, 9, 2, 0, 6}, []string{"-1..2", "6..7", "9"}},        // negative numbers
		{[]int{-2, -2, 0, 1, 1, 3, 4, 5, 5, 5}, []string{"-2", "0..1", "3..5"}}, // duplicate numbers
		{[]int{0, 0, 0}, []string{"0"}},
		{[]int{1, 3, 5, 7, 9}, []string{"1", "3", "5", "7", "9"}}, // no ranges
	}

	for _, test := range tests {
		result := Solution(test.input)
		if len(result) != len(test.output) {
			t.Errorf("Solution(%v) = %v, want %v", test.input, result, test.output)
		}
		for i := range result {
			if result[i] != test.output[i] {
				t.Errorf("Solution(%v) = %v, want %v", test.input, result, test.output)
			}
		}
	}
}
