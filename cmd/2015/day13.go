/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: "Aoc 2015 Day 13",

	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "13", false)
		fmt.Printf("Part 1: %d\n", day13Part1(input))
		fmt.Printf("Part 2: %d\n", day13Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day13Cmd)
}

func day13Part1(input string) int {

	personMap := make(Graph)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		// Alice would gain 54 happiness units by sitting next to Bob.
		personOne := parts[0]
		amount := cast.ToInt(parts[3])
		if parts[2] == "lose" {
			amount *= -1
		}

		personTwo := strings.ReplaceAll(parts[10], ".", "")

		if personMap[personOne] == nil {
			personMap[personOne] = make(map[string]int)
		}

		if personMap[personTwo] == nil {
			personMap[personTwo] = make(map[string]int)
		}

		personMap[personOne][personTwo] += amount
		personMap[personTwo][personOne] += amount

	}

	dfsMax := maxDistance(personMap, "Alice", map[string]bool{"Alice": true})
	return dfsMax
}

func day13Part2(input string) int {

	personMap := make(Graph)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		// Alice would gain 54 happiness units by sitting next to Bob.
		personOne := parts[0]
		amount := cast.ToInt(parts[3])
		if parts[2] == "lose" {
			amount *= -1
		}

		personTwo := strings.ReplaceAll(parts[10], ".", "")

		if personMap[personOne] == nil {
			personMap[personOne] = make(map[string]int)
		}

		if personMap[personTwo] == nil {
			personMap[personTwo] = make(map[string]int)
		}

		if personMap["me"] == nil {
			personMap["me"] = make(map[string]int)
		}

		personMap["me"][personOne] = 0
		personMap[personOne]["me"] = 0
		personMap["me"][personTwo] = 0
		personMap[personTwo]["me"] = 0

		personMap[personOne][personTwo] += amount
		personMap[personTwo][personOne] += amount

	}

	dfsMax := maxDistance(personMap, "Alice", map[string]bool{"Alice": true})
	return dfsMax
}

func (g Graph) printGraph() {
	for k := range g {
		for v := range g[k] {
			fmt.Print("  ", v, "[", g[k][v], "]")
		}
	}
}

func maxDistance(graph Graph, entry string, visited map[string]bool) int {
	if len(visited) == len(graph) {
		// We have reached the end but we need to make sure we complete the cycle
		return graph[entry]["Alice"]
	}

	maxDis := 0

	for k := range graph {
		if !visited[k] {
			visited[k] = true
			weight := graph[entry][k]
			maxRec := maxDistance(graph, k, visited)
			maxDis = util.Max(maxDis, maxRec+weight)

			delete(visited, k)
		}

	}

	return maxDis

}
