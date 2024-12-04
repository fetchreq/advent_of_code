/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"math"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day11Cmd represents the day08 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		part1, _ := cmd.Flags().GetBool("part1")
		input := util.ReadFile("2023", "11", false)
		if part1 {
			fmt.Printf("Part 1: %d\n", day11Part1(input))
		} else {
			fmt.Printf("Part 2: %d\n", day11Part2(input))
		}

	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day11Cmd)
	day11Cmd.PersistentFlags().Bool("part1", false, "")
}

// a galaxy position is a number: x + y * uw
var univ []int

// the univ dimensions: width and height of its grid. thus usize = uw*uh
var uw, uh int

// the list of galaxies in each row and col, as their ID + 1
var rows, cols [][]int

func day11Part1(input string) int {
	return computeDistances(strings.Split(input, "\n"), 1)
}

func day11Part2(input string) int {
	return computeDistances(strings.Split(input, "\n"), 1_000_000-1)
}

func computeDistances(lines []string, expansion int) (sum int) {
	parse(lines)
	expandUniv(expansion)
	for g1 := 0; g1 < len(univ)-1; g1++ {
		for g2 := g1 + 1; g2 < len(univ); g2++ {
			sum += distance(g1, g2)
		}
	}
	return

}

func parse(lines []string) {
	uw = len(lines[0])
	uh = len(lines)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '#' {
				galaxy := j + i*uw
				univ = append(univ, galaxy)
			}
		}
	}
}

func expandUniv(expansion int) {
	// compute rows and cols
	rows = make([][]int, uh)
	for i := range rows {
		rows[i] = []int{}
	}
	cols = make([][]int, uw)
	for i := range cols {
		cols[i] = []int{}
	}

	for gid, gal := range univ {
		rows[gal/uw] = append(rows[gal/uw], gid)
		cols[gal%uw] = append(cols[gal%uw], gid)
	}

	tmd := make([]int, len(univ), len(univ))
	tmr := make([]int, len(univ), len(univ))
	ouw := uw
	for i, exp := 0, 0; i < len(rows); i++ {
		if len(rows[i]) == 0 {
			exp += expansion
			uh += expansion
			continue
		}
		if exp == 0 {
			continue
		}
		for _, gid := range rows[i] {
			tmd[gid] = exp
		}
	}

	for i, exp := 0, 0; i < len(cols); i++ {
		if len(cols[i]) == 0 {
			exp += expansion
			uw += expansion
			continue
		}
		if exp == 0 {
			continue
		}
		for _, gid := range cols[i] {
			tmr[gid] = exp
		}
	}

	for gid, pos := range univ {
		ox := pos % ouw
		oy := pos / ouw
		x := ox + tmr[gid]
		y := oy + tmd[gid]
		univ[gid] = x + y*uw
	}
}

// mahattan distance
func distance(g1, g2 int) int {
	xVal := math.Abs(float64(univ[g2]%uw - univ[g1]%uw))
	yVal := math.Abs(float64(univ[g2]/uw - univ[g1]/uw))
	return int(xVal) + int(yVal)
}

func intAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
