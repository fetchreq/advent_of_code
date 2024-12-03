/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "2", false)
		//fmt.Printf("Day 2 Part 1 = %d\n", day2Part1(input))
		fmt.Printf("Day 2 Part 2 = %d\n", day2Part2(input)) // 686 wrong
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

func processReport(report []int) bool {
	fmt.Printf("looking at report %v (", report)
	direction := "de"
	if report[0] < report[1] {
		direction = "in"
	}
	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		if curr-next > 0 && direction == "in" {
			fmt.Printf("unsafe)\n")
			return false
		} else if curr-next < 0 && direction == "de" {
			fmt.Printf("unsafe)\n")
			return false
		}
		diff := int(math.Abs(float64(curr - next)))

		if diff < 1 || diff > 3 {
			fmt.Printf("unsafe)\n")
			return false
		}
	}

	fmt.Printf("safe)\n")
	return true
}

func processReportWithError(report []int) bool {
	safe := processReport(report)
	if !safe {
		for i := range len(report) {
			fmt.Printf("Remove %d \n", report[i])
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
