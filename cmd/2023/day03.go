/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day03 called")
		input := util.ReadFile("2023", "3", false)
		fmt.Printf("Part 1: %d\n", day3Part1(input))
		fmt.Printf("Part 2: %d\n", day3Part2(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day03Cmd)
}

func day3Part1(input string) int {
	matrix := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(row, ""))
	}
	r := regexp.MustCompile(`\d`)
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			var num strings.Builder

			if r.MatchString(matrix[i][j]) {
				k := j
				// get all the digits of the number
				for k < len(matrix[i]) && r.MatchString(matrix[i][k]) {
					num.WriteString(matrix[i][k])
					k++
				}

				if checkNeighborsForSymbol(i, j, num.Len(), len(matrix)-1, len(matrix[0])-1, matrix) {
					sum += cast.ToInt(num.String())
				}
				// skip the rest of the digits since they have all been checked
				j = k

			}

		}
	}
	return sum
}

func day3Part2(input string) int {
	matrix := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(row, ""))
	}
	r := regexp.MustCompile(`\d`)
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {

			if matrix[row][col] == "*" {

				numSet := make(map[int]bool)
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if r.MatchString(matrix[row+i][col+j]) {
							numSet[getNum(row+i, col+j, matrix)] = true
						}

					}
				}

				if len(numSet) == 2 {
					running := 1
					for key := range numSet {
						running *= key
					}

					sum += running
				}
			}
		}
	}
	return sum
}

func getNum(row, col int, matrix [][]string) int {
	r := regexp.MustCompile(`\d`)
	var num strings.Builder
	for col > 0 && r.MatchString(matrix[row][col-1]) {
		col--

	}

	for col < len(matrix[0]) && r.MatchString(matrix[row][col]) {
		num.WriteString(matrix[row][col])
		col++
	}
	return cast.ToInt(num.String())
}

// Starts at the first col that is a number in the matrix and checks all neighbors for lenght of the number
func checkNeighborsForSymbol(row, col, numLength, numRows, numCols int, matrix [][]string) bool {
	r := regexp.MustCompile(`\d`)
	if row == 0 {
		for k := util.Max(col-1, 0); k < util.Min(numCols, col+numLength+1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) ||
				(matrix[row+1][k] != "." && !r.MatchString(matrix[row+1][k])) {
				return true
			}
		}
	} else if row == numRows {
		for k := util.Max(col-1, 0); k < util.Min(numCols, col+numLength+1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) ||
				(matrix[row-1][k] != "." && !r.MatchString(matrix[row-1][k])) {
				return true
			}
		}
	} else {
		for k := util.Max(col-1, 0); k < util.Min(numCols, col+numLength+1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) ||
				(matrix[row-1][k] != "." && !r.MatchString(matrix[row-1][k])) ||
				(matrix[row+1][k] != "." && !r.MatchString(matrix[row+1][k])) {
				return true
			}
		}
	}
	return false
}
