/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "3", false)
		fmt.Printf("Day 3 Part 1 %d\n", day3Part1(input))
		fmt.Printf("Day 3 Part 2 %d\n", day3Part2(input))
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day3Cmd)
}

func day3Part1(input string) int {
	sum := 0
	for i := range len(input) {
		curr := input[i]
		if curr == 'm' && input[i+1] == 'u' && input[i+2] == 'l' && input[i+3] == '(' {
			start := i + 4
			end := start
			for isDigit(input[end]) {
				end++
			}

			if input[end] != ',' {
				continue
			}

			strOne := input[start:end]

			start = end + 1
			end = start
			for isDigit(input[end]) {
				end++
			}
			strTwo := input[start:end]
			if len(strOne) > 3 || len(strTwo) > 3 {
				continue
			}

			if input[end] == ')' {
				numOne := cast.ToInt(strOne)
				numTwo := cast.ToInt(strTwo)
				sum += numOne * numTwo

			}
		}
	}
	return sum
}

func day3Part2(input string) int {
	enabled := true
	sum := 0
	for i := range len(input) {
		curr := input[i]

		if curr == 'd' && input[i+1] == 'o' {
			if input[i+2] == 'n' && input[i+3] == '\'' && input[i+4] == 't' {
				enabled = false
			} else {
				enabled = true
			}
		}
		if curr == 'm' && input[i+1] == 'u' && input[i+2] == 'l' && input[i+3] == '(' {
			start := i + 4
			end := start
			for isDigit(input[end]) {
				end++
			}

			if input[end] != ',' {
				continue
			}

			strOne := input[start:end]

			start = end + 1
			end = start
			for isDigit(input[end]) {
				end++
			}
			strTwo := input[start:end]
			if len(strOne) > 3 || len(strTwo) > 3 {
				continue
			}

			if input[end] == ')' {
				if enabled {
					numOne := cast.ToInt(strOne)
					numTwo := cast.ToInt(strTwo)
					sum += numOne * numTwo

				}

			}
		}
	}
	return sum
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
