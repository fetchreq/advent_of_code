/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strings"

	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day08Cmd represents the day08 command
var day08Cmd = &cobra.Command{
	Use:   "day08",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day08 called")
		input := util.ReadFile("2023", "8", false)
		fmt.Printf("Part 1: %d\n", day8Part1(input))
		fmt.Printf("Part 2: %d\n", day8Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day08Cmd)
}

type nodePair struct {
	left  string
	right string
}

func day8Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	path := parts[0]
	nodeMap := make(map[string]nodePair)
	for _, row := range strings.Split(parts[1], "\n") {
		var nodeId, leftNode, rightNode string
		row = strings.ReplaceAll(row, ",", "")
		row = strings.ReplaceAll(row, ")", "")
		row = strings.ReplaceAll(row, "(", "")
		fmt.Sscanf(row, "%s = %s %s", &nodeId, &leftNode, &rightNode)
		nodeMap[nodeId] = nodePair{left: leftNode, right: rightNode}
	}

	atEnd := func(input string) bool {
		return input == "ZZZ"
	}

	return solve("AAA", path, atEnd, nodeMap)
}

type AtEndNode func(string) bool

func solve(start string, path string, fn AtEndNode, nodeMap map[string]nodePair) int {
	currNode := start
	stepCount := 0
	found := false
	for !found {
		for _, instruction := range strings.Split(path, "") {
			pair := nodeMap[currNode]
			if instruction == "L" {
				currNode = pair.left
			} else {
				currNode = pair.right
			}

			stepCount++
			if fn(currNode) {
				found = true
				break
			}

		}
	}

	return stepCount
}

func day8Part2(input string) int {
	parts := strings.Split(input, "\n\n")
	path := parts[0]
	nodeMap := make(map[string]nodePair)
	nodes := []string{}
	for _, row := range strings.Split(parts[1], "\n") {
		var nodeId, leftNode, rightNode string
		row = strings.ReplaceAll(row, ",", "")
		row = strings.ReplaceAll(row, ")", "")
		row = strings.ReplaceAll(row, "(", "")
		fmt.Sscanf(row, "%s = %s %s", &nodeId, &leftNode, &rightNode)
		if strings.HasSuffix(nodeId, "A") {
			nodes = append(nodes, nodeId)
		}
		nodeMap[nodeId] = nodePair{left: leftNode, right: rightNode}
	}
	stepCounts := []int{}
	atEnd := func(input string) bool {
		return strings.HasSuffix(input, "Z")
	}
	for _, currNode := range nodes {
		stepCounts = append(stepCounts, solve(currNode, path, atEnd, nodeMap))
	}

	return LCM(stepCounts[0], stepCounts[1], stepCounts...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
