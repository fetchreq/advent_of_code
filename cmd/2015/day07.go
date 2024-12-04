/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"math"
	"regexp"

	// "math"
	// "strconv"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Aoc 2015 Day 7",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day7 called")
		input := util.ReadFile("2015", "day7", false)
		part1, part2 := day7Part1(input)
		fmt.Printf("Part 1: %d\n", part1)
		fmt.Printf("Part 2: %d\n", part2)
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day7Part1(input string) (int, int) {
	wireToRule := map[string]string{}

	// generate graph of wires to their source rule
	for _, inst := range strings.Split(input, "\n") {
		if len(strings.TrimSpace(inst)) == 0 {
			continue
		}
		parts := strings.Split(inst, " -> ")
		wireToRule[parts[1]] = parts[0]
	}

	afirst := memoDFS(wireToRule, "a", map[string]int{})

	wireToRule["b"] = cast.ToString(afirst)
	aSecond := memoDFS(wireToRule, "a", map[string]int{})

	return afirst, aSecond
}

func memoDFS(graph map[string]string, entry string, memo map[string]int) int {
	if memoVal, ok := memo[entry]; ok {
		return memoVal
	}

	// if it's a number, return the casted value
	if regexp.MustCompile("[0-9]").MatchString(entry) {
		return cast.ToInt(entry)
	}

	sourceRule := graph[entry]
	parts := strings.Split(sourceRule, " ")

	var result int
	switch {
	case len(parts) == 1:
		result = memoDFS(graph, parts[0], memo)
	case parts[0] == "NOT":
		start := memoDFS(graph, parts[1], memo)
		result = (math.MaxUint16) ^ start
	case parts[1] == "AND":
		result = memoDFS(graph, parts[0], memo) & memoDFS(graph, parts[2], memo)
	case parts[1] == "OR":
		result = memoDFS(graph, parts[0], memo) | memoDFS(graph, parts[2], memo)
	case parts[1] == "LSHIFT":
		result = memoDFS(graph, parts[0], memo) << memoDFS(graph, parts[2], memo)
	case parts[1] == "RSHIFT":
		result = memoDFS(graph, parts[0], memo) >> memoDFS(graph, parts[2], memo)
	}

	memo[entry] = result
	return result
}
