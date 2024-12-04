/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day08Cmd represents the day08 command
var day09Cmd = &cobra.Command{
	Use:   "day09",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day09 called")
		input := util.ReadFile("2023", "9", true)
		//fmt.Printf("Part 1: %d\n", day9Part1(input))
		fmt.Printf("Part 2: %d\n", day9Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day09Cmd)
}

func day9Part1(input string) int {
	sum := 0
	for idx, row := range strings.Split(input, "\n") {
		toSolve := []int{}
		for _, field := range strings.Fields(row) {
			toSolve = append(toSolve, cast.ToInt(field))
		}
		next := toSolve[len(toSolve)-1] + solveForNext(toSolve)
		fmt.Printf("For row %d next value is %d\n", idx, next)
		sum += next
	}

	return sum
}

func solveForNext(input []int) int {
	if isAllZeros(input) {
		return 0
	}
	difference := []int{}
	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		difference = append(difference, diff)
	}

	return difference[len(difference)-1] + solveForNext(difference)
}

func solveForPrevious(input []int) int {
	if isAllZeros(input) {
		return 0
	}
	difference := []int{}
	for i := 0; i < len(input)-1; i++ {
		diff := input[i] - input[i+1]
		difference = append(difference, diff)
	}

	fmt.Printf("Adding %d\n", difference[len(difference)-1])

	return difference[len(difference)-1] + solveForNext(difference)
}

func isAllZeros(input []int) bool {
	for _, item := range input {
		if item != 0 {
			return false
		}
	}

	return true
}

func day9Part2(input string) int {
	sum := 0

	for idx, row := range strings.Split(input, "\n") {
		toSolve := []int{}

		for _, field := range strings.Fields(row) {
			toSolve = append(toSolve, cast.ToInt(field))
		}

		for i := len(toSolve)/2 - 1; i >= 0; i-- {
			opp := len(toSolve) - 1 - i
			toSolve[i], toSolve[opp] = toSolve[opp], toSolve[i]
		}
		fmt.Println(toSolve)
		next := toSolve[len(toSolve)-1] - solveForPrevious(toSolve)
		fmt.Printf("For row %d next value is %d\n", idx, next)
		sum += next
	}

	return sum
}
