/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "Aoc 2015 Day 10",

	Run: func(cmd *cobra.Command, args []string) {
		input := "1113122113"

		fmt.Printf("Part 1: %d\n", day10Part1(input))
		fmt.Printf("Part 1: %d\n", day10Part2(input))

	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day10Cmd)
}

func day10Part1(input string) int {
	var sb strings.Builder
	for loops := 0; loops < 40; loops++ {
		sb.Reset()
		// output = ""
		for i := 0; i < len(input); i ++ {
			curr := input[i];
			count := 1
			for j := i + 1; j < len(input); j++ {
				if input[j] == curr {
					count += 1
				} else {
					break
				}

			}

			i += count - 1;
			val := fmt.Sprintf("%d%s", count, string(curr))
			sb.WriteString(val)
		}
		//fmt.Printf("%d: %s\n", loops + 1, output)
	
		input = sb.String()
	}
	return len(sb.String())
}

func day10Part2(input string) int {
	var sb strings.Builder
	for loops := 0; loops < 50; loops++ {
		sb.Reset()
		// output = ""
		for i := 0; i < len(input); i ++ {
			curr := input[i];
			count := 1
			for j := i + 1; j < len(input); j++ {
				if input[j] == curr {
					count += 1
				} else {
					break
				}

			}

			i += count - 1;
			val := fmt.Sprintf("%d%s", count, string(curr))
			sb.WriteString(val)
		}
		//fmt.Printf("%d: %s\n", loops + 1, output)
	
		input = sb.String()
	}
	return len(sb.String())
}
