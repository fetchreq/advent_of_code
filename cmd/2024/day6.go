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

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "6", false)
		fmt.Println("Day 6 Part 1 = ", day6Part1(input))
		fmt.Println("Day 6 Part 2 = ", day6Part2(input))
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day6Cmd)
}

type GridType int

const (
	OBSTRUCTIONS GridType = iota
	PATH
	OPEN
)

func buildPuzzle(input string) ([][]GridType, Coordinates) {
	updates := make([][]GridType, 0)
	var starting Coordinates

	for i := range updates {
		updates[i] = make([]GridType, 0)
	}

	for i, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, "")
		update := make([]GridType, 0, len(splits))
		for j, cell := range splits {
			if cell == "." {
				update = append(update, OPEN)
			} else if cell == "#" {
				update = append(update, OBSTRUCTIONS)
			} else if cell == "^" {
				update = append(update, PATH)
				starting = Coordinates{x: i, y: j, dir: util.NORTH}
			} else if cell == ">" {
				update = append(update, PATH)
				starting = Coordinates{x: i, y: j, dir: util.EAST}
			} else if cell == "v" {
				update = append(update, PATH)
				starting = Coordinates{x: i, y: j, dir: util.SOUTH}
			} else if cell == "<" {
				update = append(update, PATH)
				starting = Coordinates{x: i, y: j, dir: util.WEST}
			}
		}
		updates = append(updates, update)
	}
	return updates, starting
}

func day6Part1(input string) int {
	puzzle, starting := buildPuzzle(input)
	path := findPath(starting.x, starting.y, starting.dir, puzzle)
	return len(path)
}

func findPath(startingX, startingY int, direction util.CardnialDirection, puzzle [][]GridType) map[Coordinates]struct{} {
	path := map[Coordinates]struct{}{}

	dir := direction
	x := startingX
	y := startingY
	for {
		switch dir {
		case util.NORTH:
			if puzzle[x-1][y] == OBSTRUCTIONS {
				dir = util.EAST
			} else {
				path[Coordinates{x: x, y: y}] = struct{}{}
				x--
			}
		case util.EAST:
			if puzzle[x][y+1] == OBSTRUCTIONS {
				dir = util.SOUTH
			} else {
				path[Coordinates{x: x, y: y}] = struct{}{}
				y++
			}
		case util.SOUTH:
			if puzzle[x+1][y] == OBSTRUCTIONS {
				dir = util.WEST
			} else {
				path[Coordinates{x: x, y: y}] = struct{}{}
				x++
			}
		case util.WEST:
			if puzzle[x][y-1] == OBSTRUCTIONS {
				dir = util.NORTH
			} else {
				path[Coordinates{x: x, y: y}] = struct{}{}
				y--
			}
		}
		if x+1 == len(puzzle) || x == 0 || y+1 == len(puzzle[0]) || y == 0 {
			path[Coordinates{x: x, y: y}] = struct{}{}
			break
		}
	}

	return path

}

func printPuzzle(puzzle [][]GridType) {
	count := 1
	for _, line := range puzzle {
		for _, cell := range line {
			if cell == PATH {
				fmt.Printf("%d ", count)
			} else if cell == OBSTRUCTIONS {
				fmt.Printf("# ")
			} else if cell == OPEN {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}

}
func hasLoop(puzzle [][]GridType, starting Coordinates) bool {
	path := make(map[Coordinates]struct{})
	dir := starting.dir
	x := starting.x
	y := starting.y
	for {
		switch dir {
		case util.NORTH:
			if puzzle[x-1][y] == OBSTRUCTIONS {
				dir = util.EAST
			} else {
				path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}
				x--
			}
		case util.EAST:
			if puzzle[x][y+1] == OBSTRUCTIONS {
				dir = util.SOUTH
			} else {
				path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}
				y++
			}
		case util.SOUTH:
			if puzzle[x+1][y] == OBSTRUCTIONS {
				dir = util.WEST
			} else {
				path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}
				x++
			}
		case util.WEST:
			if puzzle[x][y-1] == OBSTRUCTIONS {
				dir = util.NORTH
			} else {
				path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}
				y--
			}
		}

		if _, ok := path[Coordinates{x: x, y: y, dir: dir}]; !ok {
			path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}

		} else {
			return true
		}
		if x+1 == len(puzzle) || x == 0 || y+1 == len(puzzle[0]) || y == 0 {
			path[Coordinates{x: x, y: y, dir: dir}] = struct{}{}
			break
		}

	}
	return false
}
func day6Part2(input string) int {
	puzzle, starting := buildPuzzle(input)
	path := findPath(starting.x, starting.y, starting.dir, puzzle)
	cycles := 0
	for step, _ := range path {
		if step.x == starting.x && step.y == starting.y {
			continue
		}
		puzzle[step.x][step.y] = OBSTRUCTIONS
		if hasLoop(puzzle, starting) {
			cycles++
		}
		puzzle[step.x][step.y] = OPEN
	}
	return cycles
}

type Coordinates struct {
	x   int
	y   int
	dir util.CardnialDirection
}
