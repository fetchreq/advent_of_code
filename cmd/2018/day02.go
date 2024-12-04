/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package eightteen

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day01 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "AoC 2018 Day 01",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2018", "2", false)
		fmt.Printf("Part 1: %d\n", day02Part1(input))
		fmt.Printf("Part 2: %s\n", day02Part2(input))
	},
}

func init() {
	cmd.EightTeenCmd.AddCommand(day02Cmd)
}

func day02Part1(input string) int {
	withTwo := 0
	withThree := 0
	for _, line := range strings.Split(input, "\n") {
		foundTwo := false
		foundThree := false
		for _, letter := range strings.Split(line, "") {

			count := strings.Count(line, letter)
			if count == 2 && !foundTwo {
				withTwo += 1
				foundTwo = true
			} else if count == 3 && !foundThree {
				withThree += 1
				foundThree = true
			}

		}

	}
	return withThree * withTwo
}

func day02Part2(input string) string {
	retVal := ""
	lines := strings.Split(input, "\n")
	found := false
	str1 := ""
	str2 := ""
	for i := 0; i < len(lines)-1 && !found; i += 1 {
		letters := strings.Split(lines[i], "")
		for j := i + 1; j < len(lines) && !found; j += 1 {
			missedCount := 0
			for idx, letter := range letters {
				if string(lines[j][idx]) != letter {
					missedCount += 1
				}

				if missedCount >= 2 {
					break
				}
			}

			if missedCount == 1 {
				found = true
				str1 = lines[i]
				str2 = lines[j]

			}

		}
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			retVal += string(str1[i])
		}
	}

	return retVal
}
