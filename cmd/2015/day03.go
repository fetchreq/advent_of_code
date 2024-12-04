/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Aoc 2015 Day 3",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day3", false)
		fmt.Printf("Part 1: %d\n", day3Part1(input))

		fmt.Printf("Part 2: %d\n", day3Part2(input))

	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type location struct {
	x int
	y int
}

func (l *location) moveNorth() {
	l.y += 1
}
func (l *location) moveSouth() {
	l.y -= 1
}
func (l *location) moveEast() {
	l.x += 1
}
func (l *location) moveWest() {
	l.x -= 1
}

func day3Part1(input string) int {
	moves := strings.Split(input, "")
	locationSet := make(map[location]bool)
	current := location{x: 0, y: 0}

	locationSet[current] = true

	for _, move := range moves {
		if move == ">" {
			current.moveEast()
		} else if move == "<" {
			current.moveWest()
		} else if move == "^" {
			current.moveNorth()
		} else if move == "v" {
			current.moveSouth()

		}
		locationSet[current] = true
	}

	return len(locationSet)
}

func day3Part2(input string) int {
	moves := strings.Split(input, "")
	locationSet := make(map[location]bool)
	santa := location{x: 0, y: 0}
	roboSanta := location{x: 0, y: 0}

	current := &santa
	locationSet[*current] = true

	for i, move := range moves {
		if i%2 == 0 {
			current = &santa
		} else {
			current = &roboSanta
		}

		if move == ">" {
			current.moveEast()
		} else if move == "<" {
			current.moveWest()
		} else if move == "^" {
			current.moveNorth()
		} else if move == "v" {
			current.moveSouth()
		}

		locationSet[*current] = true
	}

	return len(locationSet)
}
