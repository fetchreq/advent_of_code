/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Aoc 2015 Day 5",

	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day5", false);
		fmt.Printf("Part 1: %d\n", day5Part1(input));
		fmt.Printf("Part 2: %d\n", day5Part2(input));
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day5Part1(input string) int {
	niceStrings := 0
	stingList := strings.Split(input, "\n")
	for _, val := range stingList {
		if strings.Contains(val, "ab") || strings.Contains(val, "cd")|| strings.Contains(val, "pq") || strings.Contains(val, "xy") {
			continue
		}

		if hasThreeVowels(val) && hasRepeatingChars(val) {
			niceStrings += 1;
		}


	}
				

	return niceStrings
}

func day5Part2(input string) int {
	niceStrings := 0

	for _, val := range strings.Split(input, "\n") {
		if doRuleOneCheck(val) && doRuleTwoCheck(val) {
			niceStrings += 1
		}

	}
				

	return niceStrings
}

func doRuleOneCheck(val string) bool {
	for i := 0; i < len(val) - 2; i++ {

		for j := i + 2; j < len(val) - 1; j++ {
			if val[i] == val[j] && val[i+1] == val[j+1] {
				return true
			}
		}
	}
	return false
}
func doRuleTwoCheck(val string) bool {
	for i := 0; i < len(val) - 2; i++ {
		if val[i] == val[i + 2] {
				return true
			
		}
	}
	return false
}


func hasThreeVowels(s string) bool {
	vowelCount := 0
	for _, r := range strings.ToLower(s) {
		if strings.ContainsRune("aeiou", r) {
			vowelCount += 1
		}
	}

	return vowelCount >= 3

}

func hasRepeatingChars(s string) bool {
	var rr rune
	for _, r := range strings.ToLower(s) {
		if rr == r {
			return true
		} 

		rr = r
	}

	return false

}
