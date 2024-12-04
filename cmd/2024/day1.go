/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Aoc Day 1 2024",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "1", false)
		fmt.Printf("Day 1 Part 1 = %d\n", day1Part1(input))
		fmt.Printf("Day 1 Part 2 = %d\n", day1Part2(input))
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day1Cmd)
}

func day1Part1(input string) int {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "   ")
		left = append(left, cast.ToInt(parts[0]))
		right = append(right, cast.ToInt(parts[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += util.AbsDiffInt(left[i], right[i])
	}
	return sum
}

func day1Part2(input string) int {
	items := make(map[int]int)
	left := []int{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "   ")
		left = append(left, cast.ToInt(parts[0]))
		right := cast.ToInt(parts[1])
		if count, ok := items[right]; !ok {
			items[right] = 1
		} else {
			items[right] = count + 1
		}
	}

	sum := 0
	for _, val := range left {
		if occ, ok := items[val]; ok {

			sum += (occ * val)
		}
	}
	return sum
}
