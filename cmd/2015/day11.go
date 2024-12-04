/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "Aoc 2015 Day 11",

	Run: func(cmd *cobra.Command, args []string) {

		// input := "abz"
		oldPassword := "hepxcrrq"
		part1 := day11Part1(oldPassword)

		fmt.Printf("Part 1: %s\n", part1)

		// We have to update the password before we start
		// otherwise we already have a valid password
		part2 := day11Part1(updatePassword(part1))
		fmt.Printf("Part 2: %s\n", part2)

		//fmt.Printf("Part 1: %s", day11Part1(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day11Cmd)
}

func day11Part1(input string) string {

	password := input
	for !isValid(password) {
		password = updatePassword(password)
	}
	return password
}

func updatePassword(input string) string {
	chars := strings.Split(input, "")
	for i := len(input) - 1; i >= 0; i-- {

		if chars[i] == "z" {
			chars[i] = "a"
		} else {
			chars[i] = string(byte(input[i]) + 1)
			break
		}
	}
	return strings.Join(chars, "")
}

func isValid(input string) bool {
	// Must include one increasing straight of at least three letters
	// abc, bcd, cde, and so on up to xyz
	ruleOne := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			if (s[i]+1) == s[i+1] && (s[i]+2 == s[i+2]) {
				return true
			}
		}
		return false
	}

	ruleTwo := func(s string) bool {
		return !regexp.MustCompile("[iol]").MatchString(s)
	}

	ruleThree := func(s string) bool {
		count := 0
		for i := 0; i < len(s)-1; i++ {
			if (s[i]) == s[i+1] {
				count += 1
				i += 1
			}
		}
		return count >= 2
	}

	return ruleOne(input) && ruleTwo(input) && ruleThree(input)
}
