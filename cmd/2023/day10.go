/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day10Cmd represents the day08 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day10 called")
		input := util.ReadFile("2023", "10", false)
		fmt.Printf("Part 1: %d\n", day10Part1(input))
		fmt.Printf("Part 2: %d\n", day10Part2(input))
	},
}
func init() {
	cmd.TwentyThreeCmd.AddCommand(day10Cmd)
}

type pipe int

const (
	Vertical pipe = iota + 1
	Horizontal
	NorthToEast	
	NorthToWest	
	SouthToEast
	SouthToWest
	Ground
)

type direction int
const (
	North direction = iota + 1
	East
	South
	West 
)

// String - Creating common behavior - give the type a String function
func (p pipe) String() string {
	return [...]string{"Vertical", "Horizontal", "NorthToEast", "NorthToWest", "SouthToEast", "SouthToWest", "Ground"}[p-1]
}

// String - Creating common behavior - give the type a String function
func (d direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d-1]
}


func (p pipe) Next(curr Pair) Pair {
	var next Pair
	if p == Horizontal {
		if curr.dir == East {
			next = Pair{x: curr.x, y: curr.y + 1, dir: curr.dir}
		} else if curr.dir == West {
			next = Pair{x: curr.x, y: curr.y - 1, dir: curr.dir}
		} else {
				panic(fmt.Sprintf("In horizontal going north/south (%d)", curr.dir))
		}
	} else if p == Vertical {
		if curr.dir == North {
			next = Pair{x: curr.x - 1, y: curr.y, dir: curr.dir}
		} else if curr.dir == South {
			next = Pair{x: curr.x + 1, y: curr.y, dir: curr.dir}
		} else {
			panic(fmt.Sprintf("In vertical going east/west (%d)", curr.dir))
		}

	} else if p == NorthToEast {

		if curr.dir == South {
			next = Pair{x: curr.x, y: curr.y + 1, dir: East}
		} else if curr.dir == West {
			next = Pair{x: curr.x - 1, y: curr.y, dir: North}
		} else {
			panic(fmt.Sprintf("In North to east with dir (%d)", curr.dir))
		}

	} else if p == NorthToWest {
		if curr.dir == South {
			next = Pair{x: curr.x, y: curr.y - 1, dir: West}
		} else if curr.dir == East {
			next = Pair{x: curr.x -1, y: curr.y, dir: North}
		} else {
			panic(fmt.Sprintf("In North to west with dir (%d)", curr.dir))
		}
	} else if p == SouthToEast {
		if curr.dir == North {
			next = Pair{x: curr.x, y: curr.y + 1, dir: East}
		} else if curr.dir == West {
			next = Pair{x: curr.x + 1, y: curr.y, dir: South}
		} else {
			panic(fmt.Sprintf("In Sourth to east with dir (%d)", curr.dir))
		}
	} else if p == SouthToWest {
		if curr.dir == North {
			next = Pair{x: curr.x, y: curr.y -1, dir: West}
		} else if curr.dir == East {
			next = Pair{x: curr.x + 1, y: curr.y, dir: South}
		} else {
			panic(fmt.Sprintf("In South to west with dir (%d)", curr.dir))
		}
	} else if p == Ground {
		panic(fmt.Sprintf("Reached Ground at (%d, %d)", curr.x, curr.y))
	}

	return next
}

type Pair struct {
    x, y int
    dir direction

}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (p pipe) EnumIndex() int {
	return int(p)
}
func day10Part1(input string) int {
	grid := [][]pipe{}
	var start Pair
	for i, row := range strings.Split(input, "\n") {
		nextRow := []pipe{}
		grid = append(grid, nextRow)
		for j, item := range strings.Split(row, "") {
			if item == "S" {
				start = Pair{x: i, y: j, dir: South}
				grid[i] = append(grid[i], pipe(NorthToWest))
			}

			if item == "." {
				grid[i] = append(grid[i], pipe(Ground))
			} else if item == "|" {
				grid[i] = append(grid[i], pipe(Vertical))

			} else if item == "-" {
				grid[i] = append(grid[i], pipe(Horizontal))

			} else if item == "L" {
				grid[i] = append(grid[i], pipe(NorthToEast))

			} else if item == "J" {
				grid[i] = append(grid[i], pipe(NorthToWest))

			} else if item == "7" {
				grid[i] = append(grid[i], pipe(SouthToWest))

			} else if item == "F" {
				grid[i] = append(grid[i], pipe(SouthToEast))

			}

		}
	}


	foundLoop := false
	curr := start
	steps := 0

	for !foundLoop {
		if steps != 0 && curr.x == start.x && curr.y == start.y {
			break
		}

		currPipe := grid[curr.x][curr.y]
		fmt.Printf("At %s (%d, %d) going %s\n", currPipe.String(), curr.x, curr.y, curr.dir.String())
		curr = currPipe.Next(curr)
		steps++
	}


	return steps / 2;
}


func day10Part2(input string) int {
	return 0
}

