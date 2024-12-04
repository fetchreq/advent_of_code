/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "AoC 2023 Day 2",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2023", "2", true)
		fmt.Printf("Part 1: %d\n", day2Part1(input))
		fmt.Printf("Part 2: %d\n", day2Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day02Cmd)
}

const (
	RedMax   = 12
	GreenMax = 13
	BlueMax  = 14
)

type game struct {
	match     *regexp.Regexp
	color     string
	maxAmount int
}

func day2Part1(input string) int {
	sum := 0
	colors := [3]game{
		game{color: "red", match: regexp.MustCompile(`(\d+) red`), maxAmount: RedMax},
		game{color: "green", match: regexp.MustCompile(`(\d+) green`), maxAmount: GreenMax},
		game{color: "blue", match: regexp.MustCompile(`(\d+) blue`), maxAmount: BlueMax},
	}
	for gameId, row := range strings.Split(input, "\n") {
		validGame := true

		for _, color := range colors {
			for _, amount := range getMarbleAmounts(color.match, row) {
				if amount > color.maxAmount {
					validGame = false
				}
			}
		}

		if validGame {
			sum += gameId + 1
		}
	}
	return sum
}

func day2Part2(input string) int {
	sum := 0
	colors := [3]game{
		game{color: "red", match: regexp.MustCompile(`(\d+) red`), maxAmount: RedMax},
		game{color: "green", match: regexp.MustCompile(`(\d+) green`), maxAmount: GreenMax},
		game{color: "blue", match: regexp.MustCompile(`(\d+) blue`), maxAmount: BlueMax},
	}

	for _, row := range strings.Split(input, "\n") {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, color := range colors {
			for _, amount := range getMarbleAmounts(color.match, row) {
				switch color.color {
				case "red":
					maxRed = util.Max(amount, maxRed)
				case "green":
					maxGreen = util.Max(amount, maxGreen)
				case "blue":
					maxBlue = util.Max(amount, maxBlue)
				}
			}
		}
		sum += (maxRed * maxGreen * maxBlue)
	}
	return sum
}

func getMarbleAmounts(match *regexp.Regexp, input string) []int {
	amounts := []int{}
	for _, match := range match.FindAllStringSubmatch(input, -1) {
		amounts = append(amounts, cast.ToInt(match[1]))
	}

	return amounts

}
