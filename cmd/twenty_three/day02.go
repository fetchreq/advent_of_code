/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day02Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day02Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day2Part1(input string) int {
	sum := 0
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(row, ":")
		var gameIdStr string
		fmt.Sscanf(parts[0], "Game %s:", &gameIdStr)
		gameId := cast.ToInt(gameIdStr)
		validGame := true

		for _, set := range strings.Split(parts[1], ";") {
			for _, marble := range strings.Split(set, ",") {
				marble = strings.TrimSpace(marble)
				var color, amountStr string
				fmt.Sscanf(marble, "%s %s", &amountStr, &color)
				amount := cast.ToInt(amountStr)
				if color == "blue" && amount > 14 {
					validGame = false

				} else if color == "green" && amount > 13 {
					validGame = false

				} else if color == "red" && amount > 12 {
					validGame = false

				}
			}

		}

		if validGame {
			sum += gameId
		}
		
	}
	return sum
}

func day2Part2(input string) int {
	sum := 0
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(row, ":")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, set:= range strings.Split(parts[1], ";") {
			for _, marble := range strings.Split(set, ",") {
				marble = strings.TrimSpace(marble)
				var color, amountStr string
				fmt.Sscanf(marble, "%s %s", &amountStr, &color)
				amount := cast.ToInt(amountStr)
				if color == "blue" {
					maxBlue = util.Max(maxBlue, amount)

				} else if color == "green" {
					maxGreen = util.Max(maxGreen, amount)

				} else if color == "red" {
					maxRed = util.Max(maxRed, amount)
				}
			}

		}

		sum += (maxRed * maxGreen * maxBlue)
	}
	return sum
}
