/*
Copyright © 2023 Chris Griffis <csgriffis@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var input string

var rootCmd = &cobra.Command{
	Use:   "ranges",
	Short: "ranges returns all ranges contained in a list of integers",
	Run: func(cmd *cobra.Command, args []string) {
		// check if input is empty
		// if empty, print usage
		if input == "" {
			// explicitly ignore error returned by Usage()
			_ = cmd.Usage()
			os.Exit(1)
		}

		values := strings.Split(input, ",")

		// parse input
		// convert to int slice
		input := make([]int, 0, len(values))
		for _, v := range values {
			n, err := strconv.Atoi(v)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error parsing integers: %v\n", err)
				os.Exit(1)
			}

			input = append(input, n)
		}

		// calculate ranges
		ranges := Solution(input)

		// print result
		output := "[\"" + strings.Join(ranges, `", "`) + `"]`
		_, _ = fmt.Fprintf(os.Stdout, "%v\n", output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&input, "input", "", "comma-separated list of integers for which to find ranges")
}
