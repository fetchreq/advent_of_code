/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package eightteen

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day01 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "AoC 2018 Day 01",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2018", "3", false)
		fmt.Printf("Part 1: %d\n", day03Part1(input))
		fmt.Printf("Part 2: %d\n", day03Part2(input))
	},
}

func init() {
	cmd.EightTeenCmd.AddCommand(day03Cmd)
}

type point struct {
	x, y int
}
type block struct {
	id             string
	x1, x2, y1, y2 int
}

func day03Part1(input string) int {
	blocks := []block{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " @ ")
		dims := strings.Split(parts[1], ": ")
		dimsFirst := strings.Split(dims[0], ",")
		size := strings.Split(dims[1], "x")
		left := cast.ToInt(dimsFirst[0])
		top := cast.ToInt(dimsFirst[1])

		width := cast.ToInt(size[0])
		heigth := cast.ToInt(size[1])

		curr := block{
			id: parts[0],
			x1: left,
			y1: top,
			x2: left + width,
			y2: top + heigth,
		}
		blocks = append(blocks, curr)
	}

	totalOverLap := 0
	for i := 0; i < len(blocks)-1; i++ {
		block1 := blocks[i]
		for j := i + 1; j < len(blocks); j++ {
			block2 := blocks[j]
			if block1.isOverLapping(block2) {
				xLeft := util.Max(block1.x1, block2.x1)
				yTop := util.Max(block1.y1, block2.y1)
				xRight := util.Min(block1.x2, block2.x2)
				yBottom := util.Min(block1.y2, block2.y2)

				if xRight < xLeft || yBottom < yTop {
					continue
				}
				area := (xRight - xLeft) * (yBottom - yTop)
				totalOverLap += area

			}
		}
	}

	return totalOverLap
}

func (b block) isOverLapping(other block) bool {
	return b.x1 < other.x2 && b.y1 < other.y2 && other.x1 < b.x2 && other.y1 < b.y2
}

func day03Part2(input string) int {
	return 0
}
