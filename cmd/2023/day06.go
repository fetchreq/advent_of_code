/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

type boatRace struct {
	time     float64
	distance float64
}

// day06Cmd represents the day06 command
var day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: "AoC 2023 Day 6",
	Run: func(cmd *cobra.Command, args []string) {

		input := util.ReadFile("2023", "6", false)

		fmt.Printf("Part 1: %d\n", day6Part1(input))
		fmt.Printf("Part 2: %d\n", day6Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day06Cmd)

}

func day6Part1(input string) int {
	boatRaces := getBoatRaces(input)

	winning := []int{}
	for _, boatRace := range boatRaces {

		a := 1.0
		b := boatRace.time
		c := boatRace.distance
		sqrtPart := math.Sqrt(math.Pow(b, 2) - 4*a*c)

		minHold := math.Floor((b - sqrtPart) / 2)
		maxHold := math.Floor((b + sqrtPart) / 2)

		winning = append(winning, int(maxHold-minHold))
	}

	product := 1
	for _, winCount := range winning {
		product *= winCount
	}

	return product
}

func day6Part2(input string) int {
	boatRace := getBoatRace(input)

	a := 1.0
	b := boatRace.time
	c := boatRace.distance

	sqrtPart := math.Sqrt(math.Pow(b, 2) - 4*a*c)

	minHold := math.Floor((b - sqrtPart) / 2)
	maxHold := math.Floor((b + sqrtPart) / 2)

	return int(maxHold - minHold)
}

func getBoatRaces(input string) []boatRace {
	boatRaces := []boatRace{}
	rows := strings.Split(input, "\n")

	distanceRow := strings.TrimPrefix(rows[1], "Distance: ")
	// Get a list of distance values
	distanceFields := strings.Fields(distanceRow)

	timeRow := strings.TrimPrefix(rows[0], "Time: ")
	for idx, val := range strings.Fields(timeRow) {
		time, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not parse %s to float", timeRow))
		}

		distance, err := strconv.ParseFloat(distanceFields[idx], 64)
		if err != nil {
			panic(fmt.Sprintf("Could not parse %s to float", timeRow))
		}
		boatRaces = append(boatRaces, boatRace{time: time, distance: distance})
	}

	return boatRaces

}

func getBoatRace(input string) boatRace {

	rows := strings.Split(input, "\n")

	timeRow := strings.TrimPrefix(rows[0], "Time: ")
	timeRow = strings.ReplaceAll(timeRow, " ", "")
	time, err := strconv.ParseFloat(timeRow, 64)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to float", timeRow))
	}

	distanceRow := strings.TrimPrefix(rows[1], "Distance: ")
	distanceRow = strings.ReplaceAll(distanceRow, " ", "")
	distance, err := strconv.ParseFloat(distanceRow, 64)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to float", distanceRow))
	}

	return boatRace{time: time, distance: distance}

}
