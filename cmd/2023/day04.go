/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day04 called")
		input := util.ReadFile("2023", "4", false)

		fmt.Printf("Part 1: %d\n", day4Part1(input))
		fmt.Printf("Part 2: %d\n", day4Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day04Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day04Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day04Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day4Part1(input string) int {
	sum := 0

	for _, row := range strings.Split(input, "\n") {
		game := strings.Split(row, ":")
		gameNums := strings.Split(game[1], " | ")

		winners := createIntArray(strings.Split(gameNums[0], " "))
		values := createIntArray(strings.Split(gameNums[1], " "))

		count := 0
		for _, val := range values {
			if contains(winners, val) {
				if count == 0 {
					count++
				} else {
					count *= 2
				}
			}
		}
		sum += count
	}

	return sum
}

type cardValues struct {
	winners []int
	values []int
	count int
}


func day4Part2(input string) int {
	sum := 0
	games := make(map[int]int)
	for i, _ := range strings.Split(input, "\n") {
		games[i] = 1
	}
	for idx, row := range strings.Split(input, "\n") {
		game := strings.Split(row, ":")
		gameNums := strings.Split(game[1], " | ")

		winners := createIntArray(strings.Split(gameNums[0], " "))
		values := createIntArray(strings.Split(gameNums[1], " "))

		count := 0
		for _, val := range values {
			if contains(winners, val) {
				count++
			}

		}
		for i := 0; i < games[idx]; i++ {
			for j := 1; j <= count; j++ {
				games[idx+j] = games[idx+j] + 1
			}
		}
	}

	for _, value := range games {
		sum += value
	}

	return sum
}

func createIntArray(input []string) []int {
	values := []int{}
	for _, val := range input{
		if strings.TrimSpace(val) == "" {
			continue
		}
		values = append(values, cast.ToInt(strings.TrimSpace(val)))
	}
	return values
}

func contains(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}

	return false
}
