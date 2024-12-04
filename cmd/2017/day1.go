/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
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
		input := util.ReadFile("2017", "1", false)
		fmt.Printf("Part 1: %d\n", day1Part1(input))
		fmt.Printf("Part 2: %d\n", day1Part2(input))
	},
}

func init() {
	cmd.SeventeenCmd.AddCommand(day1Cmd)
}

func day1Part1(input string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[i+1] {
			sum += cast.ToInt(string(input[i]))
		}
	}

	if input[0] == input[len(input)-1] {
		sum += cast.ToInt(string(input[0]))
	}

	return sum
}

func day1Part2(input string) int {
	sum := 0
	factor := len(input) / 2
	for i := 0; i < len(input)-1; i++ {
		if i+factor > len(input) {
		}
		if input[i] == input[i+1] {
			sum += cast.ToInt(string(input[i]))
		}
	}

	if input[0] == input[len(input)-1] {
		sum += cast.ToInt(string(input[0]))
	}

	return sum
}
