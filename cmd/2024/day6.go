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

var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "AoC 2024 Day 6",
	Run: func(cmd *cobra.Command, args []string) {
		defer util.Duration(util.Track("Day 6"))
		input := util.ReadFile("2024", "6", false)
		puzzle, starting := convertToGrid(input)
		path := findPath(puzzle, starting)
		fmt.Println("Day 6 Part 1 = ", day6Part1(path))
		fmt.Println("Day 6 Part 2 = ", day6Part2(path, puzzle, starting))

	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day6Cmd)

}

type GridType int

const (
	OBSTRUCTIONS GridType = iota
	OPEN
)

type Coordinates struct {
	X   int
	Y   int
	Dir util.CardnialDirection
}

// Converts the string input into a grid
// while doing that it also finds the starting point
func convertToGrid(input string) ([][]GridType, Coordinates) {
	updates := make([][]GridType, 0)
	var starting Coordinates

	for i := range updates {
		updates[i] = make([]GridType, 0)
	}

	for i, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, "")
		update := make([]GridType, 0, len(splits))
		for j, cell := range splits {
			switch cell {
			case ".":
				update = append(update, OPEN)
			case "#":
				update = append(update, OBSTRUCTIONS)
			case "^":
				update = append(update, OPEN)
				starting = Coordinates{X: i, Y: j, Dir: util.NORTH}
			case ">":
				update = append(update, OPEN)
				starting = Coordinates{X: i, Y: j, Dir: util.WEST}
			case "v":
				update = append(update, OPEN)
				starting = Coordinates{X: i, Y: j, Dir: util.SOUTH}
			case "<":
				update = append(update, OPEN)
				starting = Coordinates{X: i, Y: j, Dir: util.EAST}
			default:
				error := fmt.Sprintf("Unknown gird '%s'", cell)
				panic(error)
			}
		}
		updates = append(updates, update)
	}
	return updates, starting
}

func day6Part1(path map[Coordinates]struct{}) int {
	defer util.Duration(util.Track("Part 1"))
	return len(path)
}

func findPath(puzzle [][]GridType, starting Coordinates) map[Coordinates]struct{} {
	defer util.Duration(util.Track("Find path"))
	path := map[Coordinates]struct{}{}

	x, y, dir := starting.X, starting.Y, starting.Dir

	for {
		path[Coordinates{X: x, Y: y}] = struct{}{}
		switch dir {
		case util.NORTH:
			if puzzle[x-1][y] == OBSTRUCTIONS {
				dir = util.EAST
			} else {
				x--
			}
		case util.EAST:
			if puzzle[x][y+1] == OBSTRUCTIONS {
				dir = util.SOUTH
			} else {
				y++
			}
		case util.SOUTH:
			if puzzle[x+1][y] == OBSTRUCTIONS {
				dir = util.WEST
			} else {
				x++
			}
		case util.WEST:
			if puzzle[x][y-1] == OBSTRUCTIONS {
				dir = util.NORTH
			} else {
				y--
			}
		}

		if x+1 == len(puzzle) || x == 0 || y+1 == len(puzzle[0]) || y == 0 {
			path[Coordinates{X: x, Y: y}] = struct{}{}
			break
		}
	}

	return path

}

func printPuzzle(puzzle [][]GridType) {
	for _, line := range puzzle {
		for _, cell := range line {
			if cell == OBSTRUCTIONS {
				fmt.Printf("# ")
			} else if cell == OPEN {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}

}
func hasLoop(obstructionX, obstructionY int, puzzle [][]GridType, starting Coordinates) bool {
	path := make(map[Coordinates]struct{})
	x, y, dir := starting.X, starting.Y, starting.Dir
	for {
		switch dir {
		case util.NORTH:
			if puzzle[x-1][y] == OBSTRUCTIONS || (x-1 == obstructionX && y == obstructionY) {
				dir = util.EAST
			} else {
				x--
			}
		case util.EAST:
			if puzzle[x][y+1] == OBSTRUCTIONS || (x == obstructionX && y+1 == obstructionY) {
				dir = util.SOUTH
			} else {
				y++
			}
		case util.SOUTH:
			if puzzle[x+1][y] == OBSTRUCTIONS || (x+1 == obstructionX && y == obstructionY) {
				dir = util.WEST
			} else {
				x++
			}
		case util.WEST:
			if puzzle[x][y-1] == OBSTRUCTIONS || (x == obstructionX && y-1 == obstructionY) {
				dir = util.NORTH
			} else {
				y--
			}
		}
		// Check if we've already been at a spot going in the dir
		if _, ok := path[Coordinates{X: x, Y: y, Dir: dir}]; !ok {
			// Add it to the path if we haven't
			path[Coordinates{X: x, Y: y, Dir: dir}] = struct{}{}
		} else {
			// We've been here before going in the second dir so we're in a loop
			return true
		}

		// Check if we're go off the map
		if x+1 == len(puzzle) || x == 0 || y+1 == len(puzzle[0]) || y == 0 {
			path[Coordinates{X: x, Y: y, Dir: dir}] = struct{}{}
			return false
		}
	}
}
func day6Part2(path map[Coordinates]struct{}, puzzle [][]GridType, starting Coordinates) int {
	defer util.Duration(util.Track("Part 2"))
	cycles := 0
	c := make(chan bool)
	for step := range path {
		if step.X == starting.X && step.Y == starting.Y {
			continue
		}

		go func() {
			hasLoop := hasLoop(step.X, step.Y, puzzle, starting)
			c <- hasLoop
		}()
		// // Make the current step on the path an obstacle
		// // and check the path for a loop
		// puzzle[step.X][step.Y] = OBSTRUCTIONS
		// if hasLoop(puzzle, starting) {
		// 	cycles++
		// }
		// // Reset the puzzle
		// puzzle[step.X][step.Y] = OPEN
	}
	for range len(path) - 1 {
		if <-c {
			cycles++
		}
	}
	return cycles
}
