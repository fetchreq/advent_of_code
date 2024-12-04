/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Aoc Day 3 2024",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "2", false)
		fmt.Printf("Day 2 Part 1 = %d\n", day2Part1(input))
		fmt.Printf("Day 2 Part 2 = %d\n", day2Part2(input))
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day2Cmd)
}
func day2Part1(input string) int {
	reports := []bool{}
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		report := []int{}
		for _, val := range numbers {
			report = append(report, cast.ToInt(val))
		}
		reports = append(reports, processReport(report))
	}
	safe := 0
	for _, val := range reports {
		if val {
			safe++
		}
	}
	return safe
}

func day2Part2(input string) int {
	reports := []bool{}
	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		report := []int{}
		for _, val := range numbers {
			report = append(report, cast.ToInt(val))
		}
		reports = append(reports, processReportWithError(report))
	}
	safe := 0
	for _, val := range reports {
		if val {
			safe++
		}
	}
	return safe
}

type dir int

const (
	INC dir = iota
	DEC
)

func processReport(report []int) bool {
	direction := DEC
	if report[0] < report[1] {
		direction = INC
	}
	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		diff := curr - next
		if diff > 0 && direction == INC {
			return false
		} else if diff < 0 && direction == DEC {
			return false
		}

		absDiff := util.AbsDiffInt(curr, next)

		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}

func processReportWithError(report []int) bool {
	safe := processReport(report)
	if !safe {
		for i := range len(report) {
			ret := make([]int, 0)
			ret = append(ret, report[:i]...)
			ret = append(ret, report[util.Min(i+1, len(report)):]...)
			if processReport(ret) {
				return true
			}
		}
		return false
	}
	return true
}
