/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "AoC 2023 Day 2",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2023", "2", false)
		fmt.Printf("Part 1: %d\n", day2Part1(input))
		fmt.Printf("Part 2: %d\n", day2Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day02Cmd)
}

const (
	RedMax = 12
	GreenMax = 13
	BlueMax = 14
)

func day2Part1(input string) int {
	sum := 0
	for gameId, row := range strings.Split(input, "\n") {
		validGame := true
		red := regexp.MustCompile(`(\d+) red`)
		for _, match := range red.FindAllStringSubmatch(row, -1){
			if cast.ToInt(match[1]) > RedMax {
				validGame = false
			}
		}
		green := regexp.MustCompile(`(\d+) green`)
		for _, match := range green.FindAllStringSubmatch(row, -1){
			if cast.ToInt(match[1]) > GreenMax {
				validGame = false
			}
		}

		blue := regexp.MustCompile(`(\d+) blue`)
		for _, match := range blue.FindAllStringSubmatch(row, -1){
			if cast.ToInt(match[1]) > BlueMax {
				validGame = false
			}
		}
		if validGame {
			sum +=  gameId + 1
		}
	}
	return sum
}

func day2Part2(input string) int {
	sum := 0
	for _, row := range strings.Split(input, "\n") {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		red := regexp.MustCompile(`(\d+) red`)
		for _, match := range red.FindAllStringSubmatch(row, -1){
			maxRed = util.Max(cast.ToInt(match[1]), maxRed) 
		}
		green := regexp.MustCompile(`(\d+) green`)
		for _, match := range green.FindAllStringSubmatch(row, -1){
			maxGreen = util.Max(cast.ToInt(match[1]), maxGreen) 
		}

		blue := regexp.MustCompile(`(\d+) blue`)
		for _, match := range blue.FindAllStringSubmatch(row, -1){
			maxBlue = util.Max(cast.ToInt(match[1]), maxBlue) 
		}


		sum += (maxRed * maxGreen * maxBlue)
	}
	return sum
}
