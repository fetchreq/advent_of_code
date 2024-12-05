/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2024", "5", false)
		fmt.Printf("Day 5 Part 1 = %d\n", day5Part1(input))
		fmt.Printf("Day 5 Part 2 = %d\n", day5Part2(input)) // 5264 too low 5583 too big
	},
}

func init() {
	cmd.TwentyFourCmd.AddCommand(day5Cmd)
}

func getRulesAndUpdaes(data string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	for i := range updates {
		updates[i] = make([]int, 0)
	}
	inputParts := strings.Split(data, "\n\n")

	for _, line := range strings.Split(inputParts[0], "\n") {

		splits := strings.Split(line, "|")
		lhs := cast.ToInt(splits[0])
		rhs := cast.ToInt(splits[1])

		rules[lhs] = append(rules[lhs], rhs)
	}

	for _, line := range strings.Split(inputParts[1], "\n") {
		splits := strings.Split(line, ",")
		update := make([]int, 0, len(splits))
		for _, i := range splits {
			update = append(update, cast.ToInt(i))
		}
		updates = append(updates, update)
	}
	return rules, updates
}
func day5Part1(input string) int {
	rules, updates := getRulesAndUpdaes(input)

	sum := 0
	for _, update := range updates {
		passed := true
		for i, page := range update {
			if i == len(update)-1 {
				break
			}
			if !gt(page, update[i+1], rules) {
				passed = false
				break
			}
		}
		if passed {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func gt(a int, b int, rules map[int][]int) bool {
	return slices.Contains(rules[a], b)
}

func day5Part2(input string) int {
	rules, updates := getRulesAndUpdaes(input)

	sum := 0
	for _, update := range updates {
		passed := true
		for i, page := range update {
			if i == len(update)-1 {
				break
			}
			if !gt(page, update[i+1], rules) {
				passed = false
				break
			}
		}
		if !passed {
			sort.SliceStable(update, func(i, j int) bool {
				return gt(update[i], update[j], rules)
			})
			sum += update[len(update)/2]
		}
	}
	return sum
}
