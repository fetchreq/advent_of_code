/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
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
		sum += getFirstNumberInString(row)*10 + getFirstNumberInString(reverseString(row))
	}
	return sum
}

func checkCharIsNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func day1Part2(input string) int {
	sum := 0

	numberStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, row := range strings.Split(input, "\n") {
		found := false
		var first int
		// from beginning of the word
		for i, letter := range row {
			if unicode.IsDigit(letter) {
				first = cast.ToInt(string(letter))
				found = true
			}
			for index, num := range numberStrings {
				if row[i:util.Min(i+len(num), len(row))] == num {
					first = index + 1
					found = true
				}
			}

			if found {
				break
			}
		}

		var last int
		found = false
		// from the end of the word
		for j := len(row) - 1; j >= 0; j-- {
			// current is a number
			if '0' <= row[j] && row[j] <= '9' {
				last = cast.ToInt(string(row[j]))
				found = true
			}

			// check for string number
			for index, num := range numberStrings {
				if row[j:util.Min(j+len(num), len(row))] == num {
					last = index + 1
					found = true
				}
			}

			if found {
				break
			}
		}
		num := first*10 + last
		sum += num
	}

	return sum
}

func getFirstNumberInString(input string) int {
	for _, i := range strings.Split(input, "") {

		if _, err := strconv.Atoi(i); err == nil {
			return cast.ToInt(i)
		}
	}

	return 0
}

func reverseString(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
