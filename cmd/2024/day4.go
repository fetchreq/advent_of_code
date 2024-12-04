/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Aoc Day 4 2024",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "4", false)
		fmt.Printf("Day 4 Part 1: %d\n", day4Part1(input))
		fmt.Printf("Day 4 Part 2: %d\n", day4Part2(input))
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day4Cmd)
}

func day4Part1(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			if lines[i][j] == 'X' {
				if searchLineBackwards(lines[i], j) {
					count++
				}
				if searchLineForeward(lines[i], j) {
					count++
				}
				if searchColUpwards(lines, i, j) {
					count++
				}
				if searchColDownwards(lines, i, j) {
					count++
				}
				if searchDiagonallyUpRight(lines, i, j) {
					count++
				}
				if searchDiagonallyUpLeft(lines, i, j) {
					count++
				}
				if searchDiagonallyDownRight(lines, i, j) {
					count++
				}
				if searchDiagonallyDownLeft(lines, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func searchLineBackwards(line string, j int) bool {
	return j >= 3 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S'
}

func searchLineForeward(line string, j int) bool {
	return j < len(line)-3 && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S'
}

func searchColUpwards(lines []string, i, j int) bool {
	return i >= 3 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S'
}

func searchColDownwards(lines []string, i, j int) bool {
	return i < len(lines)-3 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S'
}

func searchDiagonallyUpRight(lines []string, i, j int) bool {
	return j >= 3 && i >= 3 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S'
}

func searchDiagonallyUpLeft(lines []string, i, j int) bool {
	return j >= 3 && i < len(lines)-3 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S'
}

func searchDiagonallyDownRight(lines []string, i, j int) bool {
	return j < len(lines[i])-3 && i >= 3 && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S'
}

func searchDiagonallyDownLeft(lines []string, i, j int) bool {
	return j < len(lines[i])-3 && i < len(lines)-3 && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S'
}

func day4Part2(input string) int {
	lines := strings.Split(input, "\n")
	// Stores the location of a the 'A' in MAS
	queue := make(map[string]struct{})
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			if lines[i][j] == 'M' {
				// search diagonally up to the right

				if j >= 2 && i >= 2 {
					index := fmt.Sprintf("(%d,%d)", i-1, j-1)
					if lines[i-1][j-1] == 'A' && lines[i-2][j-2] == 'S' {
						// Check if we've seen the 'A' before
						if _, ok := queue[index]; !ok {
							// If not it's a new MAS that we might need later
							queue[index] = struct{}{}
						} else {
							// if we have found this 'A' before we have a valid X-MAS
							delete(queue, index)
							count++
						}
					}
				}

				// search diagonally up to the left
				if j >= 2 && i < len(lines)-2 {
					index := fmt.Sprintf("(%d,%d)", i+1, j-1)
					if lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'S' {
						if _, ok := queue[index]; !ok {
							queue[index] = struct{}{}
						} else {
							delete(queue, index)
							count++
						}
					}
				}
				// search diagonally down to the right
				if j < len(lines[i])-2 && i >= 2 {
					index := fmt.Sprintf("(%d,%d)", i-1, j+1)
					if lines[i-1][j+1] == 'A' && lines[i-2][j+2] == 'S' {
						if _, ok := queue[index]; !ok {
							queue[index] = struct{}{}
						} else {
							delete(queue, index)
							count++
						}
					}
				}
				// search diagonally down to the left
				if j < len(lines[i])-2 && i < len(lines)-2 {
					index := fmt.Sprintf("(%d,%d)", i+1, j+1)
					if lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'S' {
						if _, ok := queue[index]; !ok {
							queue[index] = struct{}{}
						} else {
							delete(queue, index)
							count++
						}
					}
				}
			}
		}
	}
	return count
}
