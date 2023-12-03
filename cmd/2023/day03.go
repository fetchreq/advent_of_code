/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
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

func day3Part1(input string) int{
	matrix := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(row, ""))	
	}
	r := regexp.MustCompile(`\d`)
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			var num string

			if r.MatchString(matrix[i][j]) {
				k := j;
				for k < len(matrix[i]) && r.MatchString(matrix[i][k]) {
					num = fmt.Sprintf("%s%s",num, matrix[i][k])
					k++
				}
				if checkNeighbors(i, j, len(num), len(matrix) -1, len(matrix[0]) - 1, matrix) {
					sum += cast.ToInt(num)
				} 				
				j = k

			}
		

		}
	}
	return sum
}

func day3Part2(input string) int{
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
				if r.MatchString(matrix[row-1][col-1]) {
					numSet[getNum(row - 1, col - 1, matrix)] = true
				}

				if r.MatchString(matrix[row-1][col]) {
					numSet[getNum(row - 1, col, matrix)] = true
				}

				if r.MatchString(matrix[row-1][col+1]) {
					numSet[getNum(row - 1, col + 1, matrix)] = true
				}

				if r.MatchString(matrix[row][col-1]) {
					numSet[getNum(row, col - 1, matrix)] = true
				}

				if r.MatchString(matrix[row][col+1]) {
					numSet[getNum(row, col + 1, matrix)] = true
				}

				if r.MatchString(matrix[row+1][col-1]) {
					numSet[getNum(row+1, col-1, matrix)] = true
				}

				if r.MatchString(matrix[row+1][col]) {
					numSet[getNum(row + 1, col, matrix)] = true
				}

				if r.MatchString(matrix[row+1][col+1]) {
					numSet[getNum(row + 1, col + 1, matrix)] = true
				}

				if len(numSet) == 2 {
					running := 1
					for key, _ := range numSet {
						running *= key
					}

					sum += running
				}
			}
		}
	}
	return sum
}

func getNum(row, col int,  matrix [][]string) int {
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


func checkNeighbors(row, col, numLength, numRows, numCols int, matrix [][]string) bool {
	r := regexp.MustCompile(`\d`)
	if row == 0 {
		for k := util.Max(col - 1, 0); k < util.Min(numCols, col + numLength + 1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) || 
			(matrix[row + 1][k] != "." && !r.MatchString(matrix[row + 1][k])) {
				return true
			}
		}
	} else if row == numRows {
		for k := util.Max(col - 1, 0); k < util.Min(numCols, col + numLength + 1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) || 
			(matrix[row - 1][k] != "." && !r.MatchString(matrix[row - 1][k])) {
				return true
			}
		}
	} else {
		for k := util.Max(col - 1, 0); k < util.Min(numCols, col + numLength + 1); k++ {
			if (matrix[row][k] != "." && !r.MatchString(matrix[row][k])) || 
			(matrix[row - 1][k] != "." && !r.MatchString(matrix[row - 1][k])) ||
			(matrix[row + 1][k] != "." && !r.MatchString(matrix[row + 1][k])) {
				return true
			}
		}
	}
	return false
}













