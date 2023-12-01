/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use:   "day18",
	Short: "AoC 2015 Day 18",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "18", false)

		fmt.Printf("Part 1: %d\n", day18Part1(input))
		fmt.Printf("Part 2: %d\n", day18Part2(input))


	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day18Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day18Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day18Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day18Part1(input string) int {
		
	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	for loop := 0; loop < 100; loop++ {
		grid = next(grid)
	}


	var count int
	for _, i := range grid {
		for _, j := range i {
			if j == "#" {
				count++
			}
			//fmt.Printf("%s", j)
		}
		//fmt.Printf("\n")
	}


	return count
}

func day18Part2(input string) int {
		
	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	grid[0][0] = "#"
	grid[0][len(grid) - 1] = "#"
	grid[len(grid) - 1][0] = "#"
	grid[len(grid) - 1][len(grid) - 1] = "#"

	for loop := 0; loop < 100; loop++ {


		grid = next(grid)
		grid[0][0] = "#"
		grid[0][len(grid) - 1] = "#"
		grid[len(grid) - 1][0] = "#"
		grid[len(grid) - 1][len(grid) - 1] = "#"
	}


	var count int
	for _, i := range grid {
		for _, j := range i {
			if j == "#" {
				count++
			}
			//fmt.Printf("%s", j)
		}
		//fmt.Printf("\n")
	}


	return count
}

func next(grid [][]string) [][]string {
	var nextGrid [][]string

	for r, row := range grid {
		nextGrid = append(nextGrid, make([]string, len(grid[0])))
		for c, cell := range row {
			var neighbors int
			for rDiff := -1; rDiff <= 1; rDiff++ {
				for cDiff := -1; cDiff <= 1; cDiff++ {
					if !(rDiff == 0 && cDiff == 0) {
						nextRow := r + rDiff
						nextCol := c + cDiff
						if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) &&
							grid[nextRow][nextCol] == "#" {
							neighbors++
						}
					}
				}
			}
			if cell == "#" && (neighbors == 2 || neighbors == 3) {
				nextGrid[r][c] = "#"
			} else if cell == "." && neighbors == 3 {
				nextGrid[r][c] = "#"
			} else {
				nextGrid[r][c] = "."
			}
		}
	}
	return nextGrid
}
