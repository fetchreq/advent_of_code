/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"regexp"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "AoC 2015 Day 12",

	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day12", false)
		fmt.Printf("Part 1: %d\n", day12Part1(input))
		fmt.Printf("Part 2: %d\n", day12Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day12Cmd)
}

func day12Part1(input string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		curr := string(input[i])

		if regexp.MustCompile("[0-9]").MatchString(curr) || input[i] == '-' {
			isNegitive := false
			if input[i] == '-' {
				isNegitive = true
				i += 1
			}
			start := i;
			for regexp.MustCompile("[0-9]").MatchString(string(input[i])) {
				i++	
			}
			if isNegitive {
				sum -= cast.ToInt(input[start:i])
			} else {
				sum += cast.ToInt(input[start:i])
			}
		}
	}
	return sum
}

func day12Part2(input string) int {
	sum := 0
	
	for i := 0; i < len(input); i++ {
		curr := string(input[i])
		
		if curr == "{" {
			objectStart := i
			for input[i] != '}' {
				i += 1
			}

			if !regexp.MustCompile("red").MatchString(input[objectStart:i]) {
				i = objectStart	
			}
		}
		
		if regexp.MustCompile("[0-9]").MatchString(curr) || input[i] == '-' {
			isNegitive := false
			if input[i] == '-' {
				isNegitive = true
				i += 1
			}
			start := i;
			for regexp.MustCompile("[0-9]").MatchString(string(input[i])) {
				i++	
			}
			if isNegitive {
				sum -= cast.ToInt(input[start:i])
			} else {
				sum += cast.ToInt(input[start:i])
			} 
		}
	}
	return sum
}
