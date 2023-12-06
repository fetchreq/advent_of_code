/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
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
	// for i := 1; i < race.time; i++ {
	// 	speed := i;
	// 	timeLeft := race.time - i;
	// 	totalDistance := speed * timeLeft
	// 	if totalDistance > race.distance {
	// 		winCount++
	// 	}
	// }
	//

	return int(maxHold-minHold);
}
func getBoatRaces(input string) []boatRace {
	boatRaces := []boatRace{}
	times := []int{}
	distances := []int{}
	for idx, row := range strings.Split(input, "\n") {
		for  _, val := range strings.Split(row, " ") {
			if strings.TrimSpace(val) == "" {
				continue
			}
			if idx == 0 {
				if num, err := strconv.Atoi(val); err == nil {
					times = append(times, num)
				}
			} else {
				if val, err := strconv.Atoi(val); err == nil {
					distances = append(distances, val)
				}
			}

		}
	}

	for idx, time := range times {
		boatRaces = append(boatRaces, boatRace{time: float64(time), distance: float64(distances[idx])})
	}

	return boatRaces

}

func getBoatRace(input string) boatRace {

	time := ""
	distance := ""
	for idx, row := range strings.Split(input, "\n") {
		for  _, val := range strings.Split(row, " ") {
			if strings.TrimSpace(val) == "" {
				continue
			}
			if idx == 0 {
				if _, err := strconv.Atoi(val); err == nil {
					time += val
				}
			} else {
				if _, err := strconv.Atoi(val); err == nil {
					distance += val
				}
			}

		}
	}
	return boatRace{time: float64(cast.ToInt(time)), distance: float64(cast.ToInt(distance))}


}
