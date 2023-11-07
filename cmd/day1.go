/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day1", false);
		fmt.Printf("Part One: %d\n", part1(input));
		fmt.Printf("Part Two: %d\n", part2(input));
	},
}

func init() {
	fifteenCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
