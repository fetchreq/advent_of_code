/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

type boatRace struct {
	time float64
	distance float64
}

// day06Cmd represents the day06 command
var day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: "AoC 2023 Day 6",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day06 called")

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
	 
		part := math.Sqrt(math.Pow(b, 2) - 4 * a * c)

		minHold := math.Floor((b - part) / 2);
		maxHold := math.Floor((b + part) / 2);

		winning = append(winning, int(maxHold-minHold))
	}

	product := 1
	for _, winCount := range winning {
		product *= winCount
	}

	return product;
}

func day6Part2(input string) int {
	boatRace := getBoatRace(input)
	
	a := 1.0
	b := boatRace.time
	c := boatRace.distance

	part := math.Sqrt(math.Pow(b, 2) - 4 * a * c)

	minHold := math.Floor((b - part) / 2);
	maxHold := math.Floor((b + part) / 2);

	return int(maxHold-minHold);
}

func getBoatRaces(input string) []boatRace {
	boatRaces := []boatRace{}
	times := []float64{}
	distances := []float64{}
	for _, row := range strings.Split(input, "\n") {
		// Check if we are on the time row
		isTime := strings.HasPrefix(row, "Time: ")
		if isTime {
			row = strings.TrimPrefix(row, "Time: ")
		} else {
			row = strings.TrimPrefix(row, "Distance: ")
		}
		// Get all the numbers
		for  _, val := range strings.Fields(row) {
			if isTime {
				num, _ := strconv.ParseFloat(val, 64)
				times = append(times, num)
			} else {
				num, _ := strconv.ParseFloat(val, 64)
				distances = append(distances, num)
			}

		}
	}

	for idx, time := range times {
		boatRaces = append(boatRaces, boatRace{time: time, distance: distances[idx]})
	}

	return boatRaces

}

func getBoatRace(input string) boatRace {

	var time float64
	var distance float64

	for _, row := range strings.Split(input, "\n") {
		isTime := strings.HasPrefix(row, "Time: ")
		if isTime {
			row = strings.TrimPrefix(row, "Time: ")
			row = strings.ReplaceAll(row, " ", "")
			time, _ = strconv.ParseFloat(row, 64)
		} else {
			row = strings.TrimPrefix(row, "Distance: ")
			row = strings.ReplaceAll(row, " ", "")
			distance, _ = strconv.ParseFloat(row, 64)
		}
	}
	return boatRace{time: time, distance: distance}


}
