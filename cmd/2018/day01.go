/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package eightteen

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "AoC 2018 Day 01",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2018", "1", false)
		fmt.Printf("Part 1: %d\n", day01Part1(input))
		fmt.Printf("Part 2: %d\n", day01Part2(input))
	},
}

func init() {
	cmd.EightTeenCmd.AddCommand(day01Cmd)
}

func day01Part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		num := cast.ToInt(line[1:])
		if strings.HasPrefix(line, "-") {
			num *= -1
		}
		sum += num

	}
	return sum
}

func day01Part2(input string) int {
	sum := 0
	retVal := 0
	found := make(map[int]bool)
	searching := true

	for searching {
		for _, line := range strings.Split(input, "\n") {
			num := cast.ToInt(line[1:])
			if strings.HasPrefix(line, "-") {
				num *= -1
			}
			sum += num
			if _, ok := found[sum]; ok {
				retVal = sum
				searching = false
				break
			}
			found[sum] = true

		}
	}
	return retVal
}
