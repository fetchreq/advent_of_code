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

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Aoc 2015 Day 8",

	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day8", false)
		fmt.Println("day8 called")
		fmt.Printf("Part 1: %d\n", day8Part1(input))
		fmt.Printf("Part 2: %d\n", day8Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day8Cmd)
}

func day8Part1(input string ) int {
	var codeChars, stringchars int
	for _, line := range strings.Split(input, "\n"){
		codeChars += len(line)
		for i := 1; i < len(line) - 1; i++ {
			if line[i] == '\\' && (line[i+1] == '\\' || line[i+1] == '"') {
				i += 1
			} else if  line[i] == '\\' && line[i + 1] == 'x' {
				i += 3
			} 
			
			stringchars += 1
		}
	}
	return codeChars - stringchars
}

func day8Part2(input string ) int {
	var codeChars, encodedChars int
	for _, line := range strings.Split(input, "\n"){
		codeChars += len(line)
		encodedChars += 2
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"', '\\':
				encodedChars += 2
			default:
				encodedChars += 1
			}

		}

	}
	return encodedChars - codeChars
}
