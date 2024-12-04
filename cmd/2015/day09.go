/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"math"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Aoc 2015 Day 9",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day9 called")
		input := util.ReadFile("2015", "day9", false)
		fmt.Printf("Part 1: %d\n", day9Part1(input))
		fmt.Printf("Part 2: %d\n", day9Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day9Cmd)
}

type Graph map[string]map[string]int

func day9Part1(input string) int {
	graph := make(Graph)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		start, end := parts[0], parts[2]
		weight := cast.ToInt(parts[4])
		if graph[start] == nil {
			graph[start] = make(map[string]int)
		}

		if graph[end] == nil {
			graph[end] = make(map[string]int)
		}
		graph[start][end] = weight
		graph[end][start] = weight
	}

	min := math.MaxInt
	for k := range graph {
		fmt.Println(k)
		dfsMin, _ := distance(graph, k, map[string]bool{k: true})

		min = util.Min(min, dfsMin)
	}

	return min
}
func day9Part2(input string) int {
	graph := make(Graph)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		start, end := parts[0], parts[2]
		weight := cast.ToInt(parts[4])
		if graph[start] == nil {
			graph[start] = make(map[string]int)
		}

		if graph[end] == nil {
			graph[end] = make(map[string]int)
		}
		graph[start][end] = weight
		graph[end][start] = weight
	}
	max := 0
	for k := range graph {
		_, dfsMax := distance(graph, k, map[string]bool{k: true})

		max = util.Max(max, dfsMax)
	}

	return max
}

func distance(graph Graph, entry string, visited map[string]bool) (minDistance int, maxDistance int) {
	if len(visited) == len(graph) {
		return 0, 0
	}

	minDis := math.MaxInt
	maxDis := 0

	for k := range graph {
		if !visited[k] {
			visited[k] = true
			weight := graph[entry][k]
			minRec, maxRec := distance(graph, k, visited)
			minDis = util.Min(minDis, weight+minRec)
			maxDis = util.Max(maxDis, weight+maxRec)

			delete(visited, k)
		}
	}

	return minDis, maxDis

}
