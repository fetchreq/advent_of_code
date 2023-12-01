/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "AoC 2023 Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2023", "1", false)
		fmt.Printf("Part 1: %d\n", day1Part1(input))
		fmt.Printf("Part 2: %d\n", day1Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day1Cmd)
}

func day1Part1(input string) int {
	sum := 0
	for _, row := range strings.Split(input, "\n") {
		first := getFirstNumberInString(row)
		last := getFirstNumberInString(reverseString(row))
		val, _:= strconv.Atoi(fmt.Sprintf("%s%s", first, last))

		sum += val
	}
	return sum
}

func checkCharIsNumber(c byte) bool {
	return '0' <= c && c <= '9'
}


func day1Part2(input string) int {
	sum := 0

	numberStrings:= [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range strings.Split(input, "\n"){
		numBytes := make([]byte, 2)
		found := false
		for i := 0; i < len(line); i++ { // Check from start of line
			// Check for numeric value
			if checkCharIsNumber(line[i]) {
				numBytes[0] = line[i]
				found = true
			}
			// Check for number as string
			for numCheck, strCheck := range numberStrings {
				if line[i:util.Min(i+len(strCheck), len(line))] == strCheck {
					numBytes[0] = fmt.Sprint(numCheck + 1)[0]
					found = true
				}
			}
			if found { // Exit loop if number is found
				break
			}
		}
		found = false
		for j := len(line) - 1; j >= 0; j-- { // Check from end of line
			// Check for numeric value
			if checkCharIsNumber(line[j]) {
				numBytes[1] = line[j]
				found = true
			}
			// Check for number as string
			for numCheck, strCheck := range numberStrings {
				if line[j:util.Min(j+len(strCheck), len(line))] == strCheck {
					numBytes[1] = fmt.Sprint(numCheck + 1)[0]
					found = true
				}
			}
			if found { // Exit loop if number is found
				break
			}
		}
		num, _ := strconv.Atoi(string(numBytes))
		sum += num
	}

	return sum
}

func getFirstNumberInString(input string) string {
	for _, i := range strings.Split(input, "") {

		if _, err := strconv.Atoi(i); err == nil {
			return i
		}
	}

	return ""
}

func reverseString(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
