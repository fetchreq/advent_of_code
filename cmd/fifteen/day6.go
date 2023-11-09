/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Aoc 2015 Day 6",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day6", false)
		fmt.Printf("Part 1: %d", day6Part1(input))
		fmt.Printf("Part 2: %d", day6Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
type InstructionType int
const (
	toggle InstructionType = 0
	turn InstructionType = 1
)

type Instruction struct {
	instructionType InstructionType
	isOn bool
	startX int
	endX int
	startY int
	endY int
}

func newInstruction(input string ) Instruction {
	parts := strings.Split(input, " ")
	if (parts[0] == "turn") {
		instructionType := turn
		isOn := parts[1] == "on";


		start := strings.Split(parts[2], ",")

		startX, err := strconv.Atoi(start[0])
		util.CheckErr(err)
		startY, err := strconv.Atoi(start[1])
		util.CheckErr(err)

		end := strings.Split(parts[4], ",")
		endX, err := strconv.Atoi(end[0])
		util.CheckErr(err)
		endY, err := strconv.Atoi(end[1])
		util.CheckErr(err)

		return Instruction{instructionType: instructionType, isOn: isOn, startX: startX, startY: startY, endX: endX, endY: endY}

	} else if parts[0] == "toggle" {
		instructionType := toggle
		start := strings.Split(parts[1], ",")
		startX, err := strconv.Atoi(start[0])
		util.CheckErr(err)
		startY, err := strconv.Atoi(start[1])
		util.CheckErr(err)

		end := strings.Split(parts[3], ",")
		endX, err := strconv.Atoi(end[0])
		util.CheckErr(err)
		endY, err := strconv.Atoi(end[1])
		util.CheckErr(err)

		return Instruction{instructionType: instructionType, isOn: false, startX: startX, startY: startY, endX: endX, endY: endY}

	}
	return Instruction{instructionType: toggle, isOn: false, startX: 0, startY: 0, endX: 0, endY: 0}

}

func day6Part1(input string) int {
	matrix := [1000][1000]bool{}

	for _, val := range strings.Split(input, "\n") {
		if len(strings.TrimSpace(val)) == 0 {
			continue
		}

		instruction := newInstruction(val);

		if (instruction.instructionType == turn) {

			for i := instruction.startX; i <= instruction.endX; i++ {
				for j := instruction.startY; j <= instruction.endY; j++ {
					matrix[i][j] = instruction.isOn
				}
			}


		} else {

			for i := instruction.startX; i <= instruction.endX; i++ {
				for j := instruction.startY; j <= instruction.endY; j++ {
					matrix[i][j] = !matrix[i][j]
				}
			}

		}


	}
	onCount := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if matrix[i][j] {
				onCount += 1;
			}
		}
	}


	return onCount
}

func day6Part2(input string) int {
	matrix := [1000][1000]int{}

	for _, val := range strings.Split(input, "\n") {
		if len(strings.TrimSpace(val)) == 0 {
			continue
		}

		instruction := newInstruction(val);

		if (instruction.instructionType == turn) {

			for i := instruction.startX; i <= instruction.endX; i++ {
				for j := instruction.startY; j <= instruction.endY; j++ {
					if instruction.isOn {
						matrix[i][j] += 1
					} else if !instruction.isOn && matrix[i][j] > 0 {
						matrix[i][j] -= 1
					}
				}
			}


		} else {

			for i := instruction.startX; i <= instruction.endX; i++ {
				for j := instruction.startY; j <= instruction.endY; j++ {
					matrix[i][j] += 2
				}
			}

		}


	}
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			total += matrix[i][j];
			
		}
	}


	return total
}
