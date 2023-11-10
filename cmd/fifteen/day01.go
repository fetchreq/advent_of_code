/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Aoc 2015 Day 1",

	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "1", false);
		fmt.Printf("Part One: %d\n", part1(input));
		fmt.Printf("Part Two: %d\n", part2(input));
	},
}

func init() {

	cmd.FifteenCmd.AddCommand(day1Cmd)
}

func part1(input string) int {

	floor := 0

	for i := range(input) {
		char := input[i]
		if char == '(' {
			floor += 1
		} else if char == ')' {
			floor -= 1
		}
	}
	return floor
}

func part2(input string) int {

	floor := 0

	for i := range(input) {
		char := input[i]
		if char == '(' {
			floor += 1
		} else if char == ')' {
			floor -= 1
		}

		if floor < 0 {
			// Add 1 because go starts at 0 but the floors start at 1 in the question
			return i + 1
		}
	}

	return -1; 
}
